package main

import "fmt"

type Node struct {
	key int
}

func main() {

	//stre := "insert 1 as  as "
	//
	//words_ := strings.Split(stre, " ")
	//
	//// filter
	//words := make([]string, 0, len(words_))
	//for _, word := range words_ {
	//	if strings.TrimSpace(word) != "" {
	//		words = append(words, word)
	//	}
	//}
	//
	//for i, v := range words {
	//	fmt.Printf("%d => %v\n", i, v)
	//}

	//a := []int{1, 2, 3, 4}
	//
	//fmt.Printf("golang 指针长度: %d个字节\n", unsafe.Sizeof(&a))
	//fmt.Printf("golang bool长度：%d个字节\n", unsafe.Sizeof(true))
	//fmt.Printf("golang int 长度：%d个字节\n", unsafe.Sizeof(int(0)))
	nodes := make([]Node, 10)
	for i := range nodes {
		nodes[i].key = i
	}
	fmt.Printf("%v\n", nodes)
	// insert a node
	newNode := Node{key: -1}

	for i := len(nodes) - 2; i > 3; i-- {
		nodes[i+1] = nodes[i]
	}
	nodes[3] = newNode
	fmt.Printf("%v\n", nodes)
}
