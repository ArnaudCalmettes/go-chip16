package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var orTestCases = []arithTestCase{
	// X Y EXP C O N Z
	{0, 1, 1, false, false, false, false},
	{0, 0, 0, false, false, false, true},
	{-1, 1, -1, false, false, true, false},
	{-3, 2, -1, false, false, true, false},
}

// ORI Rx, HHLL
func TestOriRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range orTestCases {
		v.Regs[3] = test.x
		op := Opcode(0x70030000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[3], v.Flags, "|")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkOriRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := oriRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// OR Rx, Ry
func TestOrRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range orTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x71420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "|")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkOrRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := orRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// OR Rx, Ry, Rz
func TestOrRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range orTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x72420600))) {
			checkOpResults(a, &test, v.Regs[6], v.Flags, "|")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkOrRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := orRxRyRz(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
