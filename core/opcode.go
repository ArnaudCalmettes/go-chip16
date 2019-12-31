package core

import "encoding/binary"

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

// Returns the word argument
//
//		Op YX LL HH
//		      ^^^^^
func (o Opcode) HHLL() uint16 {
	return uint16(o&0xFF<<8 | o&0xFF00>>8)
}

// Returns the X register argument
//
//		Op YX LL HH
//		    ^
func (o Opcode) X() uint8 {
	return uint8(o & 0x0F0000 >> 16)
}

// Returns the Y register argument
//
//		Op YX LL HH
//		   ^
func (o Opcode) Y() uint8 {
	return uint8(o & 0xF00000 >> 20)
}

// Returns the Z register argument
//
//		Op YX 0Z 00
//		       ^
func (o Opcode) Z() uint8 {
	return uint8(o & 0x0F00 >> 8)
}
