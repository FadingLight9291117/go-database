package metaCommandResult

type MetaCommandResult byte

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)
