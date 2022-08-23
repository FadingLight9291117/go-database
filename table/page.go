package table

type Page struct {
	Rows [PAGE_SIZE]Row
}

func (p *Page) Serialize() ([]byte, error) {
	bs := make([]byte, 0, len(p.Rows)*RowSize)
	for _, r := range p.Rows {
		br, err := r.Serialize()
		if err != nil {
			return nil, err
		}
		bs = append(bs, br...)
	}
	return bs, nil
}

func DeserializePage(bs []byte) (*Page, error) {
	rows := [PAGE_SIZE]Row{}
	for i, _ := range rows {
		r, err := DeserializeRow(bs[i*RowSize : i*RowSize+RowSize])
		if err != nil {
			return nil, err
		}
		rows[i] = *r
	}
	return &Page{rows}, nil
}
