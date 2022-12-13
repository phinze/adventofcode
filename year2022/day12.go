package year2022

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/dominikbraun/graph"
	"tailscale.com/util/must"
)

type MapPoint struct {
	X, Y      int
	Elevation rune

	pathTaken rune
}

func (mp *MapPoint) Height() int {
	switch mp.Elevation {
	case 'S':
		return 0
	case 'E':
		return 26
	default:
		return int(mp.Elevation - 'a')
	}
}

func (mp *MapPoint) IsStart() bool {
	return mp.Elevation == 'S'
}

func (mp *MapPoint) IsEnd() bool {
	return mp.Elevation == 'E'
}

func (mp *MapPoint) String() string {
	if mp.pathTaken != 0 {
		return fmt.Sprintf("%c", mp.pathTaken)
	} else {
		return fmt.Sprintf("%c", mp.Elevation)
	}
}

func (mp *MapPoint) PathTakenTo(next *MapPoint) {
	if mp.X+1 == next.X {
		mp.pathTaken = '>'
	} else if mp.X-1 == next.X {
		mp.pathTaken = '<'
	} else if mp.Y+1 == next.Y {
		mp.pathTaken = 'v'
	} else if mp.Y-1 == next.Y {
		mp.pathTaken = '^'
	}
}

type HeightMap struct {
	Grid  [][]*MapPoint
	Start *MapPoint
	End   *MapPoint
}

func (m *HeightMap) Width() int {
	return len(m.Grid[0])
}

func (m *HeightMap) Height() int {
	return len(m.Grid)
}

func (m *HeightMap) Neighbors(x, y int) []*MapPoint {
	n := make([]*MapPoint, 0, 4)
	if x-1 >= 0 {
		n = append(n, m.Grid[y][x-1])
	}
	if x+1 < m.Width() {
		n = append(n, m.Grid[y][x+1])
	}
	if y-1 >= 0 {
		n = append(n, m.Grid[y-1][x])
	}
	if y+1 < m.Height() {
		n = append(n, m.Grid[y+1][x])
	}
	return n
}

func (m *HeightMap) String() string {
	var b strings.Builder

	for _, row := range m.Grid {
		for _, p := range row {
			b.WriteString(p.String())
		}
		b.WriteString("\n")
	}

	return b.String()
}

type DayTwelveInput struct {
	Map *HeightMap
}

type DayTwelveOutput struct {
	PartOneAnswer int
}

func parsePoint(x, y int, c rune) *MapPoint {
	return &MapPoint{
		X:         x,
		Y:         y,
		Elevation: c,
	}
}

func pointHash(p *MapPoint) string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func parseDayTwelve(rawInput string) (*DayTwelveInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayTwelveInput{
		Map: &HeightMap{},
	}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []*MapPoint{}
		for x, c := range line {
			p := parsePoint(x, y, c)
			if p.IsStart() {
				in.Map.Start = p
			} else if p.IsEnd() {
				in.Map.End = p
			}
			row = append(row, p)
		}
		in.Map.Grid = append(in.Map.Grid, row)
		log.Printf("line: %s", line)
		y++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayTwelve(in *DayTwelveInput) (*DayTwelveOutput, error) {
	out := &DayTwelveOutput{}

	// part one
	g := graph.New(pointHash, graph.Directed())

	// Make points
	for _, row := range in.Map.Grid {
		for _, point := range row {
			g.AddVertex(point)
		}
	}

	// Draw traversible edges
	for y, row := range in.Map.Grid {
		for x, point := range row {
			// Draw edges for traverse-able neighbors
			// log.Printf("at %d,%d[%d]", x, y, point.Height())
			for _, n := range in.Map.Neighbors(x, y) {
				// log.Printf("    considering %d,%d", n.X, n.Y)
				if n.Height()-point.Height() <= 1 {
					// log.Printf("  %s -> %s", point, n)
					g.AddEdge(pointHash(point), pointHash(n))
				}
			}
		}
	}

	// Find path
	path, err := graph.ShortestPath(g, pointHash(in.Map.Start), pointHash(in.Map.End))
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(path)-1; i++ {
		this := must.Get(g.Vertex(path[i]))
		next := must.Get(g.Vertex(path[i+1]))
		this.PathTakenTo(next)
	}
	log.Printf("path:\n%s", in.Map)

	// Draw path

	log.Printf("path: %s", path)

	out.PartOneAnswer = len(path) - 1

	// Calculate shortest path

	// part two

	return out, nil
}

func DayTwelve(rawInput string) (*DayTwelveOutput, error) {
	in, err := parseDayTwelve(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayTwelve(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
