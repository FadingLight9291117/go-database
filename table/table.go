package table

type Table struct {
	Size  int
	Pager *Pager
}

func New(filename string) *Table { return new(Table).init(filename) }

func (t *Table) init(filename string) *Table {
	t.Size = 0
	pager := OpenPager(filename)
	t.Pager = pager

	return t
}

func (t *Table) IsFull() bool { return t.Size == MAX_PAGE*PAGE_SIZE }

func (t *Table) Insert(r *Row) *Table {
	if t.IsFull() {
		panic("table is full.")
	}
	pageNum := t.Size / PAGE_SIZE
	p := t.Pager.GetPage(pageNum)
	pageOffset := t.Size % PAGE_SIZE
	p.Rows[pageOffset] = *r
	t.Size += 1

	return t
}

func (t *Table) Select() []*Row {
	rows := make([]*Row, 0, t.Size)

	for i := 0; i < t.Pager.Len(); i++ {
		page := t.Pager.GetPage(i)
		if t.Size-len(rows) < PAGE_SIZE {
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
