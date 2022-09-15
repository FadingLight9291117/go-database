package BTree

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

// 1个rune是4个字节
const (
	ID_SIZE       = int(unsafe.Sizeof(uint32(0)))
	USERNAME_SIZE = int(ColumnUsernameSize * unsafe.Sizeof(rune(' ')))
	EMAIL_SIZE    = int(ColumnEmailSize * unsafe.Sizeof(rune(' ')))
)

const (
	ColumnUsernameSize = 8
	ColumnEmailSize    = 32
)

type Row struct {
	Id       uint32
	Username [ColumnUsernameSize]rune
	Email    [ColumnEmailSize]rune
}

func (r *Row) Copy(or *Row) {
	r.Id = or.Id
	r.Username = or.Username
	r.Email = or.Email
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

func (r *Row) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, &r)
	if err != nil {
		return err
	}
	return nil
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
