package year2021

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type boardSpace struct {
	X      int
	Y      int
	Number int
	Marked bool
}

type bingoBoard struct {
	Spaces [][]*boardSpace
}

// returns justWon if marking this number made it a winner
// does not mark if it already won
func (b *bingoBoard) MarkNumber(n int) bool {
	if b.IsWinner() {
		return false
	}
	for _, row := range b.Spaces {
		for _, space := range row {
			if space.Number == n {
				space.Marked = true
			}
		}
	}
	return b.IsWinner()
}

func (b *bingoBoard) IsWinner() bool {
	// winning row?
	rowHasUnmarked := make([]bool, len(b.Spaces))
	colHasUnmarked := make([]bool, len(b.Spaces[0]))
	for y, row := range b.Spaces {
		for x, space := range row {
			if !space.Marked {
				rowHasUnmarked[y] = true
				colHasUnmarked[x] = true
			}
		}
	}

	for _, r := range rowHasUnmarked {
		if r == false {
			return true
		}
	}
	for _, c := range colHasUnmarked {
		if c == false {
			return true
		}
	}

	return false
}

func (b *bingoBoard) SumUnmarked() int {
	i := 0

	for _, row := range b.Spaces {
		for _, space := range row {
			if !space.Marked {
				i += space.Number
			}
		}
	}

	return i
}

func (b *bingoBoard) String() string {
	st := strings.Builder{}
	st.WriteString("[board]\n")

	for _, row := range b.Spaces {
		for _, space := range row {
			if space.Marked {
				st.WriteString("*")
			} else {
				st.WriteString(" ")
			}
			st.WriteString(fmt.Sprintf("%d  ", space.Number))
		}
		st.WriteString("\n")
	}
	st.WriteString("\n")

	return st.String()
}

type DayFourInput struct {
	DrawNumbers []int
	Boards      []*bingoBoard
}

type DayFourOutput struct {
	PartOneAnswer int

	PartTwoAnswer int
}

func parseDayFour(rawInput string) (*DayFourInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayFourInput{}

	var curBoard *bingoBoard

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("line: %s", line)

		// first we see a single comma separated list of draw numbers
		if strings.Contains(line, ",") {
			draws := strings.Split(line, ",")
			for _, d := range draws {
				dn, err := strconv.Atoi(d)
				if err != nil {
					return nil, err
				}
				in.DrawNumbers = append(in.DrawNumbers, dn)
			}
			log.Printf("got draw numbrs: %#v", in.DrawNumbers)
		}

		// then we see blank line separated boards
		if line == "" {
			if curBoard != nil {
				log.Printf("got board: %s", curBoard)
				in.Boards = append(in.Boards, curBoard)
			}
			curBoard = &bingoBoard{}
		}

		if strings.Contains(line, " ") {
			spaces := strings.Fields(line)
			row := make([]*boardSpace, len(spaces))
			y := len(curBoard.Spaces)
			for x, s := range spaces {
				sn, err := strconv.Atoi(s)
				if err != nil {
					return nil, err
				}
				row[x] = &boardSpace{X: x, Y: y, Number: sn}
			}
			curBoard.Spaces = append(curBoard.Spaces, row)
		}

	}
	log.Printf("last board: %s", curBoard)

	// append last board
	in.Boards = append(in.Boards, curBoard)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayFour(in *DayFourInput) (*DayFourOutput, error) {
	out := &DayFourOutput{}
	winningBoardScores := []int{}

	for _, dn := range in.DrawNumbers {
		for _, b := range in.Boards {
			justWon := b.MarkNumber(dn)

			if justWon {
				score := b.SumUnmarked() * dn
				if out.PartOneAnswer == 0 {
					out.PartOneAnswer = score
				}
				winningBoardScores = append(winningBoardScores, score)
				log.Printf("got winner with score: %d; %d winners", score, len(winningBoardScores))
			}
		}

		// done when all boards have won
		if len(winningBoardScores) == len(in.Boards) {
			log.Printf("all boards have won... breaking")
			break
		}
	}

	out.PartTwoAnswer = winningBoardScores[len(winningBoardScores)-1]

	return out, nil
}

func DayFour(rawInput string) (*DayFourOutput, error) {
	in, err := parseDayFour(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayFour(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
