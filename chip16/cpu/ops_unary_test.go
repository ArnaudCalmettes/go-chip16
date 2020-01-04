package cpu

import (
	"testing"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
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
	f vm.CPUFlags,
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
	v := vm.NewState()

	for _, test := range notTestCases {
		op := vm.Opcode(0xE0010000).WithHHLL(uint16(test.x))
		if a.NoError(Eval(v, op)) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "^")
		}
	}
}

func BenchmarkNotiRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := notiRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NOT Rx

func TestNotRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range notTestCases {
		v.Regs[1] = test.x
		if a.NoError(Eval(v, vm.Opcode(0xE1010000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "^")
		}
	}
}

func BenchmarkNotRx(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := notRx(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NOT Rx, Ry

func TestNotRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range notTestCases {
		v.Regs[3] = test.x
		if a.NoError(Eval(v, vm.Opcode(0xE2310000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "^")
		}
	}
}

func BenchmarkNotRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := notRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NEGI Rx, HHLL

func TestNegiRxHHLL(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range negTestCases {
		op := vm.Opcode(0xE3010000).WithHHLL(uint16(test.x))
		if a.NoError(Eval(v, op)) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "-")
		}
	}
}

func BenchmarkNegiRxHHLL(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := negiRxHHLL(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NEG Rx

func TestNegRx(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range negTestCases {
		v.Regs[1] = test.x
		if a.NoError(Eval(v, vm.Opcode(0xE4010000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "-")
		}
	}
}

func BenchmarkNegRx(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := negRx(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}

// NEG Rx, Ry

func TestNegRxRy(t *testing.T) {
	a := assert.New(t)
	v := vm.NewState()

	for _, test := range negTestCases {
		v.Regs[3] = test.x
		if a.NoError(Eval(v, vm.Opcode(0xE5310000))) {
			checkUnary(a, &test, v.Regs[1], v.Flags, "-")
		}
	}
}

func BenchmarkNegRxRy(b *testing.B) {
	v := vm.NewState()

	for n := 0; n < b.N; n++ {
		if err := negRxRy(v, vm.Opcode(n)); err != nil {
			b.Fatal(err)
		}
	}
}
