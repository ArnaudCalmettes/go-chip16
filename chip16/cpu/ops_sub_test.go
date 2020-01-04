package cpu

import (
	"math"
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
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
	v := vm.NewState()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		op := vm.Opcode(0x50020000).WithHHLL(uint16(test.y))

		if a.NoError(Eval(v, op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "-")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSubiRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := subiRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SUB Rx, Ry

func TestSubRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x51420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "-")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSubRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := subRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// SUB Rx, Ry, Rz

func TestSubRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x52420500))) {
			checkOpResults(a, &test, v.Regs[5], v.Flags, "-")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkSubRxRyRz(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := subRxRyRz(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CMPI Rx, HHLL

func TestCmpiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		op := vm.Opcode(0x53020000).WithHHLL(uint16(test.y))

		if a.NoError(Eval(v, op)) {
			checkOpFlags(a, &test, v.Flags, "CMP")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkCmpiRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := cmpiRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CMP Rx, Ry

func TestCmpRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range subTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x54420000))) {
			checkOpFlags(a, &test, v.Flags, "CMP")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkCmpRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := cmpRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
