package dm

import (
	"fmt"
	"gorm.io/gorm/schema"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/migrator"
)

type Migrator struct {
	migrator.Migrator
}

type BuildIndexOptionsInterface interface {
	BuildIndexOptions([]schema.IndexOption, *gorm.Statement) []interface{}
}

func (m Migrator) CurrentDatabase() (name string) {
	m.DB.Raw("SELECT SYS_CONTEXT ('userenv', 'current_schema') FROM DUAL").Row().Scan(&name)
	return
}

func buildConstraint(constraint *schema.Constraint) (sql string, results []interface{}) {
	sql = "CONSTRAINT ? FOREIGN KEY ? REFERENCES ??"
	if constraint.OnDelete != "" {
		sql += " ON DELETE " + constraint.OnDelete
	}

	if constraint.OnUpdate != "" {
		sql += " ON UPDATE " + constraint.OnUpdate
	}

	var foreignKeys, references []interface{}
	for _, field := range constraint.ForeignKeys {
		foreignKeys = append(foreignKeys, clause.Column{Name: field.DBName})
	}

	for _, field := range constraint.References {
		references = append(references, clause.Column{Name: field.DBName})
	}
	results = append(results, clause.Table{Name: constraint.Name}, foreignKeys, clause.Table{Name: constraint.ReferenceSchema.Table}, references)
	return
}

func (m Migrator) CreateIndex(value interface{}, name string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if idx := stmt.Schema.LookIndex(name); idx != nil {
			opts := m.DB.Migrator().(BuildIndexOptionsInterface).BuildIndexOptions(idx.Fields, stmt)
			values := []interface{}{clause.Column{Name: m.Migrator.DB.NamingStrategy.IndexName(stmt.Table, idx.Name)}, m.CurrentTable(stmt), opts}

			createIndexSQL := "CREATE "
			if idx.Class != "" {
				createIndexSQL += idx.Class + " "
			}
			createIndexSQL += "INDEX ? ON ??"

			if idx.Type != "" {
				createIndexSQL += " USING " + idx.Type
			}

			//if idx.Comment != "" {
			//	createIndexSQL += fmt.Sprintf(" COMMENT '%s'", idx.Comment)
			//}

			if idx.Option != "" {
				createIndexSQL += " " + idx.Option
			}

			return m.DB.Exec(createIndexSQL, values...).Error
		}

		return fmt.Errorf("failed to create index with name %s", name)
	})
}

func (m Migrator) CreateTable(values ...interface{}) error {
	for _, value := range values {
		m.TryQuotifyReservedWords(value)
		m.TryRemoveOnUpdate(value)
	}

	for _, value := range m.ReorderModels(values, false) {
		tx := m.DB.Session(&gorm.Session{})
		if err := m.RunWithValue(value, func(stmt *gorm.Statement) (errr error) {
			var (
				createTableSQL          = "CREATE TABLE ? ("
				values                  = []interface{}{m.CurrentTable(stmt)}
				hasPrimaryKeyInDataType bool
			)

			for _, dbName := range stmt.Schema.DBNames {
				field := stmt.Schema.FieldsByDBName[dbName]
				if !field.IgnoreMigration {
					createTableSQL += "? ?"
					hasPrimaryKeyInDataType = hasPrimaryKeyInDataType || strings.Contains(strings.ToUpper(string(field.DataType)), "PRIMARY KEY")
					f := m.DB.Migrator().FullDataTypeOf(field)
					if field.AutoIncrement {
						f.SQL = "INTEGER IDENTITY(1, " + strconv.FormatInt(field.AutoIncrementIncrement, 10) + ")"
					}
					values = append(values, clause.Column{Name: dbName}, f)
					createTableSQL += ","
				}
			}

			if !hasPrimaryKeyInDataType && len(stmt.Schema.PrimaryFields) > 0 {
				createTableSQL += "PRIMARY KEY ?,"
				primaryKeys := []interface{}{}
				for _, field := range stmt.Schema.PrimaryFields {
					primaryKeys = append(primaryKeys, clause.Column{Name: field.DBName})
				}

				values = append(values, primaryKeys)
			}

			for _, idx := range stmt.Schema.ParseIndexes() {
				if m.CreateIndexAfterCreateTable {
					defer func(value interface{}, name string) {
						if errr == nil {
							//errr = tx.Migrator().CreateIndex(value, name)
							errr = m.CreateIndex(value, name)
						}
					}(value, idx.Name)
				} else {
					if idx.Class != "" {
						createTableSQL += idx.Class + " "
					}
					createTableSQL += "INDEX ? ?"

					if idx.Comment != "" {
						createTableSQL += fmt.Sprintf(" COMMENT '%s'", idx.Comment)
					}

					if idx.Option != "" {
						createTableSQL += " " + idx.Option
					}

					createTableSQL += ","
					values = append(values, clause.Expr{SQL: idx.Name}, tx.Migrator().(migrator.BuildIndexOptionsInterface).BuildIndexOptions(idx.Fields, stmt))
				}
			}

			for _, rel := range stmt.Schema.Relationships.Relations {
				if !m.DB.DisableForeignKeyConstraintWhenMigrating {
					if constraint := rel.ParseConstraint(); constraint != nil {
						if constraint.Schema == stmt.Schema {
							sql, vars := buildConstraint(constraint)
							createTableSQL += sql + ","
							values = append(values, vars...)
						}
					}
				}
			}
			for _, chk := range stmt.Schema.ParseCheckConstraints() {
				createTableSQL += "CONSTRAINT ? CHECK (?),"
				values = append(values, clause.Column{Name: chk.Name}, clause.Expr{SQL: chk.Constraint})
			}

			createTableSQL = strings.TrimSuffix(createTableSQL, ",")

			createTableSQL += ")"

			if tableOption, ok := m.DB.Get("gorm:table_options"); ok {
				createTableSQL += fmt.Sprint(tableOption)
			}

			errr = tx.Exec(createTableSQL, values...).Error
			return errr
		}); err != nil {
			return err
		}
	}
	return nil
}

