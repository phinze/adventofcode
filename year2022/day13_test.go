package year2022

import (
	"log"
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDayThirteen_simple(t *testing.T) {
	in := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`
	out, err := DayThirteen(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	expected := 13
	require.Equal(t, expected, out.PartOneAnswer)
}

func TestDayThirteen(t *testing.T) {
	in, err := aoc.FetchInput(2022, 13)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayThirteen(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// require.Equal(t, dayThirteenComparisonOracle, out.PartOneComparisons)
	for i, result := range out.PartOneComparisons {
		if dayThirteenComparisonOracle[i] != result {
			log.Printf("pair %d mismatch", i+1)
		}
	}
	// t.Logf("out: %s", spew.Sdump(out))
}

var dayThirteenComparisonOracle = []bool{
	false,
	false,
	true,
	true,
	true,
	true,
	false,
	false,
	true,
	false,
	true,
	false,
	true,
	false,
	false,
	true,
	true,
	true,
	true,
	true,
	false,
	false,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	true,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	true,
	true,
	false,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	true,
	true,
	true,
	true,
	false,
	true,
	false,
	true,
	true,
	false,
	false,
	true,
	false,
	true,
	true,
	true,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
	true,
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
	false,
	true,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	true,
	false,
	true,
	false,
	false,
	true,
	true,
}
