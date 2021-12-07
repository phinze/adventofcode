package year2021

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type DaySixInput struct {
	FishFrequencies []int
}

type DaySixOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

func parseDaySix(rawInput string) (*DaySixInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DaySixInput{make([]int, 9)}

	for scanner.Scan() {
		timers := strings.Split(scanner.Text(), ",")
		for _, t := range timers {
			ti, err := strconv.Atoi(t)
			if err != nil {
				return nil, fmt.Errorf("noninteger timer value: %s:", t)
			}
			in.FishFrequencies[ti] += 1
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDaySix(in *DaySixInput) (*DaySixOutput, error) {
	out := &DaySixOutput{}

	day := 0
	for ; day < 80; day++ {
		spawners := in.FishFrequencies[0]
		for i := 1; i < len(in.FishFrequencies); i++ {
			in.FishFrequencies[i-1] = in.FishFrequencies[i]
		}
		in.FishFrequencies[6] += spawners
		in.FishFrequencies[8] = spawners
	}

	for _, i := range in.FishFrequencies {
		out.PartOneAnswer += i
	}

	for ; day < 256; day++ {
		spawners := in.FishFrequencies[0]
		for i := 1; i < len(in.FishFrequencies); i++ {
			in.FishFrequencies[i-1] = in.FishFrequencies[i]
		}
		in.FishFrequencies[6] += spawners
		in.FishFrequencies[8] = spawners
	}

	for _, i := range in.FishFrequencies {
		out.PartTwoAnswer += i
	}

	return out, nil
}

func DaySix(rawInput string) (*DaySixOutput, error) {
	in, err := parseDaySix(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDaySix(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