func (m Migrator) DropTable(values ...interface{}) error {
	values = m.ReorderModels(values, false)
	for i := len(values) - 1; i >= 0; i-- {
		value := values[i]
		tx := m.DB.Session(&gorm.Session{})
		if m.HasTable(value) {
			if err := m.RunWithValue(value, func(stmt *gorm.Statement) error {
				return tx.Exec("DROP TABLE ? CASCADE CONSTRAINTS", clause.Table{Name: stmt.Table}).Error
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m Migrator) HasTable(value interface{}) bool {
	var count int64

	m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw("SELECT COUNT(*) FROM USER_TABLES WHERE TABLE_NAME = ?", stmt.Table).Row().Scan(&count)
	})

	return count > 0
}

func (m Migrator) RenameTable(oldName, newName interface{}) (err error) {
	resolveTable := func(name interface{}) (result string, err error) {
		if v, ok := name.(string); ok {
			result = v
		} else {
			stmt := &gorm.Statement{DB: m.DB}
			if err = stmt.Parse(name); err == nil {
				result = stmt.Table
			}
		}
		return
	}

	var oldTable, newTable string

	if oldTable, err = resolveTable(oldName); err != nil {
		return
	}

	if newTable, err = resolveTable(newName); err != nil {
		return
	}

	if !m.HasTable(oldTable) {
		return
	}

	return m.DB.Exec("RENAME TABLE ? TO ?",
		clause.Table{Name: oldTable},
		clause.Table{Name: newTable},
	).Error
}

func (m Migrator) AddColumn(value interface{}, field string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if field := stmt.Schema.LookUpField(field); field != nil {
			return m.DB.Exec(
				"ALTER TABLE ? ADD ? ?",
				clause.Table{Name: stmt.Table}, clause.Column{Name: field.DBName}, m.DB.Migrator().FullDataTypeOf(field),
			).Error
		}
		return fmt.Errorf("failed to look up field with name: %s", field)
	})
}

func (m Migrator) DropColumn(value interface{}, name string) error {
	if !m.HasColumn(value, name) {
		return nil
	}

	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if field := stmt.Schema.LookUpField(name); field != nil {
			name = field.DBName
		}

		return m.DB.Exec(
			"ALTER TABLE ? DROP ?",
			clause.Table{Name: stmt.Table},
			clause.Column{Name: name},
		).Error
	})
}

