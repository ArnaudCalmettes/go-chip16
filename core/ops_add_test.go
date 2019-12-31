package core

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var addTestCases = []arithTestCase{
	{1, 2, 3, false, false, false, false},
	{1, -1, 0, true, false, false, true},
	{1, -2, -1, false, false, true, false},
	{math.MaxInt16, 1, math.MinInt16, false, true, true, false},
	{math.MinInt16, -1, math.MaxInt16, true, true, false, false},
	{math.MaxInt16, math.MaxInt16, -2, false, true, true, false},
}

func checkAddResults(
	a *assert.Assertions,
	res int16,
	flags CPUFlags,
	test *arithTestCase,
) {
	a.Equalf(
		test.exp, res, "%d + %d != %d",
		test.x, test.y, test.exp,
	)
	a.Equalf(
		test.c, flags.Carry(),
		"(%d + %d) wrong carry flag", test.x, test.y,
	)
	a.Equalf(
		test.o, flags.Overflow(),
		"(%d + %d) wrong overflow flag", test.x, test.y,
	)
	a.Equalf(
		test.n, flags.Negative(),
		"(%d + %d) wrong negative flag", test.x, test.y,
	)
	a.Equalf(
		test.z, flags.Zero(),
		"(%d + %d) wrong zero flag", test.x, test.y,
	)
}

// ADDI Rx, HHLL

func TestAddiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range addTestCases {
		v.Regs[2] = test.x
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x40020000 | uint32(llhh)))) {
			checkAddResults(a, v.Regs[2], v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAddiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := addiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// ADD Rx, Ry

func TestAddRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range addTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x41420000))) {
			checkAddResults(a, v.Regs[2], v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAddRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := addRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// ADD Rx, Ry, Rz

func TestAddRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range addTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x42420500))) {
			checkAddResults(a, v.Regs[5], v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAddRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := addRxRyRz(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
