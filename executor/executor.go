package executor

import (
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"errors"
	"fmt"
	"os"
)

type MetaCommandResult byte

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

func ExecuteStatement(statement *types.Statement, table *table.Table) (ExecuteResult, error) {
	switch statement.StatementType {
	case types.StatementInsert:
		fmt.Println("This is where we would do an insert.")
		return ExecuteInsert(statement, table)
	case types.StatementSelect:
		fmt.Println("This is where we would do a select.")
		return ExecuteSelect(statement, table)
	default:
		return EXECUTE_ERROR, errors.New("unknown statement type")
	}
}

func DoMetaCommand(inputBuffer *types.InputBuffer, t *table.Table) (MetaCommandResult, error) {
	if inputBuffer.Buffer == ".exit" {
		err := t.Close()
		if err != nil {
			return 0, err
		}
		os.Exit(0)
		return MetaCommandSuccess, nil
	} else {
		return MetaCommandUnrecognizedCommand, errors.New(fmt.Sprintf("Unrecognized command '%s' .\n", inputBuffer.Buffer))
	}
}
