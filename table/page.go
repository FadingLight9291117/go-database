package table

type Page struct {
	Rows [ROWS_PER_PAGE]Row
}

func (p *Page) Serialize() ([]byte, error) {
	bs := make([]byte, 0, len(p.Rows)*ROW_SIZE)
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
	rows := [ROWS_PER_PAGE]Row{}
	for i, _ := range rows {
		r, err := DeserializeRow(bs[i*ROW_SIZE : i*ROW_SIZE+ROW_SIZE])
		if err != nil {
			return nil, err
		}
		rows[i] = *r
	}
	return &Page{rows}, nil
}
