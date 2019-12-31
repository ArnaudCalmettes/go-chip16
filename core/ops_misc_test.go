package core

import "testing"

func BenchmarkNop(b *testing.B) {
	v := NewVirtualMachine()
	for n := 0; n < b.N; n++ {
		nop(v, Opcode(n))
	}
}
