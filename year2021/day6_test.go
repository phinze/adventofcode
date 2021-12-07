package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

var daySixExampleInput = `3,4,3,1,2`

func TestDaySix(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		require := require.New(t)
		in := daySixExampleInput
		out, err := DaySix(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		require.Equal(5934, out.PartOneAnswer)
		require.Equal(26984457539, out.PartTwoAnswer)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2021, 6)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		out, err := DaySix(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
}
