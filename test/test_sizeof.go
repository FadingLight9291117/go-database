package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type Foo2 struct {
	Test  int8
	Tset1 int64
}

func main() {
	foo := Foo2{}
	fmt.Printf("CommonNodeHeader size is %d\n", unsafe.Sizeof(foo))

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, foo)
	if err != nil {
		return
	}

	fmt.Printf("CommonNodeHeader byes size of is %d\n", len(buf.Bytes()))

}
