package core

// Utility call method
func (v *VirtualMachine) call(p Pointer) {
	v.PutPointerAt(v.PC, v.SP)
	v.SP += 2
	v.PC = p
}

// Unconditional jump to HHLL
func jmpHHLL(v *VirtualMachine, o Opcode) error {
	v.PC = Pointer(o.HHLL())
	return nil
}

// Jump to HHLL if carry is set
func jmc(v *VirtualMachine, o Opcode) error {
	if v.Flags.Carry() {
		v.PC = Pointer(o.HHLL())
	}
	return nil
}

// Jump to HHLL if a flag condition X is met
func jx(v *VirtualMachine, o Opcode) error {
	if cond, err := v.Flags.Condition(o.X()); err != nil {
		return err
	} else if cond {
		v.PC = Pointer(o.HHLL())
	}
	return nil
}

// Jump to HHLL if Rx == RY
func jme(v *VirtualMachine, o Opcode) error {
	if v.Regs[o.X()] == v.Regs[o.Y()] {
		v.PC = Pointer(o.HHLL())
	}
	return nil
}

// Perform an inconditional call to HHLL
func callHHLL(v *VirtualMachine, o Opcode) error {
	v.call(Pointer(o.HHLL()))
	return nil
}

// Return from function call
func ret(v *VirtualMachine, _ Opcode) error {
	var err error
	v.SP -= 2
	v.PC, err = v.PointerAt(v.SP)
	return err
}

// Unconditional jump to Rx
func jmpRx(v *VirtualMachine, o Opcode) error {
	v.PC = Pointer(v.Regs[o.X()])
	return nil
}

// Conditional call to HHLL
func cx(v *VirtualMachine, o Opcode) error {
	if cond, err := v.Flags.Condition(o.X()); err != nil {
		return err
	} else if cond {
		v.call(Pointer(o.HHLL()))
	}
	return nil
}

// Perform an inconditional call to Rx
func callRx(v *VirtualMachine, o Opcode) error {
	v.call(Pointer(v.Regs[o.X()]))
	return nil
}

func init() {
	setOp(0x10, "JMP HHLL", jmpHHLL)
	setOp(0x11, "JMC HHLL", jmc)
	setOp(0x12, "Jx HHLL", jx)
	setOp(0x13, "JME Rx, Ry, HHLL", jme)
	setOp(0x14, "CALL HHLL", callHHLL)
	setOp(0x15, "RET", ret)
	setOp(0x16, "JMP Rx", jmpRx)
	setOp(0x17, "Cx HHLL", cx)
	setOp(0x18, "CALL Rx", callRx)
}
