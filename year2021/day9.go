package year2021

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type DayNineInput struct {
	Readings [][]*lavaTubePoint
}

type lavaTubePoint struct {
	X, Y, Depth int
	Basin       *laveTubeBasin
}

func (l *lavaTubePoint) String() string {
	return fmt.Sprintf("[%d, %d: %d]", l.X, l.Y, l.Depth)
}

type laveTubeBasin struct {
	Points []*lavaTubePoint
}

func (in *DayNineInput) Adjacents(x, y int) []*lavaTubePoint {
	adjacents := make([]*lavaTubePoint, 0, 4)
	if x-1 >= 0 {
		adjacents = append(adjacents, in.Readings[y][x-1])
	}
	if y-1 >= 0 {
		adjacents = append(adjacents, in.Readings[y-1][x])
	}
	if x+1 < len(in.Readings[y]) {
		adjacents = append(adjacents, in.Readings[y][x+1])
	}
	if y+1 < len(in.Readings) {
		adjacents = append(adjacents, in.Readings[y+1][x])
	}

	return adjacents
}

type DayNineOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

func parseDayNine(rawInput string) (*DayNineInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayNineInput{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		row := make([]*lavaTubePoint, len(line))
		for i, c := range chars {
			n, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			row[i] = &lavaTubePoint{Y: y, X: i, Depth: n}
		}
		in.Readings = append(in.Readings, row)
		y++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayNine(in *DayNineInput) (*DayNineOutput, error) {
	out := &DayNineOutput{}

	riskLevel := 0
	lowPoints := make([]*lavaTubePoint, 0)
	for y := 0; y < len(in.Readings); y++ {
		for x := 0; x < len(in.Readings[y]); x++ {
			adjacents := in.Adjacents(x, y)

			point := in.Readings[y][x]
			lowerThanAllAdjacents := true
			for _, a := range adjacents {
				if point.Depth >= a.Depth {
					lowerThanAllAdjacents = false
				}
			}
			if lowerThanAllAdjacents {
				lowPoints = append(lowPoints, point)
				riskLevel += point.Depth + 1
			}
		}
	}
	spew.Dump(lowPoints)
	out.PartOneAnswer = riskLevel

	basins := []*laveTubeBasin{}
	for _, lowPoint := range lowPoints {
		// calculate baisin size from this lowpoint
		basin := &laveTubeBasin{[]*lavaTubePoint{lowPoint}}
		lowPoint.Basin = basin

		// start with adjacents
		searchSpace := in.Adjacents(lowPoint.X, lowPoint.Y)
		toCheck := len(searchSpace)

		for {
			for i := 0; i < len(searchSpace); i++ {
				p := searchSpace[i]
				toCheck--
				// we only touch each point once
				if p.Basin != nil {
					continue
				}
				if p.Depth < 9 {
					basin.Points = append(basin.Points, p)
					p.Basin = basin
					extras := in.Adjacents(p.X, p.Y)
					searchSpace = append(searchSpace, extras...)
					toCheck += len(extras)
				}
			}
			if toCheck == 0 {
				break
			}
		}

		basins = append(basins, basin)
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i].Points) > len(basins[j].Points)
	})

	out.PartTwoAnswer = len(basins[0].Points) * len(basins[1].Points) * len(basins[2].Points)

	return out, nil
}

func DayNine(rawInput string) (*DayNineOutput, error) {
	in, err := parseDayNine(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayNine(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
