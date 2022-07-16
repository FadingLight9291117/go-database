package prepareResult

type PrepareResult byte

const (
	PrepareSuccess PrepareResult = iota
	PrepareSyntaxError
	PrepareUnrecognizedStatement
	PrepareStringTooLong
	PrepareNegativeId
)
