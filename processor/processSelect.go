package processor

import (
	"com.fadinglight/db/types"
	"errors"
)

func PrepareSelect(inputBuffer *types.InputBuffer, statement *types.Statement) (PrepareResult, error) {
	if inputBuffer.Buffer != "select" {
		return PREPARE_UNRECOGNIZED_STATEMENT, errors.New("unrecognized keyword")
	}
	statement.StatementType = types.StatementSelect

	return PREPARE_SUCCESS, nil
}
