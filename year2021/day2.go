package year2021

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type dayTwoCommand struct {
	Direction string
	Magnitude int
}

type DayTwoInput struct {
	Commands []*dayTwoCommand
}

type DayTwoOutput struct {
	Position      int
	Depth         int
	PartOneAnswer int

	PartTwoPosition int
	PartTwoDepth    int
	PartTwoAim      int
	PartTwoAnswer   int
}

func parseDayTwo(rawInput string) (*DayTwoInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayTwoInput{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		mag, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		in.Commands = append(in.Commands, &dayTwoCommand{
			Direction: fields[0],
			Magnitude: mag,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayTwo(in *DayTwoInput) (*DayTwoOutput, error) {
	out := &DayTwoOutput{}

	// part one
	for _, c := range in.Commands {
		switch c.Direction {
		case "forward":
			out.Position += c.Magnitude
		case "down":
			out.Depth += c.Magnitude
		case "up":
			out.Depth -= c.Magnitude
		default:
			return nil, fmt.Errorf("unexpected direction: %v", c.Direction)
		}
	}

	// part two
	for _, c := range in.Commands {
		switch c.Direction {
		case "forward":
			out.PartTwoPosition += c.Magnitude
			out.PartTwoDepth += out.PartTwoAim * c.Magnitude
		case "down":
			out.PartTwoAim += c.Magnitude
		case "up":
			out.PartTwoAim -= c.Magnitude
		default:
			return nil, fmt.Errorf("unexpected direction: %v", c.Direction)
		}
	}

	out.PartOneAnswer = out.Depth * out.Position
	out.PartTwoAnswer = out.PartTwoPosition * out.PartTwoDepth

	return out, nil
}

func DayTwo(rawInput string) (*DayTwoOutput, error) {
	in, err := parseDayTwo(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayTwo(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
