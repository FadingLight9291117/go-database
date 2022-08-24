package executor

import (
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"com.fadinglight/db/types/metaCommandResult"
	"errors"
	"fmt"
	"os"
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

func DoMetaCommand(inputBuffer *types.InputBuffer, t *table.Table) (metaCommandResult.MetaCommandResult, error) {
	if inputBuffer.Buffer == ".exit" {
		err := t.Close()
		if err != nil {
			return 0, err
		}
		os.Exit(0)
		return metaCommandResult.MetaCommandSuccess, nil
	} else {
		return metaCommandResult.MetaCommandUnrecognizedCommand, errors.New(fmt.Sprintf("Unrecognized command '%s' .\n", inputBuffer.Buffer))
	}
}
