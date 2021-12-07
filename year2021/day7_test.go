package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

var daySevenExampleInput = `16,1,2,0,4,2,7,1,2,14`

func TestDaySeven(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		require := require.New(t)
		in := daySevenExampleInput
		out, err := DaySeven(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		require.Equal(37, out.PartOneAnswer)
		require.Equal(168, out.PartTwoAnswer)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2021, 7)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		out, err := DaySeven(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
}
