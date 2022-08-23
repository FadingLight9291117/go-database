package types

import "com.fadinglight/db/table"

type StatementType byte

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	StatementType StatementType
	Row2Insert    *table.Row //only used by insert statement
}

