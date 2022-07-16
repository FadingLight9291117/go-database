package statementType

type StatementType byte

const (
	StatementInsert StatementType = iota
	StatementSelect
)
