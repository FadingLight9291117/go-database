package BTree

type Node interface {
	Serialize() (byte, error)
	Deserialize() error
	IsRoot() bool
}
