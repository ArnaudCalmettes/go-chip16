package core

type arithTestCase struct {
	x   int16
	y   int16
	exp int16 // expected result
	c   bool  // expected carry flag
	o   bool  // expected overflow flag
	n   bool  // expected negative flag
	z   bool  // expected zero flag
}
