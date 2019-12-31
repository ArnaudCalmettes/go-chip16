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

func checkSubFlags(a *assert.Assertions, flags CPUFlags, test *arithTestCase) {
	a.Equalf(
		test.c, flags.Carry(),
		"(%d - %d) wrong carry flag", test.x, test.y,
	)
	a.Equalf(
		test.o, flags.Overflow(),
		"(%d - %d) wrong overflow flag", test.x, test.y,
	)
	a.Equalf(
		test.n, flags.Negative(),
		"(%d - %d) wrong negative flag", test.x, test.y,
	)
	a.Equalf(
		test.z, flags.Zero(),
		"(%d - %d) wrong zero flag", test.x, test.y,
	)
}

// SUBI Rx, HHLL

func TestSubiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x50020000 | uint32(llhh)))) {
			a.Equalf(
				test.exp, v.Regs[2], "%d - %d != %d",
				test.x, test.y, test.exp,
			)
			checkSubFlags(a, v.Flags, &test)
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
			a.Equalf(
				test.exp, v.Regs[2], "%d - %d != %d",
				test.x, test.y, test.exp,
			)
			checkSubFlags(a, v.Flags, &test)
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
			a.Equalf(
				test.exp, v.Regs[5], "%d - %d != %d",
				test.x, test.y, test.exp,
			)
			checkSubFlags(a, v.Flags, &test)
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
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x53020000 | uint32(llhh)))) {
			checkSubFlags(a, v.Flags, &test)
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
			checkSubFlags(a, v.Flags, &test)
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
