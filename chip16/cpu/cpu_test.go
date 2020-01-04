package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
	"github.com/stretchr/testify/assert"
)

type arithTestCase struct {
	x   int16
	y   int16
	exp int16 // expected result
	c   bool  // expected carry flag
	o   bool  // expected overflow flag
	n   bool  // expected negative flag
	z   bool  // expected zero flag
}

func checkOpFlags(
	a *assert.Assertions,
	test *arithTestCase,
	flags vm.CPUFlags,
	op string,
) {
	a.Equalf(
		test.c, flags.Carry(),
		"(%d %s %d) wrong carry flag", test.x, op, test.y,
	)
	a.Equalf(
		test.o, flags.Overflow(),
		"(%d %s %d) wrong overflow flag", test.x, op, test.y,
	)
	a.Equalf(
		test.n, flags.Negative(),
		"(%d %s %d) wrong negative flag", test.x, op, test.y,
	)
	a.Equalf(
		test.z, flags.Zero(),
		"(%d %s %d) wrong zero flag", test.x, op, test.y,
	)
}

func checkOpResults(
	a *assert.Assertions,
	test *arithTestCase,
	res int16,
	flags vm.CPUFlags,
	op string,
) {
	a.Equalf(
		test.exp, res, "%d %s %d != %d",
		test.x, op, test.y, test.exp,
	)
	checkOpFlags(a, test, flags, op)
}

func BenchmarkEvalNop(b *testing.B) {
	v := vm.NewState()
	for n := 0; n < b.N; n++ {
		Eval(v, vm.Opcode(0x000000))
	}
}
