package executor

type ExecuteResult = byte

const (
	EXECUTE_SUCESS ExecuteResult = iota
	EXECUTE_TABLE_FULL
	EXECUTE_ERROR
)
