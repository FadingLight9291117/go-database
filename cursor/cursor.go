package cursor

import "com.fadinglight/db/table"

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
	RowNum     int  // 指向的row的位置
	EndOfTable bool // past the end of the table
}

// Value return the pointed row of the cursor
func (c *Cursor) Value() *table.Row {
	pageNum := c.RowNum / table.ROWS_PER_PAGE
	page := c.Table.Pager.GetPage(pageNum)
	rowOffset := c.RowNum % table.ROWS_PER_PAGE
	return &page.Rows[rowOffset]
}

func (c *Cursor) IsEnd() bool {
	return c.EndOfTable
}

func (c *Cursor) Next() *table.Row {
	var row *table.Row
	if !c.IsEnd() {
		row = c.Value()
		c.RowNum++
	}
	if c.RowNum >= c.Table.Size {
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
	cursor.RowNum = 0
	cursor.EndOfTable = table.Size == 0
	return cursor
}

// tableEnd returns a cursor at the end of the table
func tableEnd(table *table.Table) *Cursor {
	cursor := &Cursor{}
	cursor.Table = table
	cursor.RowNum = table.Size
	cursor.EndOfTable = true
	return cursor
}
