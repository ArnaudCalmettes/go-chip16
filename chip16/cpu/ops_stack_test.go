package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

// PUSH Rx

func TestPushRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()
	v.Regs[1] = 0x1337

	if a.NoError(Eval(v, vm.Opcode(0xC0010000))) {
		a.Equal(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equal(vm.Pointer(vm.StackStart+2), v.SP, "SP didn't move")
		x, err := v.Int16At(v.SP - 2)
		a.NoError(err)
		a.Equal(int16(0x1337), x)
	}

	v.SP = vm.IOStart
	a.Errorf(
		Eval(v, vm.Opcode(0xC0010000)),
		"Stack overflow didn't return an error",
	)
}

func BenchmarkPushRx(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := pushRx(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}

// POP Rx

func TestPopRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.PutInt16At(0x1337, v.SP)
	v.SP += 2
	if a.NoError(Eval(v, vm.Opcode(0xC1010000))) {
		a.Equal(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equal(vm.Pointer(vm.StackStart), v.SP, "SP didn't move")
		a.Equal(int16(0x1337), v.Regs[1])
	}
	a.Errorf(
		Eval(v, vm.Opcode(0xC1010000)),
		"Stack underflow didn't return an error",
	)
}

func BenchmarkPopRx(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		v.SP += 2
		if err := popRx(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// PUSHALL

func TestPushAll(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for i := 0; i < len(v.Regs); i++ {
		v.Regs[i] = int16(i)
	}
	if a.NoError(Eval(v, vm.Opcode(0xC2000000))) {
		a.Equal(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equal(vm.Pointer(vm.StackStart+32), v.SP, "SP didn't move")
		for i := 0; i < len(v.Regs); i++ {
			x, _ := v.Int16At(vm.Pointer(vm.StackStart + 2*i))
			a.Equal(v.Regs[i], x)
		}
	}

	v.SP = vm.IOStart - 16
	a.Errorf(
		Eval(v, vm.Opcode(0xC2000000)),
		"Stack overflow didn't return an error",
	)
}

func BenchmarkPushAll(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := pushAll(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 32
	}
}

// POPALL

func TestPopAll(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for i := 0; i < 16; i++ {
		v.RAM[vm.StackStart+2*i] = byte(i)
	}
	v.SP += 32
	if a.NoError(Eval(v, vm.Opcode(0xC3000000))) {
		a.Equal(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equal(vm.Pointer(vm.StackStart), v.SP, "SP didn't move")
		for i := 0; i < len(v.Regs); i++ {
			a.Equal(int16(i), v.Regs[i])
		}
	}
	a.Errorf(
		Eval(v, vm.Opcode(0xC3000000)),
		"Stack underflow didn't return an error",
	)
}

func BenchmarkPopAll(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		v.SP += 32
		if err := popAll(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// PUSHF

func TestPushF(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.Flags = vm.CPUFlags(0xAA)
	if a.NoError(Eval(v, vm.Opcode(0xC4000000))) {
		a.Equal(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equal(vm.Pointer(vm.StackStart+2), v.SP, "SP didn't move")
		a.Equal(vm.CPUFlags(0xAA), vm.CPUFlags(v.RAM[vm.StackStart]))
	}

	v.SP = vm.IOStart
	a.Errorf(
		Eval(v, vm.Opcode(0xC4000000)),
		"Stack overflow didn't return an error",
	)
}

func BenchmarkPushF(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := pushF(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
		v.SP -= 2
	}
}

// PopF

func TestPopF(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	v.RAM[v.SP] = 0xAA
	v.SP += 2

	if a.NoError(Eval(v, vm.Opcode(0xC5000000))) {
		a.Equal(vm.Pointer(vm.RAMStart), v.PC, "PC shouldn't move")
		a.Equal(vm.Pointer(vm.StackStart), v.SP, "SP didn't move")
		a.Equal(vm.CPUFlags(0xAA), v.Flags)
	}

	a.Errorf(
		Eval(v, vm.Opcode(0xC5000000)),
		"Stack underflow didn't return an error",
	)
}

func BenchmarkPopF(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		v.SP += 2
		if err := popF(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
