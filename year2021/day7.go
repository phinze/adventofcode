package year2021

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/bxcodec/saint"
)

type DaySevenInput struct {
	CrabLocations []int
	Min, Max      int
}

type DaySevenOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

func parseDaySeven(rawInput string) (*DaySevenInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DaySevenInput{}

	for scanner.Scan() {
		locs := strings.Split(scanner.Text(), ",")
		for _, l := range locs {
			li, err := strconv.Atoi(l)
			if err != nil {
				return nil, err
			}
			in.CrabLocations = append(in.CrabLocations, li)
			in.Min = saint.Min(in.Min, li)
			in.Max = saint.Max(in.Max, li)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDaySeven(in *DaySevenInput) (*DaySevenOutput, error) {
	out := &DaySevenOutput{}

	out.PartOneAnswer = math.MaxInt64
	for pos := in.Min; pos <= in.Max; pos++ {
		fuelCost := 0
		for _, loc := range in.CrabLocations {
			fuelCost += saint.Abs(pos - loc)
		}
		out.PartOneAnswer = saint.Min(out.PartOneAnswer, fuelCost)
	}

	out.PartTwoAnswer = math.MaxInt64
	for pos := in.Min; pos <= in.Max; pos++ {
		fuelCost := 0
		for _, loc := range in.CrabLocations {
			// fuel cost is Nth triangle number
			// (n * (n + 1)) / 2
			base := saint.Abs(pos - loc)
			fuelCost += (base * (base + 1)) / 2
		}
		out.PartTwoAnswer = saint.Min(out.PartTwoAnswer, fuelCost)
	}

	return out, nil
}

func DaySeven(rawInput string) (*DaySevenOutput, error) {
	in, err := parseDaySeven(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDaySeven(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
