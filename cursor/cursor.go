package cursor

import (
	"com.fadinglight/db/BTree"
	"com.fadinglight/db/table"
	"errors"
	"fmt"
	"os"
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
	node, ok := c.Table.Pager.GetPage(c.PageNum).Node.(*BTree.LeafNode)
	if !ok {
		return nil
	}
	return &node.Cells[c.CellNum].Value
}

func (c *Cursor) IsEnd() bool {
	return c.EndOfTable
}

func (c *Cursor) Next() *BTree.Row {
	row := c.Value()
	node, ok := c.Table.Pager.GetPage(c.PageNum).Node.(*BTree.LeafNode)
	if !ok {
		return nil
	}
	c.CellNum++
	if c.CellNum >= int(node.CellNums) {
		c.EndOfTable = true
	}
	return row
}

func CreateStartCursor(table *table.Table) *Cursor {
	return tableStart(table)
}

// tableStart return a cursor at the beginning of the table
func tableStart(table *table.Table) *Cursor {
	cursor := &Cursor{}
	cursor.Table = table
	cursor.PageNum = table.RootPageNum
	cursor.CellNum = 0
	rootNode, ok := table.Pager.GetPage(table.RootPageNum).Node.(*BTree.LeafNode)
	if !ok {
		return nil
	}
	cursor.EndOfTable = rootNode.CellNums == 0

	return cursor
}

//// tableEnd returns a cursor at the end of the table
//func tableEnd(table *table.Table) *Cursor {
//	cursor := &Cursor{}
//	cursor.Table = table
//	cursor.PageNums = table.RootPageNum
//	rootNode := table.Pager.GetPage(table.RootPageNum)
//	cursor.CellNum = int(rootNode.CellNums)
//	cursor.EndOfTable = true
//	return cursor
//}

// FindInTable returns a cursor at the position of the given key
func FindInTable(t *table.Table, key2Insert uint64) *Cursor {
	rootNode := t.Pager.GetPage(t.RootPageNum)
	switch rootNode.Node.(type) {
	case *BTree.LeafNode:
		return findInLeafNode(t, t.RootPageNum, key2Insert)
	case *BTree.InternalNode:
		fmt.Println("Need to implement searching an internal node.")
		os.Exit(1)
		return nil
	default:
		return nil
	}
}

func findInLeafNode(t *table.Table, pageNum int, key uint64) *Cursor {
	if page, ok := t.Pager.GetPage(pageNum).Node.(*BTree.LeafNode); ok {

		cellNums := int(page.CellNums)

		c := &Cursor{
			Table:      t,
			PageNum:    pageNum,
			EndOfTable: false,
			CellNum:    int(page.CellNums),
		}

		// binary search
		lIndex := 0
		rIndex := cellNums
		for lIndex < rIndex {
			midIndex := (lIndex + rIndex) / 2
			if page.Cells[midIndex].Key == key {
				c.CellNum = midIndex
				return c
			}
			if page.Cells[midIndex].Key > key {
				rIndex = midIndex
			} else {
				lIndex = midIndex + 1
			}
		}

		c.CellNum = lIndex
		return c
	}
	return nil
}

func (c *Cursor) InsertLeafNode(key uint64, value *BTree.Row) error {
	page, ok := c.Table.Pager.GetPage(c.PageNum).Node.(*BTree.LeafNode)
	if !ok {
		return errors.New("error: not a leaf node")
	}
	if int(page.CellNums) >= BTree.LEAF_NODE_MAX_CELLS {
		// Page is full
		c.splitLeafNodeAndInsert(key, value)
		return nil
		//return errors.New("need to implement splitting a leaf node")
	}
	if c.CellNum < int(page.CellNums) {
		// make room for new cell
		for i := int(page.CellNums); i > c.CellNum; i-- {
			page.Cells[i] = page.Cells[i-1]
		}
	}
	page.CellNums++
	page.Cells[c.CellNum] = BTree.LeafNodeCell{
		Key:   key,
		Value: *value,
	}
	return nil
}

// splitLeafNodeAndInsert create a new Node and move half cells over;
// insert the new value in one of the two nodes;
// update parent or create a new parent.
func (c *Cursor) splitLeafNodeAndInsert(key uint64, value *BTree.Row) {
	newPageNum := c.Table.Pager.GetUnusedPageNum()
	oldNode, ok1 := c.Table.Pager.GetPage(c.PageNum).Node.(*BTree.LeafNode)
	newNode, ok2 := c.Table.Pager.GetPage(newPageNum).Node.(*BTree.LeafNode)

	if !(ok1 && ok2) {
		return
	}
	/*
	  All existing keys plus new key should be divided
	  evenly between old (left) and new (right) nodes.
	  Starting from the right, move each key to correct position.
	*/
	const leafNodeLeftSplitCount = (BTree.LEAF_NODE_MAX_CELLS + 1) / 2
	const leafNodeRightSplitCount = BTree.LEAF_NODE_MAX_CELLS + 1 - leafNodeLeftSplitCount
	cells := oldNode.Cells
	for i := BTree.LEAF_NODE_MAX_CELLS; i >= 0; i-- {
		var cell *BTree.LeafNodeCell
		if i == c.CellNum {
			cell = &BTree.LeafNodeCell{
				Key:   key,
				Value: *value,
			}
		} else if i > c.CellNum {
			cell = &cells[i-1]
		} else {
			cell = &cells[i]
		}
		indexInNode := i % leafNodeLeftSplitCount
		if i < leafNodeLeftSplitCount {
			// left
			oldNode.Cells[indexInNode] = *cell
		} else {
			//right
			newNode.Cells[indexInNode] = *cell
		}
	}
	oldNode.CellNums = uint64(leafNodeLeftSplitCount)
	newNode.CellNums = uint64(leafNodeRightSplitCount)

	if oldNode.IsRoot {
		c.Table.CreateNewRoot(newPageNum)
	} else {
		fmt.Println("need to implement updating parent after split.")
		os.Exit(1)
	}
}
