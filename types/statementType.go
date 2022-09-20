package types

import (
	"com.fadinglight/db/btree"
)

type StatementType byte

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	StatementType StatementType
	Row2Insert    *BTree.Row //only used by insert statement
}
