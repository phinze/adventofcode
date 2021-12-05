package year2021

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/bxcodec/saint"
)

type VentLine struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (v *VentLine) IsHorizontal() bool {
	return v.Y1 == v.Y2
}

func (v *VentLine) IsVertical() bool {
	return v.X1 == v.X2
}

func (v *VentLine) ContainsPoint(x3, y3 int) bool {
	inXRange := (v.X1 <= x3 && x3 <= v.X2) || (v.X2 <= x3 && x3 <= v.X1)
	inYRange := (v.Y1 <= y3 && y3 <= v.Y2) || (v.Y2 <= y3 && y3 <= v.Y1)
	if !(inXRange && inYRange) {
		return false
	}

	if v.IsHorizontal() {
		return y3 == v.Y1 && inXRange
	}
	if v.IsVertical() {
		return x3 == v.X1 && inYRange
	}

	// For three points, slope of any pair of points
	// must be same as other pair.
	//
	// For example, slope of line joining (x2, y2)
	// and (x3, y3), and line joining (x1, y1) and
	// (x2, y2) must be same.
	//
	// (y3 - y2)/(x3 - x2) = (y2 - y1)/(x2 - x1)
	//
	// In other words,
	// (y3 - y2)(x2 - x1) = (y2 - y1)(x3 - x2)
	return (y3-v.Y2)*(v.X2-v.X1) == (v.Y2-v.Y1)*(x3-v.X2)
}

func (v *VentLine) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", v.X1, v.Y1, v.X2, v.Y2)
}

type DayFiveInput struct {
	VentLines []*VentLine
	MaxX      int
	MaxY      int
}

func (i *DayFiveInput) AddVentLine(x1, y1, x2, y2 int) {
	line := &VentLine{x1, y1, x2, y2}
	i.VentLines = append(i.VentLines, line)
	i.MaxX = saint.Max(i.MaxX, line.X1, line.X2)
	i.MaxY = saint.Max(i.MaxY, line.Y1, line.Y2)
}

type DayFiveOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

var ventLinePattern = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

func parseDayFive(rawInput string) (*DayFiveInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayFiveInput{}

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("line: %s", line)
		matches := ventLinePattern.FindStringSubmatch(line)
		if len(matches) != 5 {
			return nil, fmt.Errorf("invalid format: %s, %#v", line, matches)
		}
		coords := make([]int, len(matches))
		for i, m := range matches[1:] {
			coord, err := strconv.Atoi(m)
			if err != nil {
				return nil, err
			}
			coords[i] = coord
		}
		in.AddVentLine(coords[0], coords[1], coords[2], coords[3])
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

type VentField struct {
	SizeX, SizeY int
	Positions    [][]int
}

func (f *VentField) String() string {
	b := strings.Builder{}
	for y := 0; y < f.SizeY; y++ {
		for x := 0; x < f.SizeX; x++ {
			val := f.Positions[x][y]
			if val == 0 {
				b.WriteString(".")
			} else {
				b.WriteString(fmt.Sprintf("%d", val))
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (f *VentField) NumPointsWithAtLeast(i int) int {
	num := 0
	for y := 0; y < f.SizeY; y++ {
		for x := 0; x < f.SizeX; x++ {
			if f.Positions[x][y] >= i {
				num++
			}
		}
	}
	return num
}

func NewVentField(sizeX, sizeY int) *VentField {
	v := &VentField{sizeX, sizeY, make([][]int, sizeX)}
	for i := 0; i < sizeX; i++ {
		v.Positions[i] = make([]int, sizeY)
	}
	return v
}

func solveDayFive(in *DayFiveInput) (*DayFiveOutput, error) {
	out := &DayFiveOutput{}

	field := NewVentField(in.MaxX+1, in.MaxY+1)
	fieldWithDiags := NewVentField(in.MaxX+1, in.MaxY+1)
	for y := 0; y < field.SizeY; y++ {
		for x := 0; x < field.SizeX; x++ {
			for _, line := range in.VentLines {
				if line.ContainsPoint(x, y) {
					fieldWithDiags.Positions[x][y]++
				}
				if line.IsHorizontal() || line.IsVertical() {
					if line.ContainsPoint(x, y) {
						field.Positions[x][y]++
					}
				}
			}
		}
	}

	out.PartOneAnswer = field.NumPointsWithAtLeast(2)
	out.PartTwoAnswer = fieldWithDiags.NumPointsWithAtLeast(2)

	fmt.Printf("FIELD\n%s\n", field)
	fmt.Printf("FIELD WITH DIAGS\n%s\n", fieldWithDiags)

	return out, nil
}

func DayFive(rawInput string) (*DayFiveOutput, error) {
	in, err := parseDayFive(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayFive(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
