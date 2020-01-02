package core

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var shlTestCases = []arithTestCase{
	{1, 2, 4, false, false, false, false},
	{0, 4, 0, false, false, false, true},
	{0x7FFF, 2, -4, false, false, true, false},
}

var shrTestCases = []arithTestCase{
	{4, 2, 1, false, false, false, false},
	{2, 4, 0, false, false, false, true},
	{math.MinInt16, 15, 1, false, false, false, false},
}

var sarTestCases = []arithTestCase{
	{4, 2, 1, false, false, false, false},
	{2, 4, 0, false, false, false, true},
	{math.MinInt16, 15, -1, false, false, true, false},
}

// SHL Rx, N

func TestShlRxN(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range shlTestCases {
		v.Regs[2] = test.x
		op := Opcode(0xB0020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "<<")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkShlRxN(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := shlRxN(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SHR Rx, N

func TestShrRxN(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range shrTestCases {
		v.Regs[2] = test.x
		op := Opcode(0xB1020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, ">>")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")

		}
	}
}

func BenchmarkShrRxN(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := shrRxN(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SAR Rx, N

func TestSarRxN(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range sarTestCases {
		v.Regs[2] = test.x
		op := Opcode(0xB2020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, ">>")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")

		}
	}
}

func BenchmarkSarRxN(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := sarRxN(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SHL Rx, Ry

func TestShlRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range shlTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xB3420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "<<")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkShlRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := shlRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SHR Rx, Ry

func TestShrRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range shrTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xB4420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, ">>")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkShrRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := shrRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SAR Rx, Ry
func TestSarRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range sarTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xB5420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, ">>")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSarRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := sarRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
