package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var xorTestCases = []arithTestCase{
	// X Y EXP C O N Z
	{0, 1, 1, false, false, false, false},
	{0, 0, 0, false, false, false, true},
	{-1, 1, -2, false, false, true, false},
	{-3, 2, -1, false, false, true, false},
}

// XORI Rx, HHLL
func TestXoriRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range xorTestCases {
		v.Regs[3] = test.x
		op := Opcode(0x80030000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[3], v.Flags, "^")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkXoriRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := oriRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// XOR Rx, Ry
func TestXorRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range xorTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x81420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "^")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkXorRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := xorRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// XOR Rx, Ry, Rz
func TestXorRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range xorTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x82420600))) {
			checkOpResults(a, &test, v.Regs[6], v.Flags, "^")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkXorRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := xorRxRyRz(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
