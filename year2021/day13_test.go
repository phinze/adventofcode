package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

var dayThirteenExampleInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestDayThirteen(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		require := require.New(t)
		in := dayThirteenExampleInput
		out, err := DayThirteen(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		require.Equal(17, out.PartOneAnswer)
		require.Equal(0, out.PartTwoAnswer)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2021, 13)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		out, err := DayThirteen(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
}
