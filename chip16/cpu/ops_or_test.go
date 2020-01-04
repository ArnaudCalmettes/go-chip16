package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

var orTestCases = []arithTestCase{
	// X Y EXP C O N Z
	{0, 1, 1, false, false, false, false},
	{0, 0, 0, false, false, false, true},
	{-1, 1, -1, false, false, true, false},
	{-3, 2, -1, false, false, true, false},
}

// ORI Rx, HHLL
func TestOriRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range orTestCases {
		v.Regs[3] = test.x
		op := vm.Opcode(0x70030000).WithHHLL(uint16(test.y))

		if a.NoError(Eval(v, op)) {
			checkOpResults(a, &test, v.Regs[3], v.Flags, "|")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkOriRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := oriRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// OR Rx, Ry
func TestOrRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range orTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x71420000))) {
			checkOpResults(a, &test, v.Regs[2], v.Flags, "|")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkOrRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := orRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// OR Rx, Ry, Rz
func TestOrRxRyRz(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range orTestCases {
		v.Regs[2] = test.x
		v.Regs[4] = test.y

		if a.NoError(Eval(v, vm.Opcode(0x72420600))) {
			checkOpResults(a, &test, v.Regs[6], v.Flags, "|")
			a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
			a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		}
	}
}

func BenchmarkOrRxRyRz(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := orRxRyRz(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
