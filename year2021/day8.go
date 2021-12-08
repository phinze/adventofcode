package year2021

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/scylladb/go-set/strset"
)

type Digit struct {
	Segments *strset.Set
}

func (d *Digit) Len() int {
	return d.Segments.Size()
}

func (d *Digit) Intersects(other *Digit) bool {
	return d.IntersectSize(other) > 0
}

func (d *Digit) IntersectSize(other *Digit) int {
	return strset.Intersection(d.Segments, other.Segments).Size()
}

func (d *Digit) IsEqual(other *Digit) bool {
	return d.Segments.IsEqual(other.Segments)
}

func (d *Digit) String() string {
	return strings.Join(d.Segments.List(), "")
}

func NewDigit(s string) *Digit {
	return &Digit{Segments: strset.New(strings.Split(s, "")...)}
}

type DisplayNote struct {
	Observations []*Digit
	Digits       []*Digit
}

type DayEightInput struct {
	Notes []*DisplayNote
}

type DayEightOutput struct {
	PartOneAnswer int
	PartTwoAnswer int
}

func parseDayEight(rawInput string) (*DayEightInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayEightInput{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		note := &DisplayNote{}

		for i, f := range fields {
			if i < 10 {
				note.Observations = append(note.Observations, NewDigit(f))
			}
			if i > 10 {
				note.Digits = append(note.Digits, NewDigit(f))
			}
		}
		in.Notes = append(in.Notes, note)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

// num segments -> possible digits
// lengths
// 0: 6
// 1: 2
// 2: 5
// 3: 5
// 4: 4
// 5: 5
// 6: 6
// 7: 3
// 8: 7
// 9: 6
var lenToDigits map[int][]int = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

func solveDayEight(in *DayEightInput) (*DayEightOutput, error) {
	out := &DayEightOutput{}

	numUnique := 0
	digitsTotal := 0

	for _, n := range in.Notes {
		// Part 1 just count uniques
		for _, d := range n.Digits {
			candidates, ok := lenToDigits[d.Len()]
			if ok && len(candidates) == 1 {
				numUnique++
			}
		}

		// Part 2 is deduction time
		knownDigits := map[int]*Digit{}
		// first we collect known from uniques
		for _, o := range n.Observations {
			candidates, ok := lenToDigits[o.Len()]
			if ok && len(candidates) == 1 {
				knownDigits[candidates[0]] = o
			}
		}

		for _, o := range n.Observations {
			log.Printf("checking %s", o)
			switch o.Len() {
			case 2, 3, 4, 7:
				// already got'em
				log.Printf("unique")
			case 5:
				log.Printf("len 5")
				if o.IntersectSize(knownDigits[4]) == 2 {
					// 2 has length 5 and its intersection with 4 is 2 segments
					// (intersection with 4 is length 3 for 3 and 5)
					knownDigits[2] = o
					log.Printf("got 2")
				} else if o.IntersectSize(knownDigits[1]) == 2 {
					// 3 has length 5 and its intersection with 1 is 2 segments
					// (only 1 for 2 and 5)
					knownDigits[3] = o
					log.Printf("got 3")
				} else {
					// last len 5 digit is 5
					knownDigits[5] = o
					log.Printf("got 5")
				}
			case 6:
				log.Printf("len 6")
				if o.IntersectSize(knownDigits[4]) == 4 {
					// 9 has 4 intersections with 4, 0 and 6 only have three
					knownDigits[9] = o
					log.Printf("got 9")
				} else if o.IntersectSize(knownDigits[1]) == 2 {
					// 0 has 2 intersections with 1, 6 only has 1
					knownDigits[0] = o
					log.Printf("got 0")
				} else {
					// last len 6 digit is 6
					knownDigits[6] = o
					log.Printf("got 6")
				}

			default:
				panic(fmt.Sprintf("unexpected len! %s", o))
			}
		}

		// now we have a key and we can sum the digits
		value := 0
		tensPlace := 3
		for _, d := range n.Digits {
			found := false
			for digitValue, kd := range knownDigits {
				if kd.IsEqual(d) {
					found = true
					log.Printf("found %d for %s", digitValue, d)
					value += digitValue * int(math.Pow10(tensPlace))
				}
			}
			if !found {
				panic(fmt.Sprintf("no digit found for %s!", d))
			}
			tensPlace--
		}

		digitsTotal += value

		log.Printf("knownDigits for this line: %v", knownDigits)
		log.Printf("value for this line: %v", value)
	}

	out.PartOneAnswer = numUnique
	out.PartTwoAnswer = digitsTotal

	// intersect (8, 1) => right hand row, but uncler which of the two is
	// sub (7, 1) => top row

	return out, nil
}

func DayEight(rawInput string) (*DayEightOutput, error) {
	in, err := parseDayEight(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayEight(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
