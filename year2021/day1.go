package year2021

import (
	"bufio"
	"strconv"
	"strings"
)

type DayOneInput struct {
	Depths []int
}
type DayOneOutput struct {
	TotalReadings    int
	TimesIncreased   int
	WindowsIncreased int
}

func parseDayOne(rawInput string) (*DayOneInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayOneInput{}

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		in.Depths = append(in.Depths, depth)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayOne(in *DayOneInput) (*DayOneOutput, error) {
	out := &DayOneOutput{
		TotalReadings: len(in.Depths),
	}
	// part one
	for i := 1; i < len(in.Depths); i++ {
		if in.Depths[i] > in.Depths[i-1] {
			out.TimesIncreased++
		}
	}
	// part two
	windows := []int{}
	for i := 2; i < len(in.Depths); i++ {
		windows = append(windows,
			in.Depths[i-2]+in.Depths[i-1]+in.Depths[i],
		)
	}
	for i := 1; i < len(windows); i++ {
		if windows[i] > windows[i-1] {
			out.WindowsIncreased++
		}
	}

	return out, nil
}

func DayOne(rawInput string) (*DayOneOutput, error) {
	in, err := parseDayOne(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayOne(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
