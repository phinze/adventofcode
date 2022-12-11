package year2022

import (
	"bufio"
	"strconv"
	"strings"

	"tailscale.com/util/must"
)

type DayTenInput struct {
	Program []Instruction
}

type DayTenOutput struct {
	SignalStrengthSum int
	ScreenOutput      string
}

type Instruction interface {
	Execute(*CPU) (done bool)
}

type Noop struct{}

func (n *Noop) Execute(c *CPU) bool {
	return true
}

type AddX struct {
	Argument      int
	didMyThinking bool
}

func (a *AddX) Execute(c *CPU) bool {
	if a.didMyThinking {
		c.Register += a.Argument
		return true
	} else {
		a.didMyThinking = true
		return false
	}
}

func ParseInstruction(tokens []string) Instruction {
	switch tokens[0] {
	case "noop":
		return &Noop{}
	case "addx":
		return &AddX{
			Argument: must.Get(strconv.Atoi(tokens[1])),
		}
	}
	return nil
}

type CPU struct {
	Cycle    int
	Register int
	Program  []Instruction
	Screen   string

	crtCol int
	crtRow int
	pc     int
}

func NewCPU(program []Instruction) *CPU {
	return &CPU{
		Cycle:    1,
		Register: 1,
		Program:  program,
	}
}

func (c *CPU) DrawCRTPixel() {
	if c.Register-1 == c.crtCol || c.Register == c.crtCol || c.Register+1 == c.crtCol {
		c.Screen += "#"
	} else {
		c.Screen += "."
	}
	c.crtCol++
	if c.crtCol >= 40 {
		c.Screen += "\n"
		c.crtCol %= 40
		c.crtRow++
	}
}

func (c *CPU) Tick() {
	c.DrawCRTPixel()
	cur := c.Program[c.pc]
	done := cur.Execute(c)
	if done {
		c.pc++
	}
	c.Cycle++
}

func (c *CPU) SignalStrength() int {
	return c.Cycle * c.Register
}

func parseDayTen(rawInput string) (*DayTenInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayTenInput{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		in.Program = append(in.Program, ParseInstruction(tokens))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayTen(in *DayTenInput) (*DayTenOutput, error) {
	out := &DayTenOutput{}

	// part one
	cpu := NewCPU(in.Program)
	for i := 0; i < 240; i++ {
		cpu.Tick()
		if (cpu.Cycle-20)%40 == 0 {
			out.SignalStrengthSum += cpu.SignalStrength()
		}
	}

	// part two
	out.ScreenOutput = cpu.Screen

	return out, nil
}

func DayTen(rawInput string) (*DayTenOutput, error) {
	in, err := parseDayTen(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayTen(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
