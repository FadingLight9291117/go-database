package executor

import (
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"errors"
)

func ExecuteInsert(statement *types.Statement, table *table.Table) (ExecuteResult, error) {
	if table.IsFull() {
		return EXECUTE_TABLE_FULL, errors.New("error: Table full")
	}
	row2Insert := statement.Row2Insert
	table.Insert(row2Insert)

	return EXECUTE_SUCESS, nil
}
