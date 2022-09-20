package executor

import (
	"com.fadinglight/db/btree"
	"com.fadinglight/db/cursor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"errors"
)

func ExecuteInsert(statement *types.Statement, t *table.Table) (ExecuteResult, error) {
	row2Insert := statement.Row2Insert
	key2Inset := row2Insert.Id
	c := cursor.FindInTable(t, key2Inset) // create a cursor pasted the end of the t

	node := c.Table.Pager.GetPage(c.PageNum, 0).Node.(*BTree.LeafNode)

	if c.CellNum < int(node.CellNums) {
		keyAtIndex := node.Cells[c.CellNum].Key
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
