package vm

import (
	"encoding/binary"
	"fmt"
)

const (
	// MemSize is the total memory size
	MemSize = 65536

	// RAMStart is the start address of the RAM
	RAMStart = 0x0000

	// StackStart is the start address of the stack
	StackStart = 0xFDF0

	// IOStart is the start address of the IO registers
	IOStart = 0xFFF0

	// PointerMax is the maximum valid 16 bit address
	PointerMax = 0xFFFE
)

// Pointer is a 16-bit pointer type
type Pointer uint16

// State modelizes the chip16 VM's state
type State struct {
	// PC is the Program Counter
	PC Pointer

	// SP is the Stack Pointer
	SP Pointer

	// RAM is a byte slice depicting the console's RAM
	RAM []byte

	// Regs is an array of 16 CPU Registers
	Regs [16]int16

	// Flags is the set of CPU Flags
	Flags CPUFlags
}

// NewState creates a new State
func NewState() *State {
	return &State{
		PC:  RAMStart,
		SP:  StackStart,
		RAM: make([]byte, MemSize),
	}
}

// Int16At reads a signed int16 at given address in RAM
func (v *State) Int16At(addr Pointer) (int16, error) {
	if addr > PointerMax {
		return 0, fmt.Errorf("address out of bounds")
	}
	return int16(binary.LittleEndian.Uint16(v.RAM[addr:])), nil
}

// PutInt16At writes a signed int16 at given address in RAM
func (v *State) PutInt16At(val int16, addr Pointer) error {
	if addr > PointerMax {
		return fmt.Errorf("address out of bounds")
	}
	binary.LittleEndian.PutUint16(v.RAM[addr:], uint16(val))
	return nil
}

// PointerAt reads a pointer at given address in RAM
func (v *State) PointerAt(addr Pointer) (Pointer, error) {
	if addr > PointerMax {
		return 0, fmt.Errorf("address out of bounds")
	}
	return Pointer(binary.LittleEndian.Uint16(v.RAM[addr:])), nil
}

// PutPointerAt writes a pointer at given address in RAM
func (v *State) PutPointerAt(val Pointer, addr Pointer) error {
	if addr > PointerMax {
		return fmt.Errorf("address out of bounds")
	}
	binary.LittleEndian.PutUint16(v.RAM[addr:], uint16(val))
	return nil
}

// Check sanity of the current vm state
func (v *State) Check() error {
	if uint16(v.PC) >= StackStart {
		return fmt.Errorf("PC overflow: PC = %#04x", v.PC)
	}
	if uint16(v.SP) < StackStart {
		return fmt.Errorf("stack underflow: SP = %#04x", v.SP)
	}
	if uint16(v.SP) >= IOStart {
		return fmt.Errorf("stack overflow: SP = %#04x", v.SP)
	}
	return nil
}
