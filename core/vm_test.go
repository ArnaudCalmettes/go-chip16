package core

import "testing"

func BenchmarkEvalNop(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		v.Eval(Opcode(0x000000))
	}
}
