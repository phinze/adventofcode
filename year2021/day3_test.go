package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestIsBitOn(t *testing.T) {
	require := require.New(t)

	require.True(isBitOn(0b01001, 0))
	require.False(isBitOn(0b01001, 1))
	require.False(isBitOn(0b01001, 2))
	require.True(isBitOn(0b01001, 3))
	require.False(isBitOn(0b01001, 4))
	require.False(isBitOn(0b01001, 5))
}

var dayThreeExampleInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func TestBitHelpers(t *testing.T) {
	require := require.New(t)
	in, err := parseDayThree(dayThreeExampleInput)
	if err != nil {
		t.Fatal(err)
	}
	nums := in.Diagnostics

	require.Equal(7, numBitsInPosition(nums, 4))
	require.Equal(5, numBitsInPosition(nums, 3))
	require.Equal(8, numBitsInPosition(nums, 2))
	require.Equal(7, numBitsInPosition(nums, 1))
	require.Equal(5, numBitsInPosition(nums, 0))

	require.True(majorityBitInPosition(nums, 4))
	require.False(majorityBitInPosition(nums, 3))
	require.True(majorityBitInPosition(nums, 2))
	require.True(majorityBitInPosition(nums, 1))
	require.False(majorityBitInPosition(nums, 0))

	oddNums := []uint64{
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b11100,
		0b10000,
		0b11001,
	}
	require.False(majorityBitInPosition(oddNums, 3))
}

func TestDayThree(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		in := dayThreeExampleInput
		out, err := DayThree(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
	t.Run("official input", func(t *testing.T) {
		in, err := aoc.FetchInput(2021, 3)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		out, err := DayThree(in)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		t.Logf("out: %#v", out)
	})
}
