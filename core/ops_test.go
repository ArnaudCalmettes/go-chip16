package core

import "github.com/stretchr/testify/assert"

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
	flags CPUFlags,
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
	flags CPUFlags,
	op string,
) {
	a.Equalf(
		test.exp, res, "%d %s %d != %d",
		test.x, op, test.y, test.exp,
	)
	checkOpFlags(a, test, flags, op)
}
