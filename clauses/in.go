package clauses

import "gorm.io/gorm/clause"

// IN Whether a value is within a set of values
type IN struct {
	Column interface{}
	Values []interface{}
}

func (in IN) Build(builder clause.Builder) {
	builder.WriteQuoted(in.Column)
	switch len(in.Values) {
	case 0:
		builder.WriteString(" IN (NULL)")
	case 1:
		if _, ok := in.Column.([]clause.Column); ok {
			builder.WriteString(" = (")
			builder.AddVar(builder, in.Values...)
			builder.WriteString(")")
		} else {
			builder.WriteString(" = ")
			builder.AddVar(builder, in.Values...)
		}

	default:
		builder.WriteString(" IN (")
		builder.AddVar(builder, in.Values...)
		builder.WriteByte(')')

	}
}
