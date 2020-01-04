package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

// JMP HHLL

func TestJmpHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	if a.NoError(Eval(v, vm.Opcode(0x10003713))) {
		a.Equalf(vm.Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkJmpHHLL(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		if err := jmpHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// JMC

func TestJmc(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	if a.NoError(Eval(v, vm.Opcode(0x11003713))) {
		a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	v.Flags.SetCarry(true)

	if a.NoError(Eval(v, vm.Opcode(0x11003713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Truef(v.Flags.Carry(), "Carry flag should be set")
	}
}

func BenchmarkJmcNoCarry(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		if err := jmc(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJmcCarry(b *testing.B) {
	v := vm.NewState()
	v.Flags.SetCarry(true)
	for n := 0; n < b.N; n++ {
		if err := jmc(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// Jx

func TestJx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	if a.NoError(Eval(v, vm.Opcode(0x12003713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	if a.NoError(Eval(v, vm.Opcode(0x12013713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		Eval(v, vm.Opcode(0x120F3713)),
		"Unknown condition index should yield an error",
	)
}

func BenchmarkJx(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		// Avoid producing unknown conditions
		if err := jx(v, vm.Opcode(n&0xFF0EFFFF)); err != nil {
			b.Fatal(err)
		}
	}
}

// JME

func TestJme(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.Regs[4] = 42

	if a.NoError(Eval(v, vm.Opcode(0x13453713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	v.Regs[5] = 42

	if a.NoError(Eval(v, vm.Opcode(0x13453713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkJme(b *testing.B) {
	v := vm.NewState()
	// Half the registers are set to 0, others are set to 1.
	// This helps ensuring that both cases (Rx == Ry and Rx != Ry) happen
	// roughly the same number of times.
	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i % 2)
	}
	for n := 0; n < b.N; n++ {
		if err := jme(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CALL HHLL

func TestCallHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.PC = vm.Pointer(0xBEEF)

	if a.NoError(Eval(v, vm.Opcode(0x14003713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "Didn't jump to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart+2), v.SP, "SP didn't move up")
		p, _ := v.PointerAt(vm.Pointer(vm.StackStart))
		a.Equalf(vm.Pointer(0xBEEF), p, "Didn't push the right address")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkCallHHLL(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		if err := callHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2 // Avoid stack overflow
	}
}

// RET

func TestRet(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.PC = vm.Pointer(0xBEEF)
	a.NoError(Eval(v, vm.Opcode(0x14003713))) // CALL 1337
	a.NoError(v.Check())
	a.Equalf(vm.Pointer(0x1337), v.PC, "CALL didn't move PC")
	a.Equalf(vm.Pointer(vm.StackStart+2), v.SP, "CALL didn't stack return address")

	if a.NoError(Eval(v, vm.Opcode(0x15000000))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0xBEEF), v.PC, "PC didn't return to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP didn't move down")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkRet(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		v.SP += 2 // Avoid stack underflow
		if err := ret(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// JMP Rx

func TestJmpRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.Regs[5] = 0x1337

	if a.NoError(Eval(v, vm.Opcode(0x16050000))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "PC didn't move to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkJmpRx(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		if err := jmpRx(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// CX HHLL

func TestCx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	if a.NoError(Eval(v, vm.Opcode(0x17003713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move: condition isn't met")
		a.Equalf(vm.Pointer(vm.StackStart), v.SP, "SP shouldn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	if a.NoError(Eval(v, vm.Opcode(0x17013713))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "PC didn't move to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart+2), v.SP, "SP didnt move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}

	a.Errorf(
		Eval(v, vm.Opcode(0x170F3713)),
		"Unknown condition should yield an error",
	)
}

func BenchmarkCx(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		// Avoid producing unknown conditions
		if err := cx(v, vm.Opcode(n&0xFF0EFFFF)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}

// CALL Rx

func TestCallRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.Regs[5] = 0x1337

	if a.NoError(Eval(v, vm.Opcode(0x18050000))) {
		a.NoError(v.Check())
		a.Equalf(vm.Pointer(0x1337), v.PC, "PC didn't move to 0x1337")
		a.Equalf(vm.Pointer(vm.StackStart+2), v.SP, "SP didn't move")
		a.Equalf(vm.CPUFlags(0x0), v.Flags, "Flags shouldn't move")
	}
}

func BenchmarkCallRx(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		if err := callRx(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}
