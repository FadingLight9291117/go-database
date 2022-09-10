package table

import "com.fadinglight/db/BTree"

type Table struct {
	Size        int
	Pager       *Pager
	RootPageNum int
}

func NewTable(filename string) *Table { return new(Table).init(filename) }

func (t *Table) init(filename string) *Table {
	pager := NewPager(filename)
	t.Pager = pager
	t.Size = pager.FileSize / BTree.ROW_SIZE

	return t
}

func (t *Table) IsFull() bool { return t.Size == MAX_PAGE*BTree.ROWS_PER_PAGE }

func (t *Table) Insert(r *BTree.Row) *Table {
	if t.IsFull() {
		panic("table is full.")
	}
	pageNum := t.Size / BTree.ROWS_PER_PAGE
	p := t.Pager.GetPage(pageNum)
	p.Cells[p.CellNums] = BTree.NodeCell{
		Key:   r.Id,
		Value: *r,
	}
	p.CellNums++

	return t
}

func (t *Table) Select() []*BTree.Row {
	rows := make([]*BTree.Row, 0, t.Size)

	pageNum := t.Size / BTree.PAGE_SIZE

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
	//fullPageNum := t.Size / ROWS_PER_PAGE
	for i := 0; i < len(t.Pager.Pages); i++ {
		if pager.Pages[i] != nil {
			if err := pager.FlushOnePage(i); err != nil {
				return err
			}
			pager.Pages[i] = nil
		}
	}
	if err := t.Pager.File.Close(); err != nil {
		return err
	}
	return nil
}
