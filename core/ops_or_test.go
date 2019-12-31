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

func checkOrResults(
	a *assert.Assertions,
	res int16,
	flags CPUFlags,
	test *arithTestCase,
) {
	a.Equalf(test.exp, res, "%d | %d", test.x, test.y)
	checkOrFlags(a, flags, test)
}

func checkOrFlags(a *assert.Assertions, flags CPUFlags, test *arithTestCase) {
	a.Equalf(
		test.c, flags.Carry(),
		"(%d | %d) wrong carry flag", test.x, test.y,
	)
	a.Equalf(
		test.o, flags.Overflow(),
		"(%d | %d) wrong overflow flag", test.x, test.y,
	)
	a.Equalf(
		test.n, flags.Negative(),
		"(%d | %d) wrong negative flag", test.x, test.y,
	)
	a.Equalf(
		test.z, flags.Zero(),
		"(%d | %d) wrong zero flag", test.x, test.y,
	)
}

// ORI Rx, HHLL
func TestOriRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range orTestCases {
		v.Regs[3] = test.x
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x70030000 | uint32(llhh)))) {
			checkOrResults(a, v.Regs[3], v.Flags, &test)
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
			checkOrResults(a, v.Regs[2], v.Flags, &test)
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
			checkOrResults(a, v.Regs[6], v.Flags, &test)
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
