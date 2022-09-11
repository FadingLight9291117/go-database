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
	key2Inset := row2Insert.Id
	c := cursor.FindInTable(t, key2Inset) // create a cursor pasted the end of the t

	if c.CellNum < int(p.CellNums) {
		keyAtIndex := p.Cells[c.CellNum].Key
		if keyAtIndex == key2Inset {
			// key is duplicate
			return EXECUT_DUPLICATE_KEY, errors.New("error: duplicate key")
		}
	}

	err := c.InsertLeafNode(key2Inset, row2Insert)
	if err != nil {
		return EXECUTE_ERROR, err
	}

	return EXECUTE_SUCESS, nil
}
