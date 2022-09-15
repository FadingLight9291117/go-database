package main

import (
	"fmt"
	"unsafe"
)

type Foo1 struct {
	Value unsafe.Pointer
}

func main() {
	foo := &Foo1{}

	s := "test"

	foo.Value = unsafe.Pointer(&s)

	fmt.Printf("%v\n", unsafe.Sizeof(foo.Value))
}
