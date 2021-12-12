package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

var dayNineExampleInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestDayNine(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		require := require.New(t)
		in := dayNineExampleInput
		out, err := DayNine(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		require.Equal(15, out.PartOneAnswer)
		require.Equal(1134, out.PartTwoAnswer)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2021, 9)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		out, err := DayNine(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
}
