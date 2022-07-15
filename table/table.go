package model

const TableMaxRows = 40960

type Table struct {
	NumRows uint32
	Rows    []Row
}

func (t *Table) init() {
	t.NumRows = 0
	t.Rows = make([]Row, 0, TableMaxRows)
}

func New() *Table { return new(Table) }
func (t *Table) IsFull() bool {
	return t.NumRows == TableMaxRows
}
