package executor

import (
	"com.fadinglight/db/BTree"
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
	switch inputBuffer.Buffer {
	case ".exit":
		err := t.Close()
		if err != nil {
			return 0, err
		}
		os.Exit(0)
	case ".constants":
		fmt.Println("Constants:")
		printConstants()
	case ".btree":
		fmt.Println("Tree:")
		t.Pager.PrintTree(0, 0)
	default:
		return MetaCommandUnrecognizedCommand, errors.New(fmt.Sprintf("Unrecognized command '%s' .\n", inputBuffer.Buffer))
	}
	return MetaCommandSuccess, nil
}

func printConstants() {
	fmt.Printf("ROW_SIZE: %d\n", BTree.ROW_SIZE)
	fmt.Printf("COMMON_NODE_HEADER_SIZE: %d\n", BTree.COMMON_NODE_HEADER_SIZE)
	fmt.Printf("LEAF_NODE_HEADER_SIZE: %d\n", BTree.LEAF_NODE_HEADER_SIZE)
	fmt.Printf("LEAF_NODE_CELL_SIZE: %d\n", BTree.LEAF_NODE_CELL_SIZE)
	fmt.Printf("LEAF_NODE_SPACE_FOR_CELLS: %d\n", BTree.LEAF_NODE_SPACE_FOR_CELLS)
	fmt.Printf("LEAF_NODE_MAX_CELLS: %d\n", BTree.LEAF_NODE_MAX_CELLS)
	fmt.Printf("LEAF_NODE_SIZE: %d\n", BTree.LEAF_NODE_SIZE)
	fmt.Printf("LEAF_NODE_BLANK_SIZE: %d\n", BTree.LEAF_NODE_BLANK_SIZE)
	fmt.Printf("INTERNAL_NODE_HEADER_SIZE: %d\n.", BTree.INTERNAL_NODE_HEADER_SIZE)
	fmt.Printf("INTERNAL_NODE_CELL_SIZE: %d\n", BTree.INTERNAL_NODE_CELL_NUMS_SIZE)
	fmt.Printf("INTERNAL_NODE_SPACE_FOR_CELLS: %d\n", BTree.INTERNAL_NODE_SPACE_FOR_CELLS)
	fmt.Printf("INTERNAL_NODE_MAX_CELLS: %d\n", BTree.INTERNAL_NODE_MAX_CELLS)
	fmt.Printf("INTERNAL_NODE_SIZE: %d\n", BTree.INTERNAL_NODE_SIZE)
	fmt.Printf("INTERNAL_NODE_BLANK_SIZE: %d\n", BTree.INTERNAL_NODE_BLANK_SIZE)
}
