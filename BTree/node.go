package BTree

type Node interface {
	Serialize() ([]byte, error)
	Deserialize(b []byte) error
	IsRootNode() bool
	SetRootNode(isRoot bool)
	GetMaxKey() uint64
}

// CreateLeafNode 将node的cellNum置零
func CreateLeafNode() *LeafNode {
	node := &LeafNode{}
	node.Type = NODE_TYPE_LEAF
	node.IsRoot = false
	node.CellNums = 0
	return node
}

func CreateInternalNode() *InternalNode {
	node := &InternalNode{}
	node.Type = NODE_TYPE_INTERNAL
	node.IsRoot = false
	node.CellNums = 0
	return node
}
