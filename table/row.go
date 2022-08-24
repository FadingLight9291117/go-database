package table

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 1个rune是4个字节
const ID_SIZE = 4
const USERNAME_SIZE = ColumnUsernameSize * 4
const EMAIL_SIZE = ColumnEmailSize * 4
const ROW_SIZE = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE

const ColumnUsernameSize = 32
const ColumnEmailSize = 255

type Row struct {
	Id       uint32
	Username [ColumnUsernameSize]rune
	Email    [ColumnEmailSize]rune
}

func NewRow(id int, username string, email string) *Row {
	row := Row{}
	row.Id = uint32(id)
	copy(row.Username[:], []rune(username))
	copy(row.Email[:], []rune(email))
	return &row
}

func (r *Row) Print() {
	fmt.Printf("Row(id->%d, username->%s, email->%s)\n", r.Id, string(r.Username[:]), string(r.Email[:]))
}

func (r *Row) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, *r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeRow(byt []byte) (*Row, error) {
	buf := bytes.NewBuffer(byt)
	r := Row{}
	if err := binary.Read(buf, binary.BigEndian, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func main() {
	r := NewRow(100, "车亮召", "123@123.com")
	byts, _ := r.Serialize()
	fmt.Println(byts)
	rn, _ := DeserializeRow(byts)
	rn.Print()
}
