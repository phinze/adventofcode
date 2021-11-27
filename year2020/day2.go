package year2020

import (
	"fmt"
	"strconv"
	"strings"
)

type PasswordWithSpec struct {
	Min      int
	Max      int
	Letter   rune
	Password string
}

func (p *PasswordWithSpec) IsValid() bool {
	letterCount := strings.Count(p.Password, string(p.Letter))
	return letterCount >= p.Min && letterCount <= p.Max
}

func (p *PasswordWithSpec) IsValid2() bool {
	inPosOne := rune(p.Password[p.Min-1]) == p.Letter
	inPosTwo := rune(p.Password[p.Max-1]) == p.Letter

	return (inPosOne || inPosTwo) && !(inPosOne && inPosTwo)
}

type DayTwoInput struct {
	Passwords []*PasswordWithSpec
}

type DayTwoOutput struct {
	NumValid  int
	NumValid2 int
}

func parseDayTwoLine(line string) (*PasswordWithSpec, error) {
	pw := &PasswordWithSpec{}

	fields := strings.Split(line, " ")
	if len(fields) != 3 {
		return nil, fmt.Errorf("Expected line to have 3 fields, got %d: %v", len(fields), line)
	}

	occurRange := strings.Split(fields[0], "-")
	if len(occurRange) != 2 {
		return nil, fmt.Errorf("Expected line to have 3 fields, got %d: %v", len(occurRange), line)
	}
	var err error

	pw.Min, err = strconv.Atoi(occurRange[0])
	if err != nil {
		return nil, fmt.Errorf("Error converting min to num: %s", err)
	}

	pw.Max, err = strconv.Atoi(occurRange[1])
	if err != nil {
		return nil, fmt.Errorf("Error converting max to num: %s", err)
	}

	pw.Letter = rune(fields[1][0])

	pw.Password = fields[2]

	return pw, nil
}

func parseDayTwo(input string) (*DayTwoInput, error) {
	out := &DayTwoInput{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		pw, err := parseDayTwoLine(line)
		if err != nil {
			return nil, err
		}
		out.Passwords = append(out.Passwords, pw)
	}
	return out, nil
}

func solveDayTwo(in *DayTwoInput) (*DayTwoOutput, error) {
	out := &DayTwoOutput{}

	for _, pw := range in.Passwords {
		if pw.IsValid() {
			out.NumValid++
		}
		if pw.IsValid2() {
			out.NumValid2++
		}
	}

	return out, nil
}

func DayTwo(input string) (*DayTwoOutput, error) {
	in, err := parseDayTwo(input)
	if err != nil {
		return nil, err
	}

	out, err := solveDayTwo(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
