package executor

import (
	"com.fadinglight/db/BTree"
	"com.fadinglight/db/cursor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"errors"
)

func ExecuteInsert(statement *types.Statement, t *table.Table) (ExecuteResult, error) {
	p := t.Pager.GetPage(t.RootPageNum)
	if int(p.CellNums) >= BTree.LEAF_NODE_MAX_CELLS {
		return EXECUTE_TABLE_FULL, errors.New("error: Table full")
	}

	row2Insert := statement.Row2Insert
	c := cursor.CreateCursor(t, true) // create a cursor pasted the end of the t
	err := c.InsertLeafNode(int(row2Insert.Id), row2Insert)
	if err != nil {
		return EXECUTE_ERROR, err
	}

	return EXECUTE_SUCESS, nil
}