func (m Migrator) AlterColumn(value interface{}, field string) error {
	if !m.HasColumn(value, field) {
		return nil
	}

	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if field := stmt.Schema.LookUpField(field); field != nil {
			return m.DB.Exec(
				"ALTER TABLE ? MODIFY ? ?",
				clause.Table{Name: stmt.Table},
				clause.Column{Name: field.DBName},
				m.FullDataTypeOf(field),
			).Error
		}
		return fmt.Errorf("failed to look up field with name: %s", field)
	})
}

func (m Migrator) HasColumn(value interface{}, field string) bool {
	var count int64
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw("SELECT COUNT(*) FROM USER_TAB_COLUMNS WHERE TABLE_NAME = ? AND COLUMN_NAME = ?", stmt.Table, field).Row().Scan(&count)
	}) == nil && count > 0
}

func (m Migrator) CreateConstraint(value interface{}, name string) error {
	m.TryRemoveOnUpdate(value)
	return m.Migrator.CreateConstraint(value, name)
}

func (m Migrator) DropConstraint(value interface{}, name string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		for _, chk := range stmt.Schema.ParseCheckConstraints() {
			if chk.Name == name {
				return m.DB.Exec(
					"ALTER TABLE ? DROP CHECK ?",
					clause.Table{Name: stmt.Table}, clause.Column{Name: name},
				).Error
			}
		}

		return m.DB.Exec(
			"ALTER TABLE ? DROP CONSTRAINT ?",
			clause.Table{Name: stmt.Table}, clause.Column{Name: name},
		).Error
	})
}

func (m Migrator) HasConstraint(value interface{}, name string) bool {
	var count int64
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw(
			"SELECT COUNT(*) FROM USER_CONSTRAINTS WHERE TABLE_NAME = ? AND CONSTRAINT_NAME = ?", stmt.Table, name,
		).Row().Scan(&count)
	}) == nil && count > 0
}

func (m Migrator) DropIndex(value interface{}, name string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if idx := stmt.Schema.LookIndex(name); idx != nil {
			name = idx.Name
		}

		return m.DB.Exec("DROP INDEX ?", clause.Column{Name: name}, clause.Table{Name: stmt.Table}).Error
	})
}

func (m Migrator) HasIndex(value interface{}, name string) bool {
	var count int64
	m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if idx := stmt.Schema.LookIndex(name); idx != nil {
			name = idx.Name
		}

		return m.DB.Raw(
			fmt.Sprintf(`SELECT COUNT(*) FROM USER_INDEXES WHERE TABLE_NAME = ('%s') AND INDEX_NAME = ('%s')`,
				m.Migrator.DB.NamingStrategy.TableName(stmt.Table),
				m.Migrator.DB.NamingStrategy.IndexName(stmt.Table, name),
			),
		).Row().Scan(&count)
	})

	return count > 0
}

// https://docs.dm.com/database/121/SPATL/alter-index-rename.htm
func (m Migrator) RenameIndex(value interface{}, oldName, newName string) error {
	panic("TODO")
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Exec(
			"ALTER INDEX ?.? RENAME TO ?", // wat
			clause.Table{Name: stmt.Table}, clause.Column{Name: oldName}, clause.Column{Name: newName},
		).Error
	})
}

func (m Migrator) TryRemoveOnUpdate(values ...interface{}) error {
	for _, value := range values {
		if err := m.RunWithValue(value, func(stmt *gorm.Statement) error {
			for _, rel := range stmt.Schema.Relationships.Relations {
				constraint := rel.ParseConstraint()
				if constraint != nil {
					rel.Field.TagSettings["CONSTRAINT"] = strings.ReplaceAll(rel.Field.TagSettings["CONSTRAINT"], fmt.Sprintf("ON UPDATE %s", constraint.OnUpdate), "")
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func (m Migrator) TryQuotifyReservedWords(values ...interface{}) error {
	for _, value := range values {
		if err := m.RunWithValue(value, func(stmt *gorm.Statement) error {
			for idx, v := range stmt.Schema.DBNames {
				if IsReservedWord(v) {
					stmt.Schema.DBNames[idx] = fmt.Sprintf(`"%s"`, v)
				}
			}

			for _, v := range stmt.Schema.Fields {
				if IsReservedWord(v.DBName) {
					v.DBName = fmt.Sprintf(`"%s"`, v.DBName)
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}
