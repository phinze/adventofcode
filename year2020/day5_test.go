package year2020

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

var dayFiveExampleInput = `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`

func TestDayFive(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		require := require.New(t)
		in := dayFiveExampleInput
		out, err := DayFive(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		require.Equal(820, out.PartOneAnswer)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2020, 5)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		out, err := DayFive(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
}
