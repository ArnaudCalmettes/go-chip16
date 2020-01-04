package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

var mulTestCases = []arithTestCase{
	{1, 2, 2, false, false, false, false},
	{1, -1, -1, false, false, true, false},
	{30000, 2, -5536, true, false, true, false},
	{-30000, 2, 5536, true, false, false, false},
}

// MULI Rx, HHLL

func TestMuliRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		if a.NoError(Eval(v, vm.Opcode(0x90020000).WithHHLL(uint16(test.y)))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "*")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkMuliRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := muliRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// MUL Rx, Ry

func TestMulRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x91420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "*")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkMulRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := mulRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// MUL Rx, Ry, Rz

func TestMulRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range mulTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x92420500))) {
			checkOpResults(a, &test, v.Regs[5], v.Flags, "*")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkMulRxRyRz(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := mulRxRyRz(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
