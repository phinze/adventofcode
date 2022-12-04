package year2022

import (
	"bufio"
	"strings"

	"golang.org/x/exp/slices"
)

type Rucksack struct {
	Compartments []string
}

type Item rune

func (i Item) Priority() int {
	if 'a' <= i && i <= 'z' {
		return int(i - 'a' + 1)
	}
	if 'A' <= i && i <= 'Z' {
		return int(i - 'A' + 27)
	}
	return 0
}

func (r *Rucksack) CommonItem() Item {
	seenItems := map[rune][]int{}
	for compIndex, comp := range r.Compartments {
		for _, item := range comp {
			if !slices.Contains(seenItems[item], compIndex) {
				seenItems[item] = append(seenItems[item], compIndex)
			}
		}
	}

	for item, foundInCompartments := range seenItems {
		if len(foundInCompartments) == len(r.Compartments) {
			return Item(item)
		}
	}

	return 0
}

type ElfGroup struct {
	Rucksacks []*Rucksack
}

func (g *ElfGroup) CommonItem() Item {
	seenItems := map[rune][]int{}

	for ruckIndex, ruck := range g.Rucksacks {
		for _, comp := range ruck.Compartments {
			for _, item := range comp {
				if !slices.Contains(seenItems[item], ruckIndex) {
					seenItems[item] = append(seenItems[item], ruckIndex)
				}
			}
		}
	}

	// sanity check
	numCommon := 0
	for _, found := range seenItems {
		if len(found) == len(g.Rucksacks) {
			numCommon++
		}
	}

	if numCommon != 1 {
		panic("oh no")
	}

	for item, foundInRucksacks := range seenItems {
		if len(foundInRucksacks) == len(g.Rucksacks) {
			return Item(item)
		}
	}

	panic("no common items")
}

type DayThreeInput struct {
	Rucksacks []*Rucksack
	Groups    []*ElfGroup
}

type DayThreeOutput struct {
	PrioritySum int
	BadgeSum    int
}

func parseDayThree(rawInput string) (*DayThreeInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayThreeInput{}

	thisGroup := &ElfGroup{}
	for scanner.Scan() {
		line := scanner.Text()
		ruck := &Rucksack{
			Compartments: []string{
				line[0 : len(line)/2],
				line[len(line)/2:],
			},
		}
		in.Rucksacks = append(in.Rucksacks, ruck)
		thisGroup.Rucksacks = append(thisGroup.Rucksacks, ruck)
		if len(thisGroup.Rucksacks) == 3 {
			in.Groups = append(in.Groups, thisGroup)
			thisGroup = &ElfGroup{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayThree(in *DayThreeInput) (*DayThreeOutput, error) {
	out := &DayThreeOutput{}

	// part one
	for _, ruck := range in.Rucksacks {
		item := ruck.CommonItem()
		out.PrioritySum += item.Priority()
	}

	// part two
	for _, group := range in.Groups {
		item := group.CommonItem()
		out.BadgeSum += item.Priority()
	}

	return out, nil
}

func DayThree(rawInput string) (*DayThreeOutput, error) {
	in, err := parseDayThree(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayThree(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
