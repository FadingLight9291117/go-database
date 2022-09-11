package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type NodeType = uint8

type Foo struct {
	key int32
}

type CommonNodeHeader struct {
	Type      NodeType
	IsRoot    bool
	ParentPtr *CommonNodeHeader // FIXME: 无法序列化
}

func main() {
	p := &CommonNodeHeader{IsRoot: true}
	buf := new(bytes.Buffer)
	data := []any{p.Type, p.IsRoot, uint64(0)}
	for _, v := range data {
		binary.Write(buf, binary.BigEndian, v)
	}
	fmt.Printf("%x\n", buf.Bytes())
}
