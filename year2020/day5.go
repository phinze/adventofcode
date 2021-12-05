package year2020

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strings"
)

type boardingPass struct {
	Row int
	Col int
}

func (b *boardingPass) SeatID() int {
	return (b.Row * 8) + b.Col
}

func boardingPassFromCode(code string) (*boardingPass, error) {
	bp := &boardingPass{}

	low, high := 0, 127
	for i := 0; i < 7; i++ {
		switch code[i] {
		case 'F':
			high -= ((high - low) / 2) + 1
		case 'B':
			low += ((high - low) / 2) + 1
		default:
			return nil, fmt.Errorf("code %s unexpected char %s at position %d",
				code, string(code[i]), i)
		}
		log.Printf("hi: %d, low: %d", high, low)
	}
	if high != low {
		return nil, fmt.Errorf("code %s did not result in single row! high: %d, low %d",
			code, high, low)
	}
	bp.Row = high

	low, high = 0, 7
	for i := 7; i < len(code); i++ {
		switch code[i] {
		case 'L':
			high -= ((high - low) / 2) + 1
		case 'R':
			low += ((high - low) / 2) + 1
		default:
			return nil, fmt.Errorf("code %s unexpected char %s at position %d",
				code, string(code[i]), i)
		}
		log.Printf("hi: %d, low: %d", high, low)
	}
	if high != low {
		return nil, fmt.Errorf("code %s did not result in single col! high: %d, low %d",
			code, high, low)
	}
	bp.Col = high

	log.Printf("code %s gave seat %#v", code, bp)
	return bp, nil
}

type DayFiveInput struct {
	Passes []*boardingPass
}

type DayFiveOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

func parseDayFive(rawInput string) (*DayFiveInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayFiveInput{}

	for scanner.Scan() {
		pass, err := boardingPassFromCode(scanner.Text())
		if err != nil {
			return nil, err
		}
		in.Passes = append(in.Passes, pass)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayFive(in *DayFiveInput) (*DayFiveOutput, error) {
	out := &DayFiveOutput{}

	maxId := 0
	for _, pass := range in.Passes {
		id := pass.SeatID()
		if id > maxId {
			maxId = id
		}
	}
	out.PartOneAnswer = maxId

	ids := make([]int, len(in.Passes))
	for i, pass := range in.Passes {
		ids[i] = pass.SeatID()
	}
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] == 2 {
			out.PartTwoAnswer = ids[i] - 1
		}
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
