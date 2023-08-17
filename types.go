package ray

type Clause func(*Selector)

type Op string

const (
	In    Op = "IN"
	NotIn Op = "NOT IN"
	Like  Op = "LIKE"
	// Eq    Op = "="
	// Ne    Op = "!="
	// Gt    Op = ">"
	// Ge    Op = ">="
	// Lt    Op = "<"
	// Le    Op = "<="
)

type whereState struct {
	predicate string // and , or
	claim
}

// id >= 1
type claim struct {
	clause string
	op     Op
	arg    any
}
