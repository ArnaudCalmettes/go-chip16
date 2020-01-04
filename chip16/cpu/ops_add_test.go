package cpu

import (
	"math"
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
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

// ADDI Rx, HHLL

func TestAddiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range addTestCases {
		v.Regs[2] = test.x
		op := vm.Opcode(0x40020000).WithHHLL(uint16(test.y))

		if a.NoError(Eval(v, op)) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "+")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAddiRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := addiRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// ADD Rx, Ry

func TestAddRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range addTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x41420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "+")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAddRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := addRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// ADD Rx, Ry, Rz

func TestAddRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range addTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x42420500))) {
			checkOpResults(a, &test, v.Regs[5], v.Flags, "+")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkAddRxRyRz(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := addRxRyRz(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
