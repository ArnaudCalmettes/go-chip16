package vm

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

	a.Equal(0x12, o.Op(), "Op() failed")

	a.Equal(uint16(0x7856), o.HHLL(), "HHLL() failed")
	a.Equal(uint8(0x56), o.LL(), "LL() failed")
	a.Equal(uint8(0x78), o.HH(), "HH() failed")
	a.Equal(uint8(0x4), o.X(), "X() failed")
	a.Equal(uint8(0x3), o.Y(), "Y() failed")
	a.Equal(uint8(0x6), o.Z(), "Z() failed")
	a.Equal(uint8(0x6), o.N(), "N() failed")

	a.Equal(Opcode(0x12340123), o.WithHHLL(0x2301), "o.WithHHLL() failed")
}
