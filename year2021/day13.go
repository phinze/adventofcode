package year2021

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type transparencyDot struct {
	X, Y int
}

func (d *transparencyDot) String() string {
	return fmt.Sprintf("(%d,%d)", d.X, d.Y)
}

func (d *transparencyDot) Equal(other *transparencyDot) bool {
	return other.X == d.X && other.Y == d.Y
}

type transparencyFold struct {
	Axis  string
	Coord int
}

type DayThirteenInput struct {
	Dots  []*transparencyDot
	Folds []*transparencyFold

	maxX, maxY int
}

func (in *DayThirteenInput) AddDot(x, y int) {
	in.Dots = append(in.Dots, &transparencyDot{X: x, Y: y})
	if x > in.maxX {
		in.maxX = x
	}
	if y > in.maxY {
		in.maxY = y
	}
}

func (in *DayThirteenInput) RemoveDuplicateDots() {
	newDots := make([]*transparencyDot, 0, len(in.Dots))
	var newMaxX, newMaxY int
	for _, oldDot := range in.Dots {
		duped := false
		for _, newDot := range newDots {
			if newDot.Equal(oldDot) {
				duped = true
			}
		}
		if !duped {
			newDots = append(newDots, oldDot)
			if oldDot.X > newMaxX {
				newMaxX = oldDot.X
			}
			if oldDot.Y > newMaxY {
				newMaxY = oldDot.Y
			}
		}
	}
	in.Dots = newDots
	in.maxX = newMaxX
	in.maxY = newMaxY
}

func (in *DayThirteenInput) HasDot(x, y int) bool {
	for _, d := range in.Dots {
		if d.X == x && d.Y == y {
			return true
		}
	}
	return false
}

func (in *DayThirteenInput) String() string {
	b := strings.Builder{}
	for y := 0; y <= in.maxY; y++ {
		for x := 0; x <= in.maxX; x++ {
			if in.HasDot(x, y) {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

type DayThirteenOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

func parseDayThirteen(rawInput string) (*DayThirteenInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayThirteenInput{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			dotCoords := strings.Split(line, ",")
			x, err := strconv.Atoi(dotCoords[0])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(dotCoords[1])
			if err != nil {
				return nil, err
			}
			in.AddDot(x, y)
		}

		if strings.Contains(line, "fold") {
			foldRegexp := regexp.MustCompile(`^fold along ([xy])=(\d+)$`)
			matches := foldRegexp.FindStringSubmatch(line)
			coord, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, err
			}
			in.Folds = append(in.Folds, &transparencyFold{Axis: matches[1], Coord: coord})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayThirteen(in *DayThirteenInput) (*DayThirteenOutput, error) {
	out := &DayThirteenOutput{}

	for _, fold := range in.Folds {
		switch fold.Axis {
		case "y":
			// hamburger: performing vertical fold with a horizontal line
			for _, dot := range in.Dots {
				if dot.Y > fold.Coord {
					delta := (dot.Y - fold.Coord) * 2
					dot.Y -= delta
				}
			}
		case "x":
			// hotdot: performing horizontal fold with a vertical line
			for _, dot := range in.Dots {
				if dot.X > fold.Coord {
					delta := (dot.X - fold.Coord) * 2
					dot.X -= delta
				}
			}
		}

		in.RemoveDuplicateDots()

		// part 1 gets first fold dot count
		if out.PartOneAnswer == 0 {
			out.PartOneAnswer = len(in.Dots)
		}
	}

	fmt.Printf("%s", in)

	return out, nil
}

func DayThirteen(rawInput string) (*DayThirteenOutput, error) {
	in, err := parseDayThirteen(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayThirteen(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
