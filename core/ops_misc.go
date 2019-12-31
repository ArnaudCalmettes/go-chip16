package core

func nop(*VirtualMachine, Opcode) error { return nil }

func init() {
	setOp(0x00, "NOP", nop)
}
