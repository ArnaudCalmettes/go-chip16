package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mulTestCases = []arithTestCase{
	{1, 2, 2, false, false, false, false},
	{1, -1, -1, false, false, true, false},
	{30000, 2, -5536, true, false, true, false},
	{-30000, 2, 5536, true, false, false, false},
}

// MULI Rx, HHLL

func TestMuliRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		if a.NoError(v.Eval(Opcode(0x90020000).WithHHLL(uint16(test.y)))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "*")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkMuliRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := muliRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// MUL Rx, Ry

func TestMulRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x91420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "*")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkMulRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := mulRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// MUL Rx, Ry, Rz

func TestMulRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x92420500))) {
			checkOpResults(a, &test, v.Regs[5], v.Flags, "*")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkMulRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := mulRxRyRz(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
