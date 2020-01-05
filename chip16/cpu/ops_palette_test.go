package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/graphics"
	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

func setupTestPalette(v *vm.State) {
	for i := 0; i < graphics.PaletteSize; i++ {
		v.RAM[0x1337+i] = byte(i)
	}
}

// PAL HHLL

func TestPalHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	setupTestPalette(v)
	op := vm.Opcode(0xD0003713)
	if a.NoError(Eval(v, op)) {
		for i, c := range v.Graphics.Palette {
			a.Equal(uint8(3*i), c.B)
			a.Equal(uint8(3*i+1), c.G)
			a.Equal(uint8(3*i+2), c.R)
		}
	}

	a.Error(
		Eval(v, vm.Opcode(0xD000F0FF)),
		"out of bounds memory access didn't return an error",
	)
}

func BenchmarkPalHHLL(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		palHHLL(v, vm.Opcode(n&0xFFFFF0FF))
	}
}

// PAL Rx

func TestPalRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	setupTestPalette(v)
	v.Regs[4] = 0x1337
	if a.NoError(Eval(v, vm.Opcode(0xD1040000))) {
		for i, c := range v.Graphics.Palette {
			a.Equal(uint8(3*i), c.B)
			a.Equal(uint8(3*i+1), c.G)
			a.Equal(uint8(3*i+2), c.R)
		}
	}

	v.Regs[4] = -1 // 0xFFFE
	a.Error(
		Eval(v, vm.Opcode(0xD1040000)),
		"out of bounds memory access didn't return an error",
	)
}

func BenchmarkPalRx(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		palRx(v, vm.Opcode(n))
	}
}
