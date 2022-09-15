package table

import (
	"com.fadinglight/db/BTree"
)

type Page struct {
	BTree.LeafNode
}

func (p *Page) Serialize() ([]byte, error) {
	b, err := p.Serialize()
	return b, err
}

func DeserializePage(b []byte) (*Page, error) {
	p := &Page{}
	err := p.Deserialize(b)
	if err != nil {
		return nil, err
	}
	return p, err
}
