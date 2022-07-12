package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputBuffer struct {
	buffer string
}

type ExitType uint8

const (
	ExitNormal ExitType = iota
	ExitUnNormal
)

type MetaCommandResult uint8

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

type PrepareResult uint8

const (
	PrepareSuccess PrepareResult = iota
	PrepareUnrecognizedStatement
)

type StatementType uint8

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	statementType StatementType
}

func printPrompt() {
	fmt.Print("db> ")
}

func readInput() *InputBuffer {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return &InputBuffer{}
	}
	return &InputBuffer{buffer: input}
}

func exit(exitType ExitType) {
	switch exitType {
	case ExitNormal:
		os.Exit(0)
	case ExitUnNormal:
		os.Exit(1)
	}
}

func DoMetaCommand(inputBuffer *InputBuffer) MetaCommandResult {
	if inputBuffer.buffer == ".exit" {
		exit(ExitNormal)
		return MetaCommandSuccess
	} else {
		return MetaCommandUnrecognizedCommand
	}
}

// PrepareStatement SQL Compiler
func PrepareStatement(inputBuffer *InputBuffer, statement *Statement) PrepareResult {
	if strings.HasPrefix(inputBuffer.buffer, "insert") {
		statement.statementType = StatementInsert
		return PrepareSuccess
	}
	if inputBuffer.buffer == "select" {
		statement.statementType = StatementSelect
		return PrepareSuccess
	}
	return PrepareUnrecognizedStatement
}

func ExecuteStatement(statement *Statement) {
	switch statement.statementType {
	case StatementInsert:
		fmt.Println("This is where we would do an insert.")
	case StatementSelect:
		fmt.Println("This is where we would do a select.")
	}
}

func main() {
	var inputBuffer *InputBuffer
	for {
		printPrompt()
		inputBuffer = readInput()
		// 处理 `meta-command`
		// 类似 `.exit`以 `.`开头的命令被称为 `meta-command`(元命令)
		if inputBuffer.buffer[0] == '.' {
			switch DoMetaCommand(inputBuffer) {
			case MetaCommandSuccess:
				continue
			case MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command '%s' .\n", inputBuffer.buffer)
				continue
			}
		}

		statement := &Statement{}
		switch PrepareStatement(inputBuffer, statement) {
		case PrepareSuccess:
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", inputBuffer.buffer)
			continue
		}

		ExecuteStatement(statement)

		fmt.Println("Executed.")
	}
}
