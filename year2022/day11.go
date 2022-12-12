package year2022

import (
	"bufio"
	"log"
	"math/big"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/mitchellh/copystructure"
	"tailscale.com/util/must"
)

type DayElevenInput struct {
	Monkeys []*Monkey
}

type DayElevenOutput struct {
	MonkeyBusiness     *big.Int
	WorryFreeMonkeyBiz *big.Int
}

type MonkeyStuff struct {
	Worry *big.Int
}

type Monkey struct {
	Stuff              []*MonkeyStuff
	WorryOperator      string
	WorryOperandIsSelf bool
	WorryOperand       *big.Int
	TestDivBy          *big.Int
	TrueThrow          int
	FalseThrow         int
	InspectCount       *big.Int
}

var One = big.NewInt(1)
var Three = big.NewInt(3)

func (m *Monkey) Catch(s *MonkeyStuff) {
	m.Stuff = append(m.Stuff, s)
}

func MonkeyPlaytime(monkeys []*Monkey, rounds int, decayFunc func(*big.Int)) *big.Int {
	monkeys = must.Get(copystructure.Copy(monkeys)).([]*Monkey)
	modTest := new(big.Int)
	modRem := new(big.Int)
	for i := 0; i < rounds; i++ {
		if i%100 == 0 {
			log.Printf("Doing round %d", i)
		}
		for _, m := range monkeys {
			// log.Printf("Doing Monkey %d", im)
			// log.Printf("  monkey %s", spew.Sdump(m))
			for len(m.Stuff) > 0 {
				var s *MonkeyStuff
				s, m.Stuff = m.Stuff[0], m.Stuff[1:]
				// log.Printf("  inspecting stuff %#v", s)
				// log.Printf("  inspecting item %#v", s.Worry)
				m.InspectCount.Add(m.InspectCount, One)
				switch m.WorryOperator {
				case "*":
					if m.WorryOperandIsSelf {
						s.Worry.Mul(s.Worry, s.Worry)
					} else {
						s.Worry.Mul(s.Worry, m.WorryOperand)
					}
				case "+":
					if m.WorryOperandIsSelf {
						s.Worry.Add(s.Worry, s.Worry)
					} else {
						s.Worry.Add(s.Worry, m.WorryOperand)
					}
				}
				// log.Printf("    after operation %s %s: %d", m.WorryOperator, m.WorryOperand, s.Worry)
				decayFunc(s.Worry)
				// log.Printf("    after decay %d", s.Worry)
				modTest.Set(s.Worry)
				modTest.QuoRem(modTest, m.TestDivBy, modRem)
				if len(modRem.Bits()) == 0 {
					// log.Printf("    divisible by %s, tossing to %d", m.TestDivBy, m.TrueThrow)
					monkeys[m.TrueThrow].Catch(s)
				} else {
					// log.Printf("    NOT divisible by %s, tossing to %d", m.TestDivBy, m.FalseThrow)
					monkeys[m.FalseThrow].Catch(s)
				}
			}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[j].InspectCount.Cmp(monkeys[i].InspectCount) == -1
	})
	return new(big.Int).Mul(monkeys[0].InspectCount, monkeys[1].InspectCount)
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
			thisMonkey = &Monkey{InspectCount: big.NewInt(0)}
		case "Starting":
			for i := 2; i < len(tokens); i++ {
				thisMonkey.Stuff = append(thisMonkey.Stuff, &MonkeyStuff{
					Worry: big.NewInt(int64(must.Get(strconv.Atoi(
						strings.ReplaceAll(tokens[i], ",", ""),
					)))),
				})
			}
		case "Operation:":
			thisMonkey.WorryOperator = tokens[4]
			if tokens[5] == "old" {
				thisMonkey.WorryOperandIsSelf = true
			} else {
				thisMonkey.WorryOperand = big.NewInt(int64(must.Get(strconv.Atoi(tokens[5]))))
			}
		case "Test:":
			thisMonkey.TestDivBy = big.NewInt(int64(must.Get(strconv.Atoi(tokens[3]))))
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

	copystructure.Copiers[reflect.TypeOf(big.Int{})] =
		func(raw interface{}) (interface{}, error) {
			in := raw.(big.Int)
			out := new(big.Int).Set(&in)
			return *out, nil
		}

	// part one
	three := big.NewInt(3)
	divByThree := func(i *big.Int) {
		i.Div(i, three)
	}
	out.MonkeyBusiness = MonkeyPlaytime(in.Monkeys, 20, divByThree)

	// part two
	// Decay by a number that will maintain modulo checks for all monkeys
	commonMultiple := big.NewInt(1)
	for _, m := range in.Monkeys {
		commonMultiple.Mul(commonMultiple, m.TestDivBy)
	}
	reduceByMultiple := func(i *big.Int) {
		i.Mod(i, commonMultiple)
	}
	out.WorryFreeMonkeyBiz = MonkeyPlaytime(in.Monkeys, 10000, reduceByMultiple)

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
