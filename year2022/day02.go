package year2022

import (
	"bufio"
	"strings"
)

var shapeScores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var resultScores = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

type RPSRound struct {
	OpponentPlay string
	MyPlay       string
}

var LoseScore = 0
var TieScore = 3
var WinScore = 6

var RockScore = 1
var PaperScore = 2
var ScissorsScore = 3

func (r *RPSRound) PartOneScore() int {
	return r.PartOneGameScore() + shapeScores[r.MyPlay]
}

func (r *RPSRound) PartOneGameScore() int {
	switch r.OpponentPlay {
	case "A":
		switch r.MyPlay {
		case "X":
			return TieScore
		case "Y":
			return WinScore
		case "Z":
			return LoseScore
		}
	case "B":
		switch r.MyPlay {
		case "X":
			return LoseScore
		case "Y":
			return TieScore
		case "Z":
			return WinScore
		}
	case "C":
		switch r.MyPlay {
		case "X":
			return WinScore
		case "Y":
			return LoseScore
		case "Z":
			return TieScore
		}
	default:
		panic("WTF")
	}
	return -1
}

func (r *RPSRound) PartTwoScore() int {
	return resultScores[r.MyPlay] + r.PartTwoPlayScore()
}

func (r *RPSRound) PartTwoPlayScore() int {
	switch r.OpponentPlay {
	case "A":
		switch r.MyPlay {
		case "X":
			return ScissorsScore
		case "Y":
			return RockScore
		case "Z":
			return PaperScore
		}
	case "B":
		switch r.MyPlay {
		case "X":
			return RockScore
		case "Y":
			return PaperScore
		case "Z":
			return ScissorsScore
		}
	case "C":
		switch r.MyPlay {
		case "X":
			return PaperScore
		case "Y":
			return ScissorsScore
		case "Z":
			return RockScore
		}
	}
	return -1
}

type DayTwoInput struct {
	StrategyGuide []*RPSRound
}
type DayTwoOutput struct {
	PartOneScore int
	PartTwoScore int
}

func parseDayTwo(rawInput string) (*DayTwoInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayTwoInput{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		in.StrategyGuide = append(in.StrategyGuide, &RPSRound{
			OpponentPlay: fields[0],
			MyPlay:       fields[1],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayTwo(in *DayTwoInput) (*DayTwoOutput, error) {
	out := &DayTwoOutput{}

	// part one
	for _, round := range in.StrategyGuide {
		out.PartOneScore += round.PartOneScore()
	}

	// part two
	for _, round := range in.StrategyGuide {
		out.PartTwoScore += round.PartTwoScore()
	}

	return out, nil
}

func DayTwo(rawInput string) (*DayTwoOutput, error) {
	in, err := parseDayTwo(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayTwo(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
