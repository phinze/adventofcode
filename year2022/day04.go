package year2022

import (
	"bufio"
	"strconv"
	"strings"
)

type CleaningAssignment struct {
	Start int
	End   int
}

func (ca *CleaningAssignment) FullyContains(other *CleaningAssignment) bool {
	return ca.Start <= other.Start && ca.End >= other.End
}

func (ca *CleaningAssignment) Overlaps(other *CleaningAssignment) bool {
	return ca.End >= other.Start && other.End >= ca.Start
}

func NewCleaningAssignment(ran string) *CleaningAssignment {
	ca := &CleaningAssignment{}
	parts := strings.Split(ran, "-")
	var err error
	if ca.Start, err = strconv.Atoi(parts[0]); err != nil {
		panic(err)
	}
	if ca.End, err = strconv.Atoi(parts[1]); err != nil {
		panic(err)
	}
	return ca
}

type CleaningElfGroup struct {
	Assignments []*CleaningAssignment
}

type DayFourInput struct {
	Groups []*CleaningElfGroup
}

type DayFourOutput struct {
	FullyContainsCount int
	OverlapsCount      int
}

func parseDayFour(rawInput string) (*DayFourInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayFourInput{}

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		group := &CleaningElfGroup{}
		for _, a := range assignments {
			group.Assignments = append(group.Assignments, NewCleaningAssignment(a))
		}
		in.Groups = append(in.Groups, group)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayFour(in *DayFourInput) (*DayFourOutput, error) {
	out := &DayFourOutput{}

	// part one
	for _, g := range in.Groups {
		left := g.Assignments[0]
		right := g.Assignments[1]

		if left.FullyContains(right) || right.FullyContains(left) {
			out.FullyContainsCount++
		}
	}

	// part two
	for _, g := range in.Groups {
		left := g.Assignments[0]
		right := g.Assignments[1]

		if left.Overlaps(right) {
			out.OverlapsCount++
		}
	}

	return out, nil
}

func DayFour(rawInput string) (*DayFourOutput, error) {
	in, err := parseDayFour(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayFour(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
