package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var o Opcode = 0x12345678

func TestOpcodeRead(t *testing.T) {
	a := assert.New(t)
	s := []byte{0x12, 0x34, 0x56, 0x78}
	a.Equal(o, readOpcode(s))
}

func TestOpcodeMethods(t *testing.T) {
	a := assert.New(t)

	a.Equalf(0x12, o.Op(), "Op() failed")

	a.Equalf(uint16(0x7856), o.HHLL(), "HHLL() failed")
	a.Equalf(uint8(0x4), o.X(), "X() failed")
	a.Equalf(uint8(0x3), o.Y(), "Y() failed")
	a.Equalf(uint8(0x6), o.Z(), "Z() failed")
	a.Equalf(uint8(0x6), o.N(), "N() failed")

	a.Equalf(Opcode(0x12340123), o.WithHHLL(0x2301), "o.WithHHLL() failed")
}
