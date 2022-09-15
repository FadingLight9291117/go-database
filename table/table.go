package table

import (
	"com.fadinglight/db/BTree"
)

type Table struct {
	Pager       *Pager
	RootPageNum int
}

func NewTable(filename string) *Table { return new(Table).init(filename) }

func (t *Table) init(filename string) *Table {
	pager := NewPager(filename)
	t.Pager = pager
	t.RootPageNum = 0
	if pager.PageNums == 0 {
		// New database file. Initialize page 0 as leaf node.
		t.Pager.GetPage(0)
	}

	return t
}

//func (t *Table) IsFull() bool {
//	return t.Size == MAX_PAGE*BTree.ROWS_PER_PAGE
//}

//func (t *Table) Insert(r *BTree.Row) *Table {
//	if t.IsFull() {
//		panic("table is full.")
//	}
//	pageNum := t.Size / BTree.ROWS_PER_PAGE
//	p := t.Pager.GetPage(pageNum)
//	p.CellNums++
//	p.Cells[p.CellNums] = BTree.LeafNodeCell{
//		Key:   r.Id,
//		Value: *r,
//	}
//
//	return t
//}

func (t *Table) Select() []*BTree.Row {
	rows := make([]*BTree.Row, 0)

	pageNum := t.Pager.FileSize / BTree.PAGE_SIZE

	for i := 0; i < pageNum; i++ {
		page := t.Pager.GetPage(i)
		for j := 0; j < int(page.CellNums); j++ {
			rows = append(rows, &page.Cells[j].Value)
		}
	}

	return rows
}

func (t *Table) Close() error {
	pager := t.Pager
	// FIXME: This only save a single node table
	err := pager.FlushOnePage(t.RootPageNum)
	if err != nil {
		return err
	}
	err = t.Pager.File.Close()
	if err != nil {
		return err
	}

	return nil
}

func (t *Table) CreateNewRoot(rightChildPageNum int) {
	/*
		Handle splitting thr root
		Old root copied to new Pag, becomes left child.
		Address of right child passed in.
		Re-initialize root page to contain the new root node.
		New root node points to the two child
	*/
	root := t.Pager.GetPage(t.RootPageNum)
	rightChild := t.Pager.GetPage(rightChildPageNum)
	leftChild := t.Pager.GetPage(t.Pager.GetUnusedPageNum())
	leftChild.CopyFrom(root)
	root
}
