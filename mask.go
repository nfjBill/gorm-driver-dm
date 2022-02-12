package dm

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

type STable struct {
	Conn  *gorm.DB
	Table interface{}
}

var conn *gorm.DB

func Table(table ...interface{}) *STable {
	var tb interface{}
	if len(table) == 1 {
		tb = table[0]
	}
	return &STable{Table: tb, Conn: conn}
}

func TB(db *gorm.DB) {
	conn = db
}

func (stb *STable) DB() (*sql.DB, error) {
	return stb.Conn.DB()
}

func (stb *STable) Model() *gorm.DB {
	if reflect.ValueOf(stb.Table).Kind() == reflect.String {
		return stb.Conn.Table(stb.Table.(string))
	}
	return stb.Conn.Model(stb.Table)
}

func (stb *STable) Create() error {
	return stb.Conn.Create(stb.Table).Error
}

func (stb *STable) Get(dest interface{}, conds ...interface{}) error {
	return stb.Conn.Where(stb.Table).First(dest, conds...).Error
}

func (stb *STable) GetWhere(dest interface{}) error {
	return stb.Conn.Model(stb.Table).Where(stb.Table).Find(dest).Error
}

func (stb *STable) GetAll(dest interface{}) error {
	return stb.Conn.Model(stb.Table).Find(dest).Error
}

func (stb *STable) ClausesAssignmentColumns(name string, doUpdates []string) error {
	return stb.Conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: name}},
		DoUpdates: clause.AssignmentColumns(doUpdates),
	}).Create(stb.Table).Error
}

func (stb *STable) Delete() error {
	tx := stb.Conn.Begin()
	if err := tx.Where(stb.Table).Delete(stb.Table).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (stb *STable) Update(dest interface{}) error {
	tx := stb.Conn.Begin()
	if err := tx.Where(stb.Table).Updates(dest).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (stb *STable) AutoMigrate(dst ...interface{}) error {
	return stb.Conn.AutoMigrate(dst...)
}

func (stb *STable) DropTable(dest ...interface{}) error {
	return stb.Conn.Migrator().DropTable(dest...)
}

func (stb *STable) HasTable(dest interface{}) bool {
	return stb.Conn.Migrator().HasTable(dest)
}
