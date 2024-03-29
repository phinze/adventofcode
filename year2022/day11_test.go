package year2022

import (
	"math/big"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDayEleven_simple(t *testing.T) {
	in := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`
	out, err := DayEleven(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	require.Equal(t, big.NewInt(10605), out.MonkeyBusiness)
	worryFreeBiz, _ := new(big.Int).SetString("2713310158", 10)
	require.Equal(t, worryFreeBiz, out.WorryFreeMonkeyBiz)
}

func TestDayEleven(t *testing.T) {
	in, err := aoc.FetchInput(2022, 11)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayEleven(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
