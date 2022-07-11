package main

import (
	"fmt"
)

func print_prompt() {
	fmt.Print("db> ")
}

func read_input() string {
	var input string
	fmt.Scan(&input)
	return input
}

func main() {
	var input string
	for {
		print_prompt()
		input = read_input()
		if input == ".exit" {
			return
		} else {
			fmt.Printf("Unrecognized command '%s' .\n", input)
		}
	}
}
