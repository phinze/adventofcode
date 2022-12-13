package year2022

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDayTwelve_simple(t *testing.T) {
	in := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`
	out, err := DayTwelve(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	require.Equal(t, 31, out.PartOneAnswer)
	require.Equal(t, 29, out.PartTwoAnswer)
}

func TestDayTwelve(t *testing.T) {
	in, err := aoc.FetchInput(2022, 12)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayTwelve(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
