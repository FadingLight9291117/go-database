package processor

import (
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"errors"
	"strconv"
	"strings"
)

func PrepareInsert(inputBuffer *types.InputBuffer, statement *types.Statement) (PrepareResult, error) {
	statement.StatementType = types.StatementInsert

	words := strings.Split(inputBuffer.Buffer, " ")
	command := words[0]
	words = words[1:]

	if command != "insert" {
		return PREPARE_UNRECOGNIZED_STATEMENT, errors.New("unrecognized keyword")
	}
	if len(words) != 3 {
		return PREPARE_SYNTAX_ERROR, errors.New("syntax error. Could not parse statement")
	}

	id, err := strconv.Atoi(words[0])
	if err != nil {
		return PREPARE_SYNTAX_ERROR, errors.New("syntax error. Could not parse statement")
	}
	username := words[1]
	email := words[2]
	if id < 0 {
		return PREPARE_NEGATIVE_ID, errors.New("ID must be positive")
	}
	if len(words[1]) > table.ColumnUsernameSize {
		return PREPARE_STRING_TOOLONG, errors.New("string is too long")
	}
	if len(words[2]) > table.ColumnEmail {
		return PREPARE_STRING_TOOLONG, errors.New("string is too long")
	}

	r := table.NewRow(id, username, email)
	statement.Row2Insert = r

	return PREPARE_SUCCESS, nil
}
