package core

import "fmt"

const (
	flagC = 1 << 1 // Carry
	flagZ = 1 << 2 // Zero
	flagO = 1 << 6 // Overflow
	flagN = 1 << 7 // Negative
)

// CPUFlags implements flags set by the chip16's CPU
type CPUFlags uint8

func (f *CPUFlags) Clear() {
	*f = 0x00
}

// SetCarry sets the carry flag if p is true, else clears it
func (f *CPUFlags) SetCarry(p bool) {
	if p {
		*f |= flagC
	} else {
		*f &^= flagC
	}
}

// SetZero sets the zero flag if p is true, else clears it
func (f *CPUFlags) SetZero(p bool) {
	if p {
		*f |= flagZ
	} else {
		*f &^= flagZ
	}
}

// SetOverflow sets the overflow flag if p is true, else clears it
func (f *CPUFlags) SetOverflow(p bool) {
	if p {
		*f |= flagO
	} else {
		*f &^= flagO
	}
}

// SetNegative sets the negative flag if p is true, else clears it
func (f *CPUFlags) SetNegative(p bool) {
	if p {
		*f |= flagN
	} else {
		*f &^= flagN
	}
}

// Carry returns true if the Carry flag is raised
func (f CPUFlags) Carry() bool {
	return f&flagC != 0
}

// Overflow returns true if the Overflow flag is raised
func (f CPUFlags) Overflow() bool {
	return f&flagO != 0
}

// Negative returns true if the Negative flag is raised
func (f CPUFlags) Negative() bool {
	return f&flagN != 0
}

// Zero returns true if the Zero flag is raised
func (f CPUFlags) Zero() bool {
	return f&flagZ != 0
}

// Condition evaluates a conditional over the flags given its index.
// Valid indexes are in the 0x0 - 0xE range.
func (f CPUFlags) Condition(index uint8) (bool, error) {
	switch index {
	case 0x0:
		// Equal (Zero)
		return f&flagZ != 0, nil

	case 0x1:
		// Not Equal (Non-Zero)
		return f&flagZ == 0, nil

	case 0x2:
		// Negative
		return f&flagN != 0, nil

	case 0x3:
		// Non-Negative
		return f&flagN == 0, nil

	case 0x4:
		// Positive
		return f&(flagN|flagZ) == 0, nil

	case 0x5:
		// Overflow
		return f&flagO != 0, nil

	case 0x6:
		// No Overflow
		return f&flagO == 0, nil

	case 0x7:
		// Above (Unsigned Greater Than)
		return f&(flagC|flagZ) == 0, nil

	case 0x8:
		// Above or Equal (Unsigned Greater Than or Equal)
		return f&flagC == 0, nil

	case 0x9:
		// Below (Unsigned Less Than)
		return f&flagC != 0, nil

	case 0xA:
		// Below Equal (Unsigned Less Than or Equal)
		return f&(flagC|flagZ) != 0, nil

	case 0xB:
		// Signed Greater Than
		return f&flagO == f&flagN && f&flagZ == 0, nil

	case 0xC:
		// Signed Greater Than or Equal
		return f&flagO == f&flagN, nil

	case 0xD:
		// Signed Less Than
		return f&flagO != f&flagN, nil

	case 0xE:
		// Signed Less Than or Equal
		return f&flagO != f&flagN || f&flagZ != 0, nil

	default:
		return false, fmt.Errorf("unknown condition %#02x", index)
	}
}
