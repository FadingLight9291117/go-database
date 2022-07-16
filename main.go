package main

import (
	"bufio"
	"com.fadinglight/database/table"
	"com.fadinglight/database/table/row"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputBuffer struct {
	Buffer string
}

type ExitType byte

const (
	ExitNormal ExitType = iota
	ExitUnNormal
)

type MetaCommandResult byte

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

type PrepareResult byte

const (
	PrepareSuccess PrepareResult = iota
	PrepareSyntaxError
	PrepareUnrecognizedStatement
)

type StatementType byte

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	statementType StatementType
	Row2Insert    *row.Row //only used by insert statement
}

type ExecuteResult = byte

const (
	ExecuteSuccess ExecuteResult = iota
	ExecuteTableFull
	ExecuteError
)

func printPrompt() {
	fmt.Print("db> ")
}

func readInput(scanner *bufio.Scanner) *InputBuffer {
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)

	return &InputBuffer{Buffer: input}
}

func exit(exitType ExitType) {
	switch exitType {
	case ExitNormal:
		os.Exit(1)
	case ExitUnNormal:
		os.Exit(0)
	}
}

func DoMetaCommand(inputBuffer *InputBuffer, table *table.Table) MetaCommandResult {
	if inputBuffer.Buffer == ".exit" {
		exit(ExitNormal)
		return MetaCommandSuccess
	} else {
		return MetaCommandUnrecognizedCommand
	}
}

// PrepareStatement SQL Compiler
func PrepareStatement(inputBuffer *InputBuffer, statement *Statement) PrepareResult {
	if strings.HasPrefix(inputBuffer.Buffer, "insert") {
		statement.statementType = StatementInsert
		words := strings.Split(inputBuffer.Buffer, " ")[1:]
		if len(words) < 3 {
			return PrepareSyntaxError
		}
		id, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Println("id 输入异常")
			return PrepareSyntaxError
		}
		r := &row.Row{}
		r.Id = uint32(id)
		copy(r.Username[:], []rune(words[1][:]))
		copy(r.Email[:], []rune(words[2][:]))
		statement.Row2Insert = r

		return PrepareSuccess
	}
	if inputBuffer.Buffer == "select" {
		statement.statementType = StatementSelect
		return PrepareSuccess
	}
	return PrepareUnrecognizedStatement
}

func ExecuteInsert(statement *Statement, table *table.Table) ExecuteResult {
	if table.IsFull() {
		return ExecuteTableFull
	}
	row2Insert := statement.Row2Insert
	table.Append(row2Insert)

	return ExecuteSuccess
}

func ExecuteSelect(statement *Statement, table *table.Table) ExecuteResult {
	for _, v := range table.Rows {
		v.Print()
	}
	return ExecuteSuccess
}

func ExecuteStatement(statement *Statement, table *table.Table) ExecuteResult {
	switch statement.statementType {
	case StatementInsert:
		fmt.Println("This is where we would do an insert.")
		return ExecuteInsert(statement, table)
	case StatementSelect:
		fmt.Println("This is where we would do a select.")
		return ExecuteSelect(statement, table)
	default:
		return ExecuteError
	}
}

func main() {
	t := table.New()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 提示信息
		printPrompt()
		inputBuffer := readInput(scanner)
		if inputBuffer.Buffer == "" {
			continue
		}
		// 处理 `meta-command`
		// 类似 `.exit`以 `.`开头的命令被称为 `meta-command`(元命令)
		if inputBuffer.Buffer[0] == '.' {
			switch DoMetaCommand(inputBuffer, t) {
			case MetaCommandSuccess:
				continue
			case MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command '%s' .\n", inputBuffer.Buffer)
				continue
			}
		}

		// 解析普通命令
		statement := &Statement{}
		switch PrepareStatement(inputBuffer, statement) {
		case PrepareSuccess:
		case PrepareSyntaxError:
			fmt.Printf("Syntax error. Could not parse statement.\n")
			continue
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", inputBuffer.Buffer)
			continue
		}

		// 执行
		switch ExecuteStatement(statement, t) {
		case ExecuteSuccess:
			fmt.Println("Executed.")
		case ExecuteTableFull:
			fmt.Println("Error: Table full.")
		}
	}
}
