package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// JMP HHLL

func TestJmpHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	if a.NoError(v.Eval(Opcode(0x10003713))) {
		a.Equalf(Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkJmpHHLL(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := jmpHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// JMC

func TestJmc(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	if a.NoError(v.Eval(Opcode(0x11003713))) {
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	v.Flags.SetCarry(true)

	if a.NoError(v.Eval(Opcode(0x11003713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Truef(v.Flags.Carry(), "Carry flag should be set")
	}
}

func BenchmarkJmcNoCarry(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := jmc(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJmcCarry(b *testing.B) {
	v := NewVirtualMachine()
	v.Flags.SetCarry(true)
	for n := 0; n < b.N; n++ {
		if err := jmc(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// Jx

func TestJx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	if a.NoError(v.Eval(Opcode(0x12003713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	if a.NoError(v.Eval(Opcode(0x12013713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		v.Eval(Opcode(0x120F3713)),
		"Unknown condition index should yield an error",
	)
}

func BenchmarkJx(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		// Avoid producing unknown conditions
		if err := jx(v, Opcode(n&0xFF0EFFFF)); err != nil {
			b.Fatal(err)
		}
	}
}

// JME

func TestJme(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[4] = 42

	if a.NoError(v.Eval(Opcode(0x13453713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	v.Regs[5] = 42

	if a.NoError(v.Eval(Opcode(0x13453713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkJme(b *testing.B) {
	v := NewVirtualMachine()
	// Half the registers are set to 0, others are set to 1.
	// This helps ensuring that both cases (Rx == Ry and Rx != Ry) happen
	// roughly the same number of times.
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i % 2)
	}
	for n := 0; n < b.N; n++ {
		if err := jme(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CALL HHLL

func TestCallHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.PC = Pointer(0xBEEF)

	if a.NoError(v.Eval(Opcode(0x14003713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(Pointer(StackStart+2), v.SP, "SP didn't move up")
		p, _ := v.PointerAt(Pointer(StackStart))
		a.Equalf(Pointer(0xBEEF), p, "Didn't push the right address")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkCallHHLL(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := callHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2 // Avoid stack overflow
	}
}

// RET

func TestRet(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.PC = Pointer(0xBEEF)
	a.NoError(v.Eval(Opcode(0x14003713))) // CALL 1337
	a.NoError(v.Check())
	a.Equalf(Pointer(0x1337), v.PC, "CALL didn't move PC")
	a.Equalf(Pointer(StackStart+2), v.SP, "CALL didn't stack return address")

	if a.NoError(v.Eval(Opcode(0x15000000))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0xBEEF), v.PC, "PC didn't return to 0x1337")
		a.Equalf(Pointer(StackStart), v.SP, "SP didn't move down")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkRet(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		v.SP += 2 // Avoid stack underflow
		if err := ret(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// JMP Rx

func TestJmpRx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[5] = 0x1337

	if a.NoError(v.Eval(Opcode(0x16050000))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "PC didn't move to 0x1337")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkJmpRx(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := jmpRx(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CX HHLL

func TestCx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	if a.NoError(v.Eval(Opcode(0x17003713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(Pointer(StackStart), v.SP, "SP shouldn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	if a.NoError(v.Eval(Opcode(0x17013713))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "PC didn't move to 0x1337")
		a.Equalf(Pointer(StackStart+2), v.SP, "SP didnt move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		v.Eval(Opcode(0x170F3713)),
		"Unknown condition should yield an error",
	)
}

func BenchmarkCx(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		// Avoid producing unknown conditions
		if err := cx(v, Opcode(n&0xFF0EFFFF)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}

// CALL Rx

func TestCallRx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Regs[5] = 0x1337

	if a.NoError(v.Eval(Opcode(0x18050000))) {
		a.NoError(v.Check())
		a.Equalf(Pointer(0x1337), v.PC, "PC didn't move to 0x1337")
		a.Equalf(Pointer(StackStart+2), v.SP, "SP didn't move")
		a.Equalf(CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkCallRx(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		if err := callRx(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}
