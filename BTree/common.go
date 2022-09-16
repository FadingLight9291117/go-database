package BTree

import "unsafe"

type NodeType = uint8

const (
	NODE_TYPE_LEAF NodeType = iota
	NODE_TYPE_INTERNAL
)

type CommonNodeHeader struct {
	Type      NodeType
	IsRoot    bool
	ParentPtr uint64 // 无法序列化
}

/**
 * CommonNode Header Layout (10 byte)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 byte)
 */
const (
	NODE_TYPE_SIZE          = int(unsafe.Sizeof(NODE_TYPE_INTERNAL))
	NODE_TYPE_OFFSET        = int(unsafe.Sizeof(NodeType(0)))
	NODE_IS_ROOT_SIZE       = int(unsafe.Sizeof(true))
	NODE_IS_ROOT_OFFSET     = int(NODE_TYPE_OFFSET + NODE_TYPE_SIZE)
	NODE_PARENTPTR_SIZE     = int(unsafe.Sizeof(uint64(0)))
	NODE_PARENTPTR_OFFSET   = int(NODE_IS_ROOT_OFFSET + NODE_IS_ROOT_SIZE)
	COMMON_NODE_HEADER_SIZE = int(NODE_TYPE_SIZE + NODE_IS_ROOT_SIZE + NODE_PARENTPTR_SIZE)
)
