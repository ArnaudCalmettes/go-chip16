package core

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var subTestCases = []arithTestCase{
	{3, 2, 1, false, false, false, false},
	{1, 1, 0, false, false, false, true},
	{1, 2, -1, true, false, true, false},
	{math.MinInt16, 1, math.MaxInt16, false, true, false, false},
	{math.MaxInt16, -1, math.MinInt16, true, true, true, false},
	{math.MinInt16, math.MinInt16, 0, false, false, false, true},
	{-3, -2, -1, true, false, true, false},
}

// SUBI Rx, HHLL

func TestSubiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		op := Opcode(0x50020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "-")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSubiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := subiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SUB Rx, Ry

func TestSubRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x51420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "-")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSubRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := subRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SUB Rx, Ry, Rz

func TestSubRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x52420500))) {
			checkOpResults(a, &test, v.Regs[5], v.Flags, "-")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSubRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := subRxRyRz(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CMPI Rx, HHLL

func TestCmpiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		op := Opcode(0x53020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpFlags(a, &test, v.Flags, "CMP")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkCmpiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := cmpiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CMP Rx, Ry

func TestCmpRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x54420000))) {
			checkOpFlags(a, &test, v.Flags, "CMP")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkCmpRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := cmpRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
