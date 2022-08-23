package processor

type PrepareResult byte

const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_SYNTAX_ERROR
	PREPARE_UNRECOGNIZED_STATEMENT
	PREPARE_STRING_TOOLONG
	PREPARE_NEGATIVE_ID
)
