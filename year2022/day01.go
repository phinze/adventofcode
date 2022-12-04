package year2022

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Calories int
}
type DayOneInput struct {
	Elves []*Elf
}
type DayOneOutput struct {
	MaxCalories      int
	MaxThreeCalories int
}

func parseDayOne(rawInput string) (*DayOneInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayOneInput{}

	// var thisElf *Elf
	thisElf := &Elf{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			in.Elves = append(in.Elves, thisElf)
			thisElf = &Elf{}
		} else {
			cals, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}
			thisElf.Calories += cals
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayOne(in *DayOneInput) (*DayOneOutput, error) {
	out := &DayOneOutput{}

	// part one
	for _, elf := range in.Elves {
		if elf.Calories > out.MaxCalories {
			out.MaxCalories = elf.Calories
		}
	}

	// part two
	allCals := make([]int, 0, len(in.Elves))
	for _, elf := range in.Elves {
		allCals = append(allCals, elf.Calories)
	}
	sort.Ints(allCals)
	for i := 1; i < 4; i++ {
		out.MaxThreeCalories += allCals[len(allCals)-i]
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
