package year2020

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDayThreeRow(t *testing.T) {
	t.Run("detects tree at index", func(t *testing.T) {
		require := require.New(t)
		row := &DayThreeRow{Source: "...#..#...#"}

		require.False(row.HasTreeAt(0))
		require.True(row.HasTreeAt(3))
		require.True(row.HasTreeAt(6))
	})

	t.Run("wraps around indefinitely", func(t *testing.T) {
		require := require.New(t)
		row := &DayThreeRow{Source: "..#"}

		require.False(row.HasTreeAt(0))
		require.False(row.HasTreeAt(1))
		require.True(row.HasTreeAt(2))
		require.False(row.HasTreeAt(3))
		require.False(row.HasTreeAt(4))
		require.True(row.HasTreeAt(5))
		require.False(row.HasTreeAt(6))
		require.False(row.HasTreeAt(7))
		require.True(row.HasTreeAt(8))
	})
}

func TestDayThree(t *testing.T) {
	in, err := aoc.FetchInput(2020, 3)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayThree(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
