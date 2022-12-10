package year2022

import (
	"bufio"
	"fmt"
	"log"
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

type Coords struct {
	X, Y int
}

func (c *Coords) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type RopeField struct {
	Head    *Coords
	Tail    *Coords
	History map[string]int
	maxX    int
	maxY    int
}

func (rf *RopeField) String() string {
	var b strings.Builder
	for y := rf.maxY; y >= 0; y-- {
		for x := 0; x <= rf.maxX; x++ {
			if rf.Head.X == x && rf.Head.Y == y {
				b.WriteString("H")
			} else if rf.Tail.X == x && rf.Tail.Y == y {
				b.WriteString("T")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (rf *RopeField) Execute(m *RopeMove) {
	log.Printf("executing move: %s", m)
	for i := 0; i < m.Steps; i++ {
		switch m.Direction {
		case "U":
			rf.Head.Y += 1
		case "D":
			rf.Head.Y -= 1
		case "R":
			rf.Head.X += 1
		case "L":
			rf.Head.X -= 1
		}
		if rf.Head.X > rf.maxX {
			rf.maxX = rf.Head.X
		}
		if rf.Head.Y > rf.maxY {
			rf.maxY = rf.Head.Y
		}
		rf.UpdateTail()
		fmt.Printf("%s\n", rf)
	}
	log.Printf("head now: %s", rf.Head)
	log.Printf("tail now: %s", rf.Tail)
}

func (rf *RopeField) UpdateTail() {
	// If the head is ever two steps directly up, down, left, or right from the
	// tail, the tail must also move one step in that direction
	deltaX := rf.Head.X - rf.Tail.X
	deltaY := rf.Head.Y - rf.Tail.Y
	fmt.Printf("dX: %d, dY: %d\n", deltaX, deltaY)
	if deltaX == 0 {
		if deltaY > 1 {
			fmt.Println("--> moving up")
			rf.Tail.Y += 1
		} else if deltaY < -1 {
			fmt.Println("--> moving down")
			rf.Tail.Y -= 1
		}
	} else if deltaY == 0 {
		if deltaX > 1 {
			fmt.Println("--> moving right")
			rf.Tail.X += 1
		} else if deltaX < -1 {
			fmt.Println("--> moving left")
			rf.Tail.X -= 1
		}
	} else if (deltaY > 1 && deltaX > 0) || (deltaX > 1 && deltaY > 0) {
		fmt.Println("--> moving up/right")
		rf.Tail.X += 1
		rf.Tail.Y += 1
	} else if (deltaY < -1 && deltaX < 0) || (deltaX < -1 && deltaY < 0) {
		fmt.Println("--> moving down/left")
		rf.Tail.X -= 1
		rf.Tail.Y -= 1
	} else if (deltaX < -1 && deltaY > 0) || (deltaY > 1 && deltaX < 0) {
		fmt.Println("--> moving up/left")
		rf.Tail.X -= 1
		rf.Tail.Y += 1
	} else if (deltaX > 1 && deltaY < 0) || (deltaY < -1 && deltaX > 0) {
		fmt.Println("--> moving down/right")
		rf.Tail.X += 1
		rf.Tail.Y -= 1
	}
	rf.History[fmt.Sprintf("%d,%d", rf.Tail.X, rf.Tail.Y)] += 1
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
	field := &RopeField{
		Head:    &Coords{},
		Tail:    &Coords{},
		History: make(map[string]int),
	}
	for _, m := range in.Moves {
		field.Execute(m)
		log.Printf("board:\n%s", field)
	}
	out.VisitedPositions = len(field.History)

	// part two

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
