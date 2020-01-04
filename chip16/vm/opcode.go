package vm

import "encoding/binary"

// Opcode is our internal representation of an opcode.
type Opcode uint32

// Reads an Opcode from a byte slice
func readOpcode(s []byte) Opcode {
	return Opcode(binary.BigEndian.Uint32(s))
}

// Op returns the operation designated by the Opcode (leading byte)
//
//		Op YX LL HH
//		^^
func (o Opcode) Op() int {
	return int(o & 0xFF000000 >> 24)
}

// HHLL returns the word argument
//
//		Op YX LL HH
//		      ^^^^^
func (o Opcode) HHLL() uint16 {
	return uint16(o&0xFF<<8 | o&0xFF00>>8)
}

// X returns the X register argument
//
//		Op YX LL HH
//		    ^
func (o Opcode) X() uint8 {
	return uint8(o & 0x0F0000 >> 16)
}

// Y returns the Y register argument
//
//		Op YX LL HH
//		   ^
func (o Opcode) Y() uint8 {
	return uint8(o & 0xF00000 >> 20)
}

// Z returns the Z register argument
//
//		Op YX 0Z 00
//		       ^
func (o Opcode) Z() uint8 {
	return uint8(o & 0x0F00 >> 8)
}

// N returns the N nibble argument
//
//		Op YX 0N 00
//		       ^
func (o Opcode) N() uint8 {
	return uint8(o & 0x0F00 >> 8)
}

// WithHHLL returns a copy of the opcode with hhll set to given value
func (o Opcode) WithHHLL(hhll uint16) Opcode {
	return Opcode(
		uint32(o&0xFFFF0000) | uint32(hhll&0xFF00>>8|hhll&0x00FF<<8),
	)
}
