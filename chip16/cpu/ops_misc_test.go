package cpu

import (
	"fmt"
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

// NOP

func BenchmarkNop(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		nop(v, vm.Opcode(n))
	}
}

// CLS

func BenchmarkCls(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		cls(v, vm.Opcode(n))
	}
}

// BGC N

func TestBgcN(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	if a.NoError(Eval(v, vm.Opcode(0x03000400))) {
		a.Equal(uint8(4), v.Graphics.BG, "Didn't set BG")
	}
}

func BenchmarkBgc(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		bgcN(v, vm.Opcode(n))
	}
}

// SPR HHLL

func TestSprHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	if a.NoError(Eval(v, vm.Opcode(0x04002010))) {
		a.Equal(uint8(0x10), v.Graphics.SpriteH, "Wrong sprite height")
		a.Equal(uint8(0x20), v.Graphics.SpriteW, "Wrong sprite width")
	}
}

func BenchmarkSprHHLL(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		sprHHLL(v, vm.Opcode(n))
	}
}

// DRW Rx, Ry, HHLL
func TestDrwRxRyHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	// That's much more than memory can hold
	v.Graphics.SpriteW = 255
	v.Graphics.SpriteH = 255
	a.Error(
		Eval(v, vm.Opcode(0x050000FF)),
		"out of bounds sprite didn't return an error",
	)

	v.Graphics.SpriteW = 16
	v.Graphics.SpriteH = 16
	a.NoError(Eval(v, vm.Opcode(0x05000000)))
}

func BenchmarkDrwRxRyHHLL(b *testing.B) {
	v := vm.NewState()

	// fill RAM with non-null pixels
	for i := 0; i < vm.StackStart; i++ {
		v.RAM[i] = 0xAA
	}

	benches := []uint8{4, 8, 16, 32, 64, 128}
	for _, size := range benches {
		b.Run(
			fmt.Sprintf("%dx%d", size, size),
			func(b *testing.B) {
				v.Graphics.SpriteW = size / 2
				v.Graphics.SpriteH = size
				for n := 0; n < b.N; n++ {
					op := vm.Opcode(0)
					if err := drwRxRyHHLL(v, op); err != nil {
						b.Fatal(err)
					}
				}
			},
		)
	}
}

// RND HHLL

func TestRndHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	a.NoError(Eval(v, vm.Opcode(0x07010000)))
	a.NoError(Eval(v, vm.Opcode(0x0701FFFF)))
}

func BenchmarkRndHHLL(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		rndRxHHLL(v, vm.Opcode(n))
	}
}

// FLIP

func TestFlip(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range []struct {
		hhll uint16
		h    bool
		v    bool
	}{
		{0x000, false, false},
		{0x100, false, true},
		{0x200, true, false},
		{0x300, true, true},
	} {
		op := vm.Opcode(0x08000000).WithHHLL(test.hhll)
		if a.NoError(Eval(v, op)) {
			a.Equalf(test.h, v.Graphics.HFlip, "(%d) wrong HFlip", test.hhll)
			a.Equalf(test.v, v.Graphics.VFlip, "(%d) wrong VFlip", test.hhll)
		}
	}
}

func BenchmarkFlip(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		flip(v, vm.Opcode(n))
	}
}
