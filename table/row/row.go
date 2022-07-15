package table

import "fmt"

const ColumnUsernameSize = 32
const ColumnEmail = 255

type Row struct {
	Id       uint32
	Username [ColumnUsernameSize]rune
	Email    [ColumnEmail]rune
}

func (r *Row) Print() {
	fmt.Printf("(%d, %s, %s)", r.Id, string(r.Username[:]), string(r.Email[:]))
}
