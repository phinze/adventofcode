package year2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"tailscale.com/util/must"
)

type RopeMove struct {
	Direction string
	Steps     int
}

func (rm *RopeMove) String() string {
	return fmt.Sprintf("%s %d", rm.Direction, rm.Steps)
}

type DayNineInput struct {
	Moves []*RopeMove
}

type Knot struct {
	X, Y int
}

type Rope struct {
	Knots []*Knot
}

func (c *Knot) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type RopeField struct {
	Rope    []*Knot
	History map[string]int
	minX    int
	minY    int
	maxX    int
	maxY    int
}

func NewRopeField(ropeLength int) *RopeField {
	rf := &RopeField{}
	rf.History = make(map[string]int)
	rf.Rope = make([]*Knot, ropeLength)
	for i := 0; i < ropeLength; i++ {
		rf.Rope[i] = &Knot{}
	}
	return rf
}

func (rf *RopeField) String() string {
	var b strings.Builder
	for y := rf.maxY; y >= rf.minY; y-- {
		for x := rf.minX; x <= rf.maxX; x++ {
			ropeHere := false
			for i, k := range rf.Rope {
				if k.X == x && k.Y == y {
					ropeHere = true
					if i == 0 {
						b.WriteString("H")
					} else {
						b.WriteString(fmt.Sprintf("%d", i))
					}
					break
				}
			}
			if !ropeHere {
				if x == 0 && y == 0 {
					b.WriteString("s")
				} else {
					b.WriteString(".")
				}
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (rf *RopeField) Execute(m *RopeMove) {
	for i := 0; i < m.Steps; i++ {
		switch m.Direction {
		case "U":
			rf.Rope[0].Y += 1
		case "D":
			rf.Rope[0].Y -= 1
		case "R":
			rf.Rope[0].X += 1
		case "L":
			rf.Rope[0].X -= 1
		}
		if rf.Rope[0].X > rf.maxX {
			rf.maxX = rf.Rope[0].X
		}
		if rf.Rope[0].Y > rf.maxY {
			rf.maxY = rf.Rope[0].Y
		}
		if rf.Rope[0].X < rf.minX {
			rf.minX = rf.Rope[0].X
		}
		if rf.Rope[0].Y < rf.minY {
			rf.minY = rf.Rope[0].Y
		}
		rf.UpdateTail()
	}
}

func (rf *RopeField) UpdateTail() {
	for i := 1; i < len(rf.Rope); i++ {
		prev := rf.Rope[i-1]
		this := rf.Rope[i]
		// If the head is ever two steps directly up, down, left, or right from the
		// tail, the tail must also move one step in that direction
		deltaX := prev.X - this.X
		deltaY := prev.Y - this.Y
		if deltaX == 0 {
			if deltaY > 1 {
				// fmt.Println("--> moving up")
				this.Y += 1
			} else if deltaY < -1 {
				// fmt.Println("--> moving down")
				this.Y -= 1
			}
		} else if deltaY == 0 {
			if deltaX > 1 {
				// fmt.Println("--> moving right")
				this.X += 1
			} else if deltaX < -1 {
				// fmt.Println("--> moving left")
				this.X -= 1
			}
		} else if (deltaY > 1 && deltaX > 0) || (deltaX > 1 && deltaY > 0) {
			// fmt.Println("--> moving up/right")
			this.X += 1
			this.Y += 1
		} else if (deltaY < -1 && deltaX < 0) || (deltaX < -1 && deltaY < 0) {
			// fmt.Println("--> moving down/left")
			this.X -= 1
			this.Y -= 1
		} else if (deltaX < -1 && deltaY > 0) || (deltaY > 1 && deltaX < 0) {
			// fmt.Println("--> moving up/left")
			this.X -= 1
			this.Y += 1
		} else if (deltaX > 1 && deltaY < 0) || (deltaY < -1 && deltaX > 0) {
			// fmt.Println("--> moving down/right")
			this.X += 1
			this.Y -= 1
		}
	}
	tail := rf.Rope[len(rf.Rope)-1]
	rf.History[fmt.Sprintf("%d,%d", tail.X, tail.Y)] += 1
}

type DayNineOutput struct {
	VisitedPositions        int
	BigRopeVisitedPositions int
}

func parseDayNine(rawInput string) (*DayNineInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayNineInput{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		in.Moves = append(in.Moves, &RopeMove{
			Direction: tokens[0],
			Steps:     must.Get(strconv.Atoi(tokens[1])),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayNine(in *DayNineInput) (*DayNineOutput, error) {
	out := &DayNineOutput{}

	// part one
	field := NewRopeField(2)
	for _, m := range in.Moves {
		field.Execute(m)
	}
	out.VisitedPositions = len(field.History)

	// part two
	field = NewRopeField(10)
	for _, m := range in.Moves {
		field.Execute(m)
	}
	out.BigRopeVisitedPositions = len(field.History)

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
