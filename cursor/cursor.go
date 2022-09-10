package cursor

import (
	"com.fadinglight/db/BTree"
	"com.fadinglight/db/table"
	"errors"
)

/*
* 	Cursor 可以做的事情
*	1. create a cursor at the beginning of the table
* 	2. create a cursor at the end of the table
*	3. access the row the cursor is pointing to
*	4. advance the cursor to the next row
 */

/*
 *	Cursor 的操作
 *	1. delete the row pointed by a cursor
 *	2. modify the row pointed by a cursor
 *	3. search a table for a given ID, and create a cursor pointing to the ROW with that ID
 */

type Cursor struct {
	Table      *table.Table
	PageNum    int
	CellNum    int
	EndOfTable bool // past the end of the table
}

// Value return the pointed row of the cursor
func (c *Cursor) Value() *BTree.Row {
	page := c.Table.Pager.GetPage(c.PageNum)
	return &page.Cells[c.CellNum].Value
}

func (c *Cursor) IsEnd() bool {
	return c.EndOfTable
}

func (c *Cursor) Next() *BTree.Row {
	row := c.Value()
	node := c.Table.Pager.GetPage(c.PageNum)
	c.CellNum++
	if c.CellNum >= int(node.CellNums) {
		c.EndOfTable = true
	}

	return row
}

func CreateCursor(table *table.Table, end bool) *Cursor {
	if end {
		return tableEnd(table)
	} else {
		return tableStart(table)
	}
}

// tableStart return a cursor at the beginning of the table
func tableStart(table *table.Table) *Cursor {
	cursor := &Cursor{}
	cursor.Table = table
	cursor.PageNum = table.RootPageNum
	cursor.CellNum = 0
	rootNode := table.Pager.GetPage(table.RootPageNum)
	cursor.EndOfTable = rootNode.CellNums == 0
	return cursor
}

// tableEnd returns a cursor at the end of the table
func tableEnd(table *table.Table) *Cursor {
	cursor := &Cursor{}
	cursor.Table = table
	cursor.PageNum = table.RootPageNum
	rootNode := table.Pager.GetPage(table.RootPageNum)
	cursor.CellNum = int(rootNode.CellNums)
	cursor.EndOfTable = true
	return cursor
}

func (c *Cursor) InsertLeafNode(key int, value *BTree.Row) error {
	page := c.Table.Pager.GetPage(c.PageNum)
	if int(page.CellNums) >= BTree.LEAF_NODE_MAX_CELLS {
		// Page is full
		return errors.New("need to implement splitting a leaf node")
	}
	if c.CellNum < int(page.CellNums) {
		// make room for new cell
		for i := int(page.CellNums); i > c.CellNum; i-- {
			page.Cells[i+1] = page.Cells[i]
		}
	}
	page.CellNums++
	page.Cells[c.CellNum] = BTree.NodeCell{
		Key:   uint32(key),
		Value: *value,
	}
	return nil
}
