package year2020

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parse(input string) ([]int, error) {
	var out []int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

func solve(nums []int) (int, error) {
	for _, x := range nums {
		for _, y := range nums {
			if x+y == 2020 {
				return x * y, nil
			}
		}
	}
	return 0, errors.New("not found")
}

func solvePartTwo(nums []int) (int, error) {
	for _, x := range nums {
		for _, y := range nums {
			for _, z := range nums {
				if x+y+z == 2020 {
					return x * y * z, nil
				}
			}
		}
	}
	return 0, errors.New("not found")
}

func DayOne(input string) (map[string]string, error) {
	out := make(map[string]string)

	in, err := parse(input)
	if err != nil {
		return nil, err
	}

	solution, err := solve(in)
	if err != nil {
		return nil, err
	}
	out["part1"] = fmt.Sprint(solution)

	solution, err = solvePartTwo(in)
	if err != nil {
		return nil, err
	}
	out["part2"] = fmt.Sprint(solution)

	return out, nil
}
