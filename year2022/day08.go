package year2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"tailscale.com/util/must"
)

type DayEightInput struct {
	Grid *TreeGrid
}
type DayEightOutput struct {
	VisibleTrees   int
	MaxScenicScore int
}

type Direction int

const (
	Top    Direction = iota
	Right            = iota
	Bottom           = iota
	Left             = iota
)

var Directions = []Direction{Top, Right, Bottom, Left}

type TreeGrid struct {
	Trees [][]*Tree
}

func (tg *TreeGrid) Width() int {
	return len(tg.Trees[0])
}

func (tg *TreeGrid) Height() int {
	return len(tg.Trees)
}

func (tg *TreeGrid) VisibilityCheck() int {
	// From Top
	for x := 0; x < tg.Width(); x++ {
		maxSeen := -1
		for y := 0; y < tg.Height(); y++ {
			t := tg.Trees[y][x]
			if t.Height > maxSeen {
				t.VisibleFrom = append(t.VisibleFrom, Top)
				maxSeen = t.Height
			}
		}
	}
	// From Bottom
	for x := 0; x < tg.Width(); x++ {
		maxSeen := -1
		for y := tg.Height() - 1; y >= 0; y-- {
			t := tg.Trees[y][x]
			if t.Height > maxSeen {
				t.VisibleFrom = append(t.VisibleFrom, Bottom)
				maxSeen = t.Height
			}
		}
	}
	// From Left
	for y := 0; y < tg.Height(); y++ {
		maxSeen := -1
		for x := 0; x < tg.Width(); x++ {
			t := tg.Trees[y][x]
			if t.Height > maxSeen {
				t.VisibleFrom = append(t.VisibleFrom, Left)
				maxSeen = t.Height
			}
		}
	}
	// From Right
	for y := 0; y < tg.Height(); y++ {
		maxSeen := -1
		for x := tg.Width() - 1; x >= 0; x-- {
			t := tg.Trees[y][x]
			if t.Height > maxSeen {
				t.VisibleFrom = append(t.VisibleFrom, Right)
				maxSeen = t.Height
			}
		}
	}

	totalVisible := 0
	for x := 0; x < tg.Width(); x++ {
		for y := 0; y < tg.Height(); y++ {
			if tg.Trees[y][x].IsVisible() {
				totalVisible++
			}
		}
	}
	return totalVisible
}

func (tg *TreeGrid) ScenicCheck() int {
	for x := 0; x < tg.Width(); x++ {
		for y := 0; y < tg.Height(); y++ {
			t := tg.Trees[y][x]
			// Up
			for i := y - 1; i >= 0; i-- {
				t2 := tg.Trees[i][x]
				t.VisionRange[Top]++
				if t2.Height >= t.Height {
					break
				}
			}
			// Down
			for i := y + 1; i < tg.Height(); i++ {
				t2 := tg.Trees[i][x]
				t.VisionRange[Bottom]++
				if t2.Height >= t.Height {
					break
				}
			}
			// Right
			for i := x + 1; i < tg.Width(); i++ {
				t2 := tg.Trees[y][i]
				t.VisionRange[Right]++
				if t2.Height >= t.Height {
					break
				}
			}
			// Left
			for i := x - 1; i >= 0; i-- {
				t2 := tg.Trees[y][i]
				t.VisionRange[Left]++
				if t2.Height >= t.Height {
					break
				}
			}
		}
	}

	maxScore := 0
	for x := 0; x < tg.Width(); x++ {
		for y := 0; y < tg.Height(); y++ {
			score := tg.Trees[y][x].ScenicScore()
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func (tg *TreeGrid) String() string {
	b := strings.Builder{}
	for _, row := range tg.Trees {
		for _, t := range row {
			b.WriteString(t.String() + " ")
		}
		b.WriteString("\n")
	}
	return b.String()
}

type Tree struct {
	Height      int
	VisibleFrom []Direction
	VisionRange map[Direction]int
}

func (t *Tree) IsVisible() bool {
	return len(t.VisibleFrom) > 0
}

func (t *Tree) ScenicScore() int {
	score := 1
	for _, d := range Directions {
		score *= t.VisionRange[d]
	}
	return score
}

func (t *Tree) String() string {
	return fmt.Sprintf("%d[%d, %d]", t.Height, len(t.VisibleFrom), t.ScenicScore())
}

func parseDayEight(rawInput string) (*DayEightInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayEightInput{
		Grid: &TreeGrid{},
	}

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]*Tree, len(line))
		in.Grid.Trees = append(in.Grid.Trees, row)
		for i, c := range line {
			row[i] = &Tree{
				Height:      must.Get(strconv.Atoi(string(c))),
				VisionRange: make(map[Direction]int),
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayEight(in *DayEightInput) (*DayEightOutput, error) {
	out := &DayEightOutput{}

	// part one
	out.VisibleTrees = in.Grid.VisibilityCheck()

	// part two
	out.MaxScenicScore = in.Grid.ScenicCheck()

	return out, nil
}

func DayEight(rawInput string) (*DayEightOutput, error) {
	in, err := parseDayEight(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayEight(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
