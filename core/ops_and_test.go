package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var andTestCases = []arithTestCase{
	// X Y EXP C O N Z
	{0, 1, 0, false, false, false, true},
	{1, 3, 1, false, false, false, false},
	{1, -2, 0, false, false, false, true},
	{-2, -3, -4, false, false, true, false},
}

func checkAndResults(
	a *assert.Assertions,
	res int16,
	flags CPUFlags,
	test *arithTestCase,
) {
	a.Equalf(test.exp, res, "%d & %d", test.x, test.y)
	checkAndFlags(a, flags, test)
}

func checkAndFlags(a *assert.Assertions, flags CPUFlags, test *arithTestCase) {
	a.Equalf(
		test.c, flags.Carry(),
		"(%d & %d) wrong carry flag", test.x, test.y,
	)
	a.Equalf(
		test.o, flags.Overflow(),
		"(%d & %d) wrong overflow flag", test.x, test.y,
	)
	a.Equalf(
		test.n, flags.Negative(),
		"(%d & %d) wrong negative flag", test.x, test.y,
	)
	a.Equalf(
		test.z, flags.Zero(),
		"(%d & %d) wrong zero flag", test.x, test.y,
	)
}

// ANDI Rx, HHLL
func TestAndiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range andTestCases {
		v.Regs[3] = test.x
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x60030000 | uint32(llhh)))) {
			checkAndResults(a, v.Regs[3], v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAndiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := andiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// AND Rx, Ry
func TestAndRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range andTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x61420000))) {
			checkAndResults(a, v.Regs[2], v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAndRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := andRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// AND Rx, Ry, Rz
func TestAndRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range andTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x62420600))) {
			checkAndResults(a, v.Regs[6], v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAndRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := andRxRyRz(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// TSTI Rx, HHLL

func TestTstiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range andTestCases {
		v.Regs[3] = test.x
		llhh := uint16(test.y)&0xFF<<8 | uint16(test.y)&0xFF00>>8

		if a.NoError(v.Eval(Opcode(0x63030000 | uint32(llhh)))) {
			checkAndFlags(a, v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkTstiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := tstiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// TST Rx, Ry

func TestTstRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range andTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0x64420000))) {
			checkAndFlags(a, v.Flags, &test)
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkTstRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := tstRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
