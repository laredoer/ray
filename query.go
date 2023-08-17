package ray

import (
	"bytes"
	"context"
	"fmt"
)

type Table interface {
	TableName() string
}

type Builder struct {
	bytes.Buffer
	dialect string
	args    []any
	total   int // total number of parameters in query tree.
}

func (b Builder) clone() Builder {
	c := Builder{dialect: b.dialect, total: b.total}
	if len(b.args) > 0 {
		c.args = append(c.args, b.args...)
	}
	c.Buffer.Write(b.Bytes())
	return c
}

type Selector struct {
	Builder
	columns       []string
	joins         []any
	where         []whereState
	or            bool
	not           bool
	order         []string
	group         []string
	having        []claim
	limit         *int
	offset        *int
	distinct      bool
	distinctField []string
}

type Querier[T Table] struct {
	as   string
	from T
	sql  *Selector
}

func From[T Table](clauses ...Clause) *Querier[T] {

	var t T
	s := &Querier[T]{
		from: t,
		sql:  &Selector{},
	}
	for _, c := range clauses {
		c(s.sql)
	}
	return s
}

func (s *Querier[T]) All() ([]*T, error) {
	s.sqlAll(context.Background())
	return nil, nil
}

func (s *Querier[T]) First() (*T, error) {
	var t T
	return &t, nil
}

func (s *Querier[T]) sqlAll(ctx context.Context) {

	b := s.sql.Builder.clone()

	b.WriteString("SELECT ")
	if s.sql.distinct {
		b.WriteString("DISTINCT ")
	}

	if len(s.sql.columns) > 0 {
		for i, c := range s.sql.columns {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(c)
		}
	} else {
		b.WriteString("*")
	}
	b.WriteString(" FROM ")
	b.WriteString(s.from.TableName())

	if len(s.sql.where) != 0 {
		b.WriteString(" WHERE ")
		for i, w := range s.sql.where {
			if i > 0 {
				b.WriteString(" ")
				b.WriteString(w.predicate)
				b.WriteString(" ")
			}
			b.WriteString(w.clause)
			b.WriteString(" ")
			b.WriteString(string(w.op))
			b.WriteString(" ")
			b.WriteString("?")
			b.args = append(b.args, w.arg)
			b.total = len(b.args)
		}
	}

	fmt.Println(b.String(), b.args)
}
