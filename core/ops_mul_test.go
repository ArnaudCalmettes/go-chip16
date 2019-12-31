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

func checkMulResults(
	a *assert.Assertions,
	res int16,
	flags CPUFlags,
	test *arithTestCase,
) {
	a.Equalf(
		test.exp, res, "%d * %d != %d",
		test.x, test.y, test.exp,
	)
	a.Equalf(
		test.c, flags.Carry(),
		"(%d * %d) wrong carry flag", test.x, test.y,
	)
	a.Equalf(
		test.n, flags.Negative(),
		"(%d * %d) wrong negative flag", test.x, test.y,
	)
	a.Equalf(
		test.z, flags.Zero(),
		"(%d * %d) wrong zero flag", test.x, test.y,
	)
}

// MULI Rx, HHLL

func TestMuliRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x90020000 | uint32(llhh)))) {
			checkMulResults(a, v.Regs[2], v.Flags, &test)
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
			checkMulResults(a, v.Regs[2], v.Flags, &test)
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
			checkMulResults(a, v.Regs[5], v.Flags, &test)
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
