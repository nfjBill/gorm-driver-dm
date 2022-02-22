package clauses

import (
	"gorm.io/gorm/clause"
)

type WhenMatched struct {
	clause.Set
	Where, Delete clause.Where
}

func (w WhenMatched) Name() string {
	return "WHEN MATCHED"
}

func (w WhenMatched) Build(builder clause.Builder) {
	if len(w.Set) > 0 {
		builder.WriteString("THEN")
		builder.WriteString(" UPDATE")
		builder.WriteString(" SET")
		builder.WriteByte(' ')
		w.Set.Build(builder)

		buildWhere := func(where clause.Where) {
			builder.WriteString(where.Name())
			builder.WriteByte(' ')
			where.Build(builder)
		}

		if len(w.Where.Exprs) > 0 {
			buildWhere(w.Where)
		}

		if len(w.Delete.Exprs) > 0 {
			builder.WriteString(" DELETE ")
			buildWhere(w.Delete)
		}
	}
}

func (w WhenMatched) MergeClause(clause *clause.Clause) {
	clause.Name = w.Name()
	clause.Expression = w
}
