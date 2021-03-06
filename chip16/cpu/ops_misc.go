package cpu

import (
	"math/rand"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

// Do nothing
func nop(*vm.State, vm.Opcode) error { return nil }

// Clear the screen (reset BG and FG)
func cls(v *vm.State, _ vm.Opcode) error {
	v.Graphics.Clear()
	return nil
}

// Draw sprite from [HHLL] at (Rx, Ry)
func drwRxRyHHLL(v *vm.State, o vm.Opcode) error {
	c, err := v.Graphics.DrawSprite(
		int(v.Regs[o.X()]),
		int(v.Regs[o.Y()]),
		v.RAM[o.HHLL():],
	)
	v.Flags.SetCarry(c)
	return err
}

// Draw sprite from [Rz] at (Rx, Ry)
func drwRxRyRz(v *vm.State, o vm.Opcode) error {
	c, err := v.Graphics.DrawSprite(
		int(v.Regs[o.X()]),
		int(v.Regs[o.Y()]),
		v.RAM[uint16(o.Z()):],
	)
	v.Flags.SetCarry(c)
	return err
}

// Set the background color index to N
func bgcN(v *vm.State, o vm.Opcode) error {
	v.Graphics.BG = o.N()
	return nil
}

// Set sprite width to LL and height to HH
func sprHHLL(v *vm.State, o vm.Opcode) error {
	v.Graphics.SpriteW = o.LL()
	v.Graphics.SpriteH = o.HH()
	return nil
}

// Store random number in RX (max. HHLL)
func rndRxHHLL(v *vm.State, o vm.Opcode) error {
	max := int(o.HHLL()) + 1
	v.Regs[o.X()] = int16(rand.Intn(max))
	return nil
}

func flip(v *vm.State, o vm.Opcode) error {
	hh := o.HH()
	v.Graphics.HFlip = (hh&0x02 != 0)
	v.Graphics.VFlip = (hh&0x01 != 0)
	return nil
}

func init() {
	setOp(0x00, "NOP", nop)
	setOp(0x01, "CLS", cls)
	// VBLNK
	setOp(0x03, "BGC N", bgcN)
	setOp(0x04, "SPR HHLL", sprHHLL)
	setOp(0x05, "DRW RX, RY, HHLL", drwRxRyHHLL)
	setOp(0x06, "DRW RX, RY, RZ", drwRxRyRz)
	setOp(0x07, "RND Rx, HHLL", rndRxHHLL)
	setOp(0x08, "FLIP HH", flip)
	// SND0
	// SND1 HHLL
	// SND2 HHLL
	// SND3 HHLL
	// SNP Rx, HHLL
	// SNG AD VT SR
}
