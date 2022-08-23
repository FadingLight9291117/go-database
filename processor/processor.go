package processor

import (
	"com.fadinglight/db/types"
	"errors"
	"strings"
)

// PrepareStatement SQL Compiler
func PrepareStatement(inputBuffer *types.InputBuffer, statement *types.Statement) (PrepareResult, error) {
	if strings.HasPrefix(inputBuffer.Buffer, "insert") {
		return PrepareInsert(inputBuffer, statement)
	}
	if strings.HasPrefix(inputBuffer.Buffer, "select") {
		return PrepareSelect(inputBuffer, statement)
	}
	return PREPARE_UNRECOGNIZED_STATEMENT, errors.New("Unrecognized keyword.")
}
