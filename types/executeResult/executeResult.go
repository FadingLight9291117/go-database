package executeResult

type ExecuteResult = byte

const (
	ExecuteSuccess ExecuteResult = iota
	ExecuteTableFull
	ExecuteError
)
