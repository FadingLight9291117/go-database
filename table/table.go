package table

type Table struct {
	Size  int
	Pager *Pager
}

func New(filename string) *Table { return new(Table).init(filename) }

func (t *Table) init(filename string) *Table {
	pager := OpenPager(filename)
	t.Pager = pager
	t.Size = pager.FileSize / ROW_SIZE

	return t
}

func (t *Table) IsFull() bool { return t.Size == MAX_PAGE*ROWS_PER_PAGE }

func (t *Table) Insert(r *Row) *Table {
	if t.IsFull() {
		panic("table is full.")
	}
	pageNum := t.Size / ROWS_PER_PAGE
	p := t.Pager.GetPage(pageNum)
	pageOffset := t.Size % ROWS_PER_PAGE
	p.Rows[pageOffset] = *r
	t.Size += 1

	return t
}

func (t *Table) Select() []*Row {
	rows := make([]*Row, 0, t.Pager.FileSize/ROW_SIZE)

	pageNum := t.Size / ROWS_PER_PAGE
	if t.Size%ROWS_PER_PAGE != 0 {
		pageNum++
	}
	for i := 0; i < pageNum; i++ {
		page := t.Pager.GetPage(i)
		if t.Size-len(rows) < ROWS_PER_PAGE {
			for j := range page.Rows[:t.Size-len(rows)] {
				rows = append(rows, &page.Rows[j])
			}
			break
		}
		for j := range page.Rows {
			rows = append(rows, &page.Rows[j])
		}
	}

	return rows
}

func (t *Table) Close() error {
	pager := t.Pager
	fullPageNum := t.Size / ROWS_PER_PAGE
	for i := 0; i < fullPageNum; i++ {
		if pager.Pages[i] == nil {
			continue
		}
		if err := pager.FlushOnePage(i, ROWS_PER_PAGE); err != nil {
			return err
		}
		pager.Pages[i] = nil
	}
	additionRowNum := t.Size % ROWS_PER_PAGE
	if additionRowNum > 0 {
		thisPageIndex := fullPageNum
		if pager.Pages[thisPageIndex] != nil {
			if err := pager.FlushOnePage(thisPageIndex, additionRowNum); err != nil {
				return err
			}
			pager.Pages[thisPageIndex] = nil
		}
	}
	err := t.Pager.File.Close()
	if err != nil {
		return err
	}
	return nil
}
