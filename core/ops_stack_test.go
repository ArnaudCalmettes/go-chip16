package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// PUSH Rx

func TestPushRx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()
	v.Regs[1] = 0x1337

	if a.NoError(v.Eval(Opcode(0xC0010000))) {
		a.Equal(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equal(Pointer(StackStart+2), v.SP, "SP didn't move")
		x, err := v.Int16At(v.SP - 2)
		a.NoError(err)
		a.Equal(int16(0x1337), x)
	}

	v.SP = IOStart
	a.Errorf(
		v.Eval(Opcode(0xC0010000)),
		"Stack overflow didn't return an error",
	)
}

func BenchmarkPushRx(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := pushRx(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}

// POP Rx

func TestPopRx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.PutInt16At(0x1337, v.SP)
	v.SP += 2
	if a.NoError(v.Eval(Opcode(0xC1010000))) {
		a.Equal(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equal(Pointer(StackStart), v.SP, "SP didn't move")
		a.Equal(int16(0x1337), v.Regs[1])
	}
	a.Errorf(
		v.Eval(Opcode(0xC1010000)),
		"Stack underflow didn't return an error",
	)
}

func BenchmarkPopRx(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		v.SP += 2
		if err := popRx(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// PUSHALL

func TestPushAll(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i)
	}
	if a.NoError(v.Eval(Opcode(0xC2000000))) {
		a.Equal(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equal(Pointer(StackStart+32), v.SP, "SP didn't move")
		for i := 0; i < len(v.Regs); i++ {
			x, _ := v.Int16At(Pointer(StackStart + 2*i))
			a.Equal(v.Regs[i], x)
		}
	}

	v.SP = IOStart - 16
	a.Errorf(
		v.Eval(Opcode(0xC2000000)),
		"Stack overflow didn't return an error",
	)
}

func BenchmarkPushAll(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := pushAll(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 32
	}
}

// POPALL

func TestPopAll(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for i := 0; i < 16; i++ {
		v.RAM[StackStart+2*i] = byte(i)
	}
	v.SP += 32
	if a.NoError(v.Eval(Opcode(0xC3000000))) {
		a.Equal(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equal(Pointer(StackStart), v.SP, "SP didn't move")
		for i := 0; i < len(v.Regs); i++ {
			a.Equal(int16(i), v.Regs[i])
		}
	}
	a.Errorf(
		v.Eval(Opcode(0xC3000000)),
		"Stack underflow didn't return an error",
	)
}

func BenchmarkPopAll(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		v.SP += 32
		if err := popAll(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// PUSHF

func TestPushF(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.Flags = CPUFlags(0xAA)
	if a.NoError(v.Eval(Opcode(0xC4000000))) {
		a.Equal(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equal(Pointer(StackStart+2), v.SP, "SP didn't move")
		a.Equal(CPUFlags(0xAA), CPUFlags(v.RAM[StackStart]))
	}

	v.SP = IOStart
	a.Errorf(
		v.Eval(Opcode(0xC4000000)),
		"Stack overflow didn't return an error",
	)
}

func BenchmarkPushF(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := pushF(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}

// PopF

func TestPopF(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	v.RAM[v.SP] = 0xAA
	v.SP += 2

	if a.NoError(v.Eval(Opcode(0xC5000000))) {
		a.Equal(Pointer(RAMStart), v.PC, "PC shouldn't move")
		a.Equal(Pointer(StackStart), v.SP, "SP didn't move")
		a.Equal(CPUFlags(0xAA), v.Flags)
	}

	a.Errorf(
		v.Eval(Opcode(0xC5000000)),
		"Stack underflow didn't return an error",
	)
}

func BenchmarkPopF(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		v.SP += 2
		if err := popF(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
