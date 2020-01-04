package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type unaryTestCase struct {
	x   int16
	exp int16
	z   bool
	n   bool
}

var notTestCases = []unaryTestCase{
	{-21246, 21245, false, false},
	{21245, -21246, false, true},
	{-1, 0, true, false},
}

var negTestCases = []unaryTestCase{
	{1, -1, false, true},
	{-1, 1, false, false},
	{0, 0, true, false},
}

func checkUnary(
	a *assert.Assertions,
	t *unaryTestCase,
	res int16,
	f CPUFlags,
	op string,
) {
	a.Equalf(t.exp, res, "%s%d != %d", op, t.x, t.exp)
	a.Equalf(t.z, f.Zero(), "(%s%d) wrong zero flag", op, t.x)
	a.Equalf(t.n, f.Negative(), "(%s%d) wrong negative flag", op, t.x)
	a.Falsef(f.Carry(), "(%s%d) wrong carry", op, t.x)
	a.Falsef(f.Overflow(), "(%s%d) wrong overflow flag", op, t.x)
}

// NOTI Rx, HHLL

func TestNotiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range notTestCases {
		op := Opcode(0xE0010000).WithHHLL(uint16(test.x))
		if a.NoError(v.Eval(op)) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "^")
		}
	}
}

func BenchmarkNotiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := notiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NOT Rx

func TestNotRx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range notTestCases {
		v.Regs[1] = test.x
		if a.NoError(v.Eval(Opcode(0xE1010000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "^")
		}
	}
}

func BenchmarkNotRx(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := notRx(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NOT Rx, Ry

func TestNotRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range notTestCases {
		v.Regs[3] = test.x
		if a.NoError(v.Eval(Opcode(0xE2310000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "^")
		}
	}
}

func BenchmarkNotRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := notRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NEGI Rx, HHLL

func TestNegiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range negTestCases {
		op := Opcode(0xE3010000).WithHHLL(uint16(test.x))
		if a.NoError(v.Eval(op)) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "-")
		}
	}
}

func BenchmarkNegiRxHHLL(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := negiRxHHLL(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NEG Rx

func TestNegRx(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range negTestCases {
		v.Regs[1] = test.x
		if a.NoError(v.Eval(Opcode(0xE4010000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "-")
		}
	}
}

func BenchmarkNegRx(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := negRx(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NEG Rx, Ry

func TestNegRxRy(t *testing.T) {
	a := assert.New(t)
	v := NewVirtualMachine()

	for _, test := range negTestCases {
		v.Regs[3] = test.x
		if a.NoError(v.Eval(Opcode(0xE5310000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "-")
		}
	}
}

func BenchmarkNegRxRy(b *testing.B) {
	v := NewVirtualMachine()

	for n := 0; n < b.N; n++ {
		if err := negRxRy(v, Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
