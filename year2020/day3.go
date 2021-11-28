package year2020

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type DayThreeRow struct {
	Source string
}

func (r *DayThreeRow) HasTreeAt(index int) bool {
	return r.Source[index%len(r.Source)] == '#'
}

type DayThreeInput struct {
	Rows []*DayThreeRow
}

type Sled struct {
	Right       int
	Down        int
	Pos         int
	NumTreesHit int
}

type DayThreeOutput struct {
	NumTreesHit     int
	MultiSledResult int
}

func parseDayThreeLine(line string) (*DayThreeRow, error) {
	return &DayThreeRow{Source: line}, nil
}

func parseDayThree(raw string) (*DayThreeInput, error) {
	in := &DayThreeInput{}

	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		row, err := parseDayThreeLine(line)
		if err != nil {
			return nil, err
		}

		in.Rows = append(in.Rows, row)
	}

	return in, nil
}

func solveDayThree(in *DayThreeInput) (*DayThreeOutput, error) {
	out := &DayThreeOutput{}
	// part 1
	sledPos := 0
	for _, row := range in.Rows {
		if row.HasTreeAt(sledPos) {
			out.NumTreesHit++
		}
		sledPos += 3
	}

	// part 2
	sleds := []*Sled{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	for i, row := range in.Rows {
		for _, sled := range sleds {
			if i%sled.Down == 0 {
				if row.HasTreeAt(sled.Pos) {
					sled.NumTreesHit++
				}
				sled.Pos += sled.Right

			}
		}
	}

	spew.Dump(sleds)

	out.MultiSledResult = 1
	for _, s := range sleds {
		out.MultiSledResult *= s.NumTreesHit
	}

	return out, nil
}

func DayThree(rawInput string) (*DayThreeOutput, error) {
	in, err := parseDayThree(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayThree(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
