package ray

import (
	"golang.org/x/exp/constraints"
)

func Select(columns ...string) Clause {
	return func(s *Selector) {
		s.columns = append(s.columns, columns...)
	}
}

func Limit[T constraints.Integer](limit T) Clause {
	return func(s *Selector) {
		i := int(limit)
		s.limit = &i
	}
}

func Offset[T constraints.Integer](offset T) Clause {
	return func(s *Selector) {
		i := int(offset)
		s.offset = &i
	}
}

func Order(o string) Clause {
	return func(s *Selector) {
		s.order = append(s.order, o)
	}
}

func GroupBy(g ...string) Clause {
	return func(s *Selector) {
		s.group = append(s.group, g...)
	}
}

func Distinct(idents ...string) Clause {
	return func(s *Selector) {
		s.distinct = true
		s.distinctField = append(s.distinctField, idents...)
	}
}

func Having(clause string, op Op, arg any) Clause {
	return func(s *Selector) {
		s.having = append(s.having, claim{clause, op, arg})
	}
}

func Where(clause string, op Op, arg any) Clause {
	return func(s *Selector) {
		s.where = append(s.where, whereState{"AND", claim{clause, op, arg}})
	}
}

func And(clause string, op Op, arg any) Clause {
	return func(s *Selector) {
		s.where = append(s.where, whereState{"AND", claim{clause, op, arg}})
	}
}

func Or(clause string, op Op, arg any) Clause {
	return func(s *Selector) {
		s.where = append(s.where, whereState{"OR", claim{clause, op, arg}})
	}
}
