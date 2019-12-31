package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// LDI Rx HHLL

func TestLdiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	if a.NoError(v.Eval(Opcode(0x20053713))) {
		a.NoError(v.Check())
		a.Equalf(int16(0x1337), v.Regs[5], "Didn't assign 0x1337 to R5")
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkLdiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := ldiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// LDI SP HHLL

func TestLdiSPHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	if a.NoError(v.Eval(Opcode(0x2105DEFE))) {
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(0xFEDE), v.SP, "Didn't assign 0xFEDE to SP")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkLdiSPHHLL(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := ldiSPHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// LDM Rx, HHLL

func TestLdmRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.PutInt16At(0x1234, Pointer(0x1337))

	if a.NoError(v.Eval(Opcode(0x22053713))) {
		a.Equalf(int16(0x1234), v.Regs[5], "Didn't store expected value")
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		v.Eval(Opcode(0x2205FFFF)),
		"Access out of bounds should yield an error",
	)
}

func BenchmarkLdmRxHHLL(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		// Set HHLL's MSB to 0 to avoid reading out of memory boundaries
		if err := ldmRxHHLL(v, Opcode(n&0xFFFFFFEF)); err != nil {
			b.Fatal(err)
		}
	}
}

// LDM Rx, Ry

func TestLdmRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[2] = 0x1337
	v.PutInt16At(0x1234, Pointer(0x1337))

	if a.NoError(v.Eval(Opcode(0x23250000))) {
		a.NoError(v.Check())
		a.Equalf(int16(0x1234), v.Regs[5], "Didn't store expected value")
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	addr := uint16(0xFFFF)
	v.Regs[2] = int16(addr)
	a.Errorf(
		v.Eval(Opcode(0x23250000)),
		"Access out of bounds should yield an error",
	)
}

func BenchmarkLdmRxRy(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := ldmRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// MOV Rx, Ry

func TestMovRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[2] = 0x1337

	if a.NoError(v.Eval(Opcode(0x24250000))) {
		a.NoError(v.Check())
		a.Equalf(int16(0x1337), v.Regs[5], "Didn't store expeced value")
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkMovRxRy(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := mov(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
