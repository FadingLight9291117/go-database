package executor

import (
	"com.fadinglight/db/cursor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"errors"
)

func ExecuteInsert(statement *types.Statement, table *table.Table) (ExecuteResult, error) {
	if table.IsFull() {
		return EXECUTE_TABLE_FULL, errors.New("error: Table full")
	}
	row2Insert := statement.Row2Insert

	c := cursor.CreateCursor(table, true) // create a cursor pasted the end of the table
	c.Value().Copy(row2Insert)

	table.Size++
	//table.Insert(row2Insert)

	return EXECUTE_SUCESS, nil
}
