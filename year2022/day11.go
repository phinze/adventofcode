package year2022

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"tailscale.com/util/must"
)

type DayElevenInput struct {
	Monkeys []*Monkey
}

type DayElevenOutput struct {
	MonkeyBusiness int
}

type MonkeyStuff struct {
	Worry int
}

type Monkey struct {
	Stuff         []*MonkeyStuff
	WorryOperator string
	WorryOperand  string
	TestDivBy     int
	TrueThrow     int
	FalseThrow    int
	InspectCount  int
}

func (m *Monkey) String() string {
	var b strings.Builder
	for _, s := range m.Stuff {
		b.WriteString(fmt.Sprintf("%d, ", s.Worry))
	}
	return b.String()
}

func (m *Monkey) Catch(s *MonkeyStuff) {
	m.Stuff = append(m.Stuff, s)
}

func parseDayEleven(rawInput string) (*DayElevenInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayElevenInput{}

	var thisMonkey *Monkey
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			continue
		}
		switch tokens[0] {
		case "Monkey":
			if thisMonkey != nil {
				in.Monkeys = append(in.Monkeys, thisMonkey)
			}
			thisMonkey = &Monkey{}
		case "Starting":
			for i := 2; i < len(tokens); i++ {
				thisMonkey.Stuff = append(thisMonkey.Stuff, &MonkeyStuff{
					Worry: must.Get(strconv.Atoi(
						strings.ReplaceAll(tokens[i], ",", ""),
					)),
				})
			}
		case "Operation:":
			thisMonkey.WorryOperator = tokens[4]
			thisMonkey.WorryOperand = tokens[5]
		case "Test:":
			thisMonkey.TestDivBy = must.Get(strconv.Atoi(tokens[3]))
		case "If":
			switch tokens[1] {
			case "true:":
				thisMonkey.TrueThrow = must.Get(strconv.Atoi(tokens[5]))
			case "false:":
				thisMonkey.FalseThrow = must.Get(strconv.Atoi(tokens[5]))
			}
		}
	}
	in.Monkeys = append(in.Monkeys, thisMonkey)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayEleven(in *DayElevenInput) (*DayElevenOutput, error) {
	out := &DayElevenOutput{}

	// part one
	for i := 0; i < 20; i++ {
		for _, m := range in.Monkeys {
			// log.Printf("Doing Monkey %d", im)
			for len(m.Stuff) > 0 {
				var s *MonkeyStuff
				s, m.Stuff = m.Stuff[0], m.Stuff[1:]
				// log.Printf("  inspecting item %d", s.Worry)
				m.InspectCount++
				switch m.WorryOperator {
				case "*":
					switch m.WorryOperand {
					case "old":
						s.Worry *= s.Worry
					default:
						s.Worry *= must.Get(strconv.Atoi(m.WorryOperand))
					}
				case "+":
					switch m.WorryOperand {
					case "old":
						s.Worry += s.Worry
					default:
						s.Worry += must.Get(strconv.Atoi(m.WorryOperand))
					}
				}
				// log.Printf("    after operation %s %s: %d", m.WorryOperator, m.WorryOperand, s.Worry)
				s.Worry /= 3
				// log.Printf("    after decat %d", s.Worry)
				if s.Worry%m.TestDivBy == 0 {
					in.Monkeys[m.TrueThrow].Catch(s)
				} else {
					in.Monkeys[m.FalseThrow].Catch(s)
				}
			}
		}
		log.Printf("After round %d", i)
		for i, m := range in.Monkeys {
			log.Printf("  Monkey %d: %s", i, m)
		}
	}

	sort.Slice(in.Monkeys, func(i, j int) bool {
		return in.Monkeys[j].InspectCount < in.Monkeys[i].InspectCount
	})

	out.MonkeyBusiness = in.Monkeys[0].InspectCount * in.Monkeys[1].InspectCount

	// part two

	return out, nil
}

func DayEleven(rawInput string) (*DayElevenOutput, error) {
	in, err := parseDayEleven(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayEleven(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
