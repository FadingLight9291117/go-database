package table

import "com.fadinglight/database/table/row"

const MaxRows = 40960

type Table struct {
	Rows []row.Row
}

func (t *Table) init() *Table {
	t.Rows = make([]row.Row, 0, MaxRows)
	return t
}

func New() *Table { return new(Table).init() }

func (t *Table) Free() {
	t.Rows = nil
}

func (t *Table) IsFull() bool {
	return len(t.Rows) == cap(t.Rows)
}

func (t *Table) Append(row2 *row.Row) *Table {
	t.Rows = append(t.Rows, *row2)
	return t
}
