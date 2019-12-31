package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// STM RX, HHLL

func TestStmRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[5] = 0x1234

	if a.NoError(v.Eval(Opcode(0x3005EFBE))) {
		val, _ := v.Int16At(Pointer(0xBEEF))
		a.Equalf(int16(0x1234), val, "Didn't store the right value")
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		v.Eval(Opcode(0x3005FFFF)),
		"Out of bound memory access didn't return an error",
	)
}

func BenchmarkStmRxHHLL(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		// Set HHLL's MSB to 0 to avoir writing out of memory boundaries
		if err := stmRxHHLL(v, Opcode(n&0xFFFFFFEF)); err != nil {
			b.Fatal(err)
		}
	}
}

// STM Rx, Ry

func TestStmRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[5] = 0x1234
	v.Regs[2] = 0x5678

	if a.NoError(v.Eval(Opcode(0x31250000))) {
		val, _ := v.Int16At(Pointer(0x5678))
		a.Equalf(int16(0x1234), val, "Didn't store the right value")
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	addr := uint16(0xFFFF)
	v.Regs[2] = int16(addr)

	a.Errorf(
		v.Eval(Opcode(0x31250000)),
		"Out of bound memory access didn't return an error",
	)
}

func BenchmarkStmRxRy(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := stmRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
