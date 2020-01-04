package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

// STM RX, HHLL

func TestStmRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.Regs[5] = 0x1234

	if a.NoError(Eval(v, vm.Opcode(0x3005EFBE))) {
		val, _ := v.Int16At(vm.Pointer(0xBEEF))
		a.Equalf(int16(0x1234), val, "Didn't store the right value")
		a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		Eval(v, vm.Opcode(0x3005FFFF)),
		"Out of bound memory access didn't return an error",
	)
}

func BenchmarkStmRxHHLL(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		// Set HHLL's MSB to 0 to avoir writing out of memory boundaries
		if err := stmRxHHLL(v, vm.Opcode(n&0xFFFFFFEF)); err != nil {
			b.Fatal(err)
		}
	}
}

// STM Rx, Ry

func TestStmRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.Regs[5] = 0x1234
	v.Regs[2] = 0x5678

	if a.NoError(Eval(v, vm.Opcode(0x31250000))) {
		val, _ := v.Int16At(vm.Pointer(0x5678))
		a.Equalf(int16(0x1234), val, "Didn't store the right value")
		a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	addr := uint16(0xFFFF)
	v.Regs[2] = int16(addr)

	a.Errorf(
		Eval(v, vm.Opcode(0x31250000)),
		"Out of bound memory access didn't return an error",
	)
}

func BenchmarkStmRxRy(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		if err := stmRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
