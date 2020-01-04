package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

func BenchmarkNop(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		nop(v, vm.Opcode(n))
	}
}
