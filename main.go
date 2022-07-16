package main

import (
	"bufio"
	"com.fadinglight/database/table"
	"com.fadinglight/database/table/row"
	"com.fadinglight/database/types/executeResult"
	"com.fadinglight/database/types/prepareResult"
	"com.fadinglight/database/types/statementType"
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

type Statement struct {
	statementType statementType.StatementType
	Row2Insert    *row.Row //only used by insert statement
}

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

func PrepareInsert(inputBuffer *InputBuffer, statement *Statement) prepareResult.PrepareResult {
	statement.statementType = statementType.StatementInsert

	words := strings.Split(inputBuffer.Buffer, " ")
	command := words[0]
	words = words[1:]

	if command != "insert" {
		return prepareResult.PrepareUnrecognizedStatement
	}
	if len(words) != 3 {
		return prepareResult.PrepareSyntaxError
	}

	id, err := strconv.Atoi(words[0])
	if err != nil {
		return prepareResult.PrepareSyntaxError
	}
	username := words[1]
	email := words[2]
	if id < 0 {
		return prepareResult.PrepareNegativeId
	}
	if len(words[1]) > row.ColumnUsernameSize {
		return prepareResult.PrepareStringTooLong
	}
	if len(words[2]) > row.ColumnEmail {
		return prepareResult.PrepareStringTooLong
	}

	r := &row.Row{}
	r.Id = uint32(id)
	copy(r.Username[:], []rune(username))
	copy(r.Email[:], []rune(email))

	statement.Row2Insert = r

	return prepareResult.PrepareSuccess
}

func PrepareSelect(inputBuffer *InputBuffer, statement *Statement) prepareResult.PrepareResult {
	if inputBuffer.Buffer != "select" {
		return prepareResult.PrepareUnrecognizedStatement
	}

	statement.statementType = statementType.StatementSelect

	return prepareResult.PrepareSuccess
}

// PrepareStatement SQL Compiler
func PrepareStatement(inputBuffer *InputBuffer, statement *Statement) prepareResult.PrepareResult {
	if strings.HasPrefix(inputBuffer.Buffer, "insert") {
		return PrepareInsert(inputBuffer, statement)
	}
	if strings.HasPrefix(inputBuffer.Buffer, "select") {
		return PrepareSelect(inputBuffer, statement)
	}

	return prepareResult.PrepareUnrecognizedStatement
}

func ExecuteInsert(statement *Statement, table *table.Table) executeResult.ExecuteResult {
	if table.IsFull() {
		return executeResult.ExecuteTableFull
	}
	row2Insert := statement.Row2Insert
	table.Append(row2Insert)

	return executeResult.ExecuteSuccess
}

func ExecuteSelect(statement *Statement, table *table.Table) executeResult.ExecuteResult {
	for _, v := range table.Rows {
		v.Print()
	}
	return executeResult.ExecuteSuccess
}

func ExecuteStatement(statement *Statement, table *table.Table) executeResult.ExecuteResult {
	switch statement.statementType {
	case statementType.StatementInsert:
		fmt.Println("This is where we would do an insert.")
		return ExecuteInsert(statement, table)
	case statementType.StatementSelect:
		fmt.Println("This is where we would do a select.")
		return ExecuteSelect(statement, table)
	default:
		return executeResult.ExecuteError
	}
}

func main() {
	t := table.New()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 提示信息
		printPrompt()
		// 读取输入
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
		case prepareResult.PrepareSuccess:
		case prepareResult.PrepareSyntaxError:
			fmt.Println("Syntax error. Could not parse statement.")
			continue
		case prepareResult.PrepareUnrecognizedStatement:
			fmt.Println("Unrecognized keyword.")
			continue
		case prepareResult.PrepareStringTooLong:
			fmt.Println("String is too long.")
			continue
		case prepareResult.PrepareNegativeId:
			fmt.Println("ID must be positive.")
			continue
		}
		// 执行
		switch ExecuteStatement(statement, t) {
		case executeResult.ExecuteSuccess:
			fmt.Println("Executed.")
		case executeResult.ExecuteTableFull:
			fmt.Println("Error: Table full.")
		}
	}
}
