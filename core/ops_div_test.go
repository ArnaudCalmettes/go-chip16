package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var divTestCases = []arithTestCase{
	{4, 2, 2, false, false, false, false},
	{0, 2, 0, false, false, false, true},
	{10, 3, 3, true, false, false, false},
	{-10, 3, -3, true, false, true, false},
}

var modTestCases = []arithTestCase{
	{0, 2, 0, false, false, false, true},
	{5, 3, 2, false, false, false, false},
	{5, -3, -1, false, false, true, false},
	{-5, 3, 1, false, false, false, false},
	{-5, -3, -2, false, false, true, false},
}

var remTestCases = []arithTestCase{
	{0, 2, 0, false, false, false, true},
	{5, 3, 2, false, false, false, false},
	{5, -3, 2, false, false, false, false},
	{-5, 3, -2, false, false, true, false},
	{-5, -3, -2, false, false, true, false},
}

// DIVI Rx, HHLL

func TestDiviRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA0000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range divTestCases {
		v.Regs[2] = test.x
		o := Opcode(0xA0020000).WithHHLL(uint16(test.y))
		if a.NoError(v.Eval(o)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "/")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkDiviRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		// Avoid division by 0 (n|1)
		if err := diviRxHHLL(v, Opcode(n|0x0100)); err != nil {
			b.Fatal(err)
		}
	}
}

// DIV Rx, Ry

func TestDivRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA1000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range divTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xA1420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "/")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkDivRxRy(b *testing.B) {
	v := NewVirtualMachine()

	// Set R0 -> R15 to 1 -> 16 to avoid division by zero
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i + 1)
	}

	for n := 0; n < b.N; n++ {
		// Write results only in R0
		if err := divRxRy(v, Opcode(n&0xFFF0FFFF)); err != nil {
			b.Fatal(err)
		}
		// Reset R0 to 1
		v.Regs[0] = 1
	}
}

// DIV Rx, Ry, Rz

func TestDivRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA2000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range divTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xA2420100))) {
			checkOpResults(a, &test, v.Regs[1], v.Flags, "/")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkDivRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	// Set R0 -> R15 to 1 -> 16 to avoid division by zero
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i + 1)
	}

	for n := 0; n < b.N; n++ {
		// Write results only in R0
		if err := divRxRyRz(v, Opcode(n&0xFFFFF0FF)); err != nil {
			b.Fatal(err)
		}
		// Reset R0 to 1
		v.Regs[0] = 1
	}
}

// MODI Rx, HHLL

func TestModiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA3000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range modTestCases {
		v.Regs[2] = test.x
		op := Opcode(0xA3020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "MOD")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkModiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		// Avoid division by 0 (n|1)
		if err := modiRxHHLL(v, Opcode(n|0x0100)); err != nil {
			b.Fatal(err)
		}
	}
}

// MOD Rx, Ry

func TestModRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA4000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range modTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xA4420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "MOD")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkModRxRy(b *testing.B) {
	v := NewVirtualMachine()

	// Set R0 -> R15 to 1 -> 16 to avoid division by zero
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i + 1)
	}

	for n := 0; n < b.N; n++ {
		// Write results only in R0
		if err := modRxRy(v, Opcode(n&0xFFF0FFFF)); err != nil {
			b.Fatal(err)
		}
		// Reset R0 to 1
		v.Regs[0] = 1
	}
}

// MOD Rx, Ry, Rz

func TestModRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA5000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range modTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xA5420100))) {
			checkOpResults(a, &test, v.Regs[1], v.Flags, "MOD")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkModRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	// Set R0 -> R15 to 1 -> 16 to avoid division by zero
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i + 1)
	}

	for n := 0; n < b.N; n++ {
		// Write results only in R0
		if err := modRxRyRz(v, Opcode(n&0xFFFFF0FF)); err != nil {
			b.Fatal(err)
		}
		// Reset R0 to 1
		v.Regs[0] = 1
	}
}

// REMI Rx, HHLL

func TestRemiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA6000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range remTestCases {
		v.Regs[2] = test.x
		op := Opcode(0xA6020000).WithHHLL(uint16(test.y))

		if a.NoError(v.Eval(op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "%")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkRemiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		// Avoid division by 0 (n|1)
		if err := remiRxHHLL(v, Opcode(n|0x0100)); err != nil {
			b.Fatal(err)
		}
	}
}

// REM Rx, Ry

func TestRemRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA7000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range remTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xA7420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "%")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkRemRxRy(b *testing.B) {
	v := NewVirtualMachine()

	// Set R0 -> R15 to 1 -> 16 to avoid division by zero
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i + 1)
	}

	for n := 0; n < b.N; n++ {
		// Write results only in R0
		if err := remRxRy(v, Opcode(n&0xFFF0FFFF)); err != nil {
			b.Fatal(err)
		}
		// Reset R0 to 1
		v.Regs[0] = 1
	}
}

// REM Rx, Ry, Rz

func TestRemRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	a.Errorf(
		v.Eval(Opcode(0xA8000000)),
		"Division by zero didn't return an error",
	)

	for _, test := range remTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(v.Eval(Opcode(0xA8420100))) {
			checkOpResults(a, &test, v.Regs[1], v.Flags, "%")
			a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkRemRxRyRz(b *testing.B) {
	v := NewVirtualMachine()

	// Set R0 -> R15 to 1 -> 16 to avoid division by zero
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i + 1)
	}

	for n := 0; n < b.N; n++ {
		// Write results only in R0
		if err := remRxRyRz(v, Opcode(n&0xFFFFF0FF)); err != nil {
			b.Fatal(err)
		}
		// Reset R0 to 1
		v.Regs[0] = 1
	}
}
