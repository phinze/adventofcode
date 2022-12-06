package year2022

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/copystructure"
	"tailscale.com/util/must"
)

type Crate struct {
	Label string
}

type Stack struct {
	Number int
	Crates []*Crate
}

func (s *Stack) Push(c *Crate) {
	s.Crates = append(s.Crates, c)
}

func (s *Stack) PushSet(set []*Crate) {
	s.Crates = append(s.Crates, set...)
}

func (s *Stack) Pop() *Crate {
	c := s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[:len(s.Crates)-1]
	return c
}

func (s *Stack) PopN(n int) []*Crate {
	set := make([]*Crate, n)
	copy(set, s.Crates[len(s.Crates)-n:])
	s.Crates = s.Crates[:len(s.Crates)-n]
	return set
}

func (s *Stack) TopCrate() *Crate {
	return s.Crates[len(s.Crates)-1]
}

type Move struct {
	Num  int
	From int
	To   int
}

type DayFiveInput struct {
	Stacks []*Stack
	Moves  []*Move
}
type DayFiveOutput struct {
	PartOneMessage string
	PartTwoMessage string
}

func printStacks(stacks []*Stack) {
	for _, stack := range stacks {
		fmt.Printf("%d: ", stack.Number)
		for _, crate := range stack.Crates {
			fmt.Printf("%s ", crate.Label)
		}
		fmt.Print("\n")
	}
}

func parseDayFive(rawInput string) (*DayFiveInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayFiveInput{}

	stackLines := []string{}
	var stackNumbers []string
	moveLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "[") {
			stackLines = append(stackLines, line)
			log.Printf("stackLines: %s", line)
		} else if strings.HasPrefix(line, " 1") {
			stackNumbers = strings.Fields(line)
			log.Printf("stackNumbers: %s", line)
		} else if strings.HasPrefix(line, "move") {
			moveLines = append(moveLines, line)
		}
	}

	// Initialize stacks
	in.Stacks = make([]*Stack, len(stackNumbers))
	for i, stackNum := range stackNumbers {
		in.Stacks[i] = &Stack{
			Number: Atoi(stackNum),
			Crates: []*Crate{},
		}
	}

	// Place crates
	for i := len(stackLines) - 1; i >= 0; i-- {
		stackLine := stackLines[i]
		for j, stack := range in.Stacks {
			stackIndex := 1 + (j * 4)
			crateSlot := string(stackLine[stackIndex])
			if crateSlot != " " {
				stack.Push(&Crate{Label: crateSlot})
			}
		}
	}

	// Parse moves
	for _, moveLine := range moveLines {
		moveFields := strings.Fields(moveLine)
		in.Moves = append(in.Moves, &Move{
			Num:  Atoi(moveFields[1]),
			From: Atoi(moveFields[3]),
			To:   Atoi(moveFields[5]),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayFive(in *DayFiveInput) (*DayFiveOutput, error) {
	out := &DayFiveOutput{}

	// part one
	stacks := must.Get(copystructure.Copy(in.Stacks)).([]*Stack)
	for _, move := range in.Moves {
		for i := 0; i < move.Num; i++ {
			stacks[move.To-1].Push(stacks[move.From-1].Pop())
		}
	}

	for _, stack := range stacks {
		out.PartOneMessage = out.PartOneMessage + stack.TopCrate().Label
	}

	printStacks(in.Stacks)

	// part two
	stacks = must.Get(copystructure.Copy(in.Stacks)).([]*Stack)
	for _, move := range in.Moves {
		log.Printf("doing move: %#v", move)
		stacks[move.To-1].PushSet(stacks[move.From-1].PopN(move.Num))
		printStacks(stacks)
	}

	for _, stack := range stacks {
		out.PartTwoMessage = out.PartTwoMessage + stack.TopCrate().Label
	}

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

func Atoi(s string) int {
	return must.Get(strconv.Atoi(s))
}
