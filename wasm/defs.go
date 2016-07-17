package wasm

import (
	"log"
)

type ValueType int

const (
	VtNone = ValueType(-1)
	VtI32  = ValueType(1)
	VtI64  = ValueType(2)
	VtF32  = ValueType(3)
	VtF64  = ValueType(4)
)

func (v ValueType) String() string {
	switch v {
	case Vt_none:
		return "<none>"
	case Vt_i32:
		return "i32"
	case Vt_i64:
		return "i64"
	case Vt_f32:
		return "f32"
	case Vt_f64:
		return "f64"
	default:
		log.Panicf("unknown value type: %d", int(v))
	}
}

type OpcodeFlags int

const OpfNone = OpcodeFlags(0)
const (
	OpfBlock = OpcodeFlags(1 << iota)
	OpfI32
	OpfI64
	OpfF32
	OpfF64
)

type Opcode int

func (c Opcode) Mnemonic() string {
	if m, ok := opcodeMnemonics[c]; ok {
		return m
	}
	return fmt.Sprintf("<invalid opcode %#x>", c)
}

func (c Opcode) IsValid() bool {
	_, ok := opcodeFlags[c]
	return ok
}

func (c Opcode) Flags() OpcodeFlags {
	f, ok := opcodeFlags[c]
	if ok {
		return f
	}
	return 0
}
