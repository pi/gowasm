package wast

import "github.com/pi/gowasm/wasm"

// NodeKind defines kind of node
type NodeKind int

const (
	NkNone = NodeKind(iota)
	// logical kinds
	NkModule
	NkFunc
	NkMemory
	// opcode kinds
	NkUnary
	NkBinary
	NkBlock
	NkBr
	NkBrIf
	NkBrTable
	NkCall
	NkCallImport
	NkCallIndirect
	NkCompare
	NkConst
	NkConvert
	NkCurrentMemory
	NkGrowMemory
	NkGetLocal
	NkSetLocal
	NkIf
	NkIfElse
	NkLoop
	NkNop
	NkReturn
	NkSelect
	NkLoad
	NkStore
	NkUnreachable
)

type NodeVisitor interface {
	VisitDefault(n Node)

	VisitModule(n *ModuleNode)
	// opcode kinds
	VisitUnary(n *UnaryNode)
	VisitBinary(n *BinaryNode)
	VisitBlock(n *BlockNode)
	VisitBr(n *BrNode)
	VisitBrIf(n *BrIfNode)
	VisitBrTable(n *BrTableNode)
	VisitCall(n *CallNode)
	VisitCallImport(n *CallImportNode)
	VisitCallIndirect(n *CallIndirectNode)
	VisitCompare(n *CompareNode)
	VisitConst(n *ConstNode)
	VisitConvert(n *ConvertNode)
	VisitCurrentMemory(n *CurrentMemoryNode)
	VisitGetLocal(n *GetLocalNode)
	VisitSetLocal(n *SetLocalNode)
	VisitGrowMemory(n *GrowMemoryNode)
	VisitIf(n *IfNode)
	VisitIfElse(n *IfElseNode)
	VisitLoad(n *LoadNode)
	VisitLoop(n *LoopNode)
	VisitNop(n *NopNode)
	VisitReturn(n *ReturnNode)
	VisitSelect(n *SelectNode)
	VisitStore(n *StoreNode)
	VisitUnreachable(n *UnreachableNode)
}

type Node interface {
	Opcode() wasm.Opcode
	Kind() NodeKind
	ValueType() wasm.ValueType

	SubNodes() []Node
	Visit(v NodeVisitor)
}

func subNodeIndexOOB(index int) {
	panic("subnode index out of bounds")
}
func unimp(funcName string) {
	panic(funcName + " is not implemented")
}

type defaultNodeImplementation struct {
	subnodes []Node
}

func (n *defaultNodeImplementation) SubNodes() []Node {
	return n.subnodes
}
func (n *defaultNodeImplementation) Opcode() wasm.Opcode {
	return wasm.Op_None
}
func (n *defaultNodeImplementation) Kind() NodeKind {
	return NkNone
}
func (n *defaultNodeImplementation) Visit(v NodeVisitor) {
	unimp("Visit")
}
func (n *defaultNodeImplementation) ValueType() wasm.ValueType {
	return wasm.Vt_none
}

type nodeWithOpcode struct {
	defaultNodeImplementation
	opcode wasm.Opcode
}

func (n *nodeWithOpcode) Opcode() wasm.Opcode {
	return n.opcode
}

type UnaryNode struct {
	nodeWithOpcode
}

func (n *UnaryNode) Kind() NodeKind {
	return NkUnary
}

type BinaryNode struct {
	nodeWithOpcode
}

func (n *BinaryNode) Kind() NodeKind {
	return NkBinary
}
