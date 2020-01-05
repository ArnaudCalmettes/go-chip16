package cpu

import (
	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

// Load palette from [HHLL]
func palHHLL(v *vm.State, o vm.Opcode) error {
	addr := o.HHLL()
	return v.Graphics.LoadPalette(v.RAM[addr:])
}

// Load palette from [Rx]
func palRx(v *vm.State, o vm.Opcode) error {
	addr := uint16(v.Regs[o.X()])
	return v.Graphics.LoadPalette(v.RAM[addr:])
}

func init() {
	setOp(0xD0, "PAL HHLL", palHHLL)
	setOp(0xD1, "PAL Rx", palRx)
}
