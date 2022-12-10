package year2022

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDayNine_simple(t *testing.T) {
	in := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
	out, err := DayNine(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	require.Equal(t, 13, out.VisitedPositions)
}

func TestDayNine_part2(t *testing.T) {
	in := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	out, err := DayNine(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	require.Equal(t, 13, out.VisitedPositions)
}

func TestDayNine(t *testing.T) {
	in, err := aoc.FetchInput(2022, 9)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayNine(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
