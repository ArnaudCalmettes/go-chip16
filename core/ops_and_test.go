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

// ANDI Rx, HHLL
func TestAndiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range andTestCases {
		v.Regs[3] = test.x
		op := Opcode(0x60030000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[3], v.Flags, "&")
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
			checkOpResults(a, &test, v.Regs[2], v.Flags, "&")
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
			checkOpResults(a, &test, v.Regs[6], v.Flags, "&")
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
		op := Opcode(0x63030000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpFlags(a, &test, v.Flags, "TST")
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
			checkOpFlags(a, &test, v.Flags, "TST")
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
