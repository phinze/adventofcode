package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

var dayFiveExampleInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func TestDayFive(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		require := require.New(t)
		in := dayFiveExampleInput
		out, err := DayFive(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		require.Equal(5, out.PartOneAnswer)
		require.Equal(12, out.PartTwoAnswer)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2021, 5)
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
