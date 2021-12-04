package year2021

import (
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
)

type DayThreeInput struct {
	Diagnostics []uint64
	BitLength   int
}

type DayThreeOutput struct {
	GammaRate   uint64
	EpsilonRate uint64

	PartOneAnswer int

	OxygenRating   uint64
	ScrubberRating uint64

	PartTwoAnswer int
}

func parseDayThree(rawInput string) (*DayThreeInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))

	in := &DayThreeInput{}

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.ParseUint(line, 2, len(line))
		if err != nil {
			return nil, err
		}
		in.BitLength = len(line)
		in.Diagnostics = append(in.Diagnostics, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func isBitOn(num uint64, bitPlace int) bool {
	mask := uint64(math.Pow(2, float64(bitPlace)))
	return num&mask > 0
}

func numBitsInPosition(nums []uint64, bitPlace int) int {
	total := 0
	for _, n := range nums {
		if isBitOn(n, bitPlace) {
			total++
		}
	}
	return total
}

func majorityBitInPosition(nums []uint64, bitPlace int) bool {
	count := numBitsInPosition(nums, bitPlace)
	other := len(nums) - count

	return count >= other
}

func minorityBitInPosition(nums []uint64, bitPlace int) bool {
	count := numBitsInPosition(nums, bitPlace)
	other := len(nums) - count

	return count < other
}

func dumpBits(nums []uint64) {
	log.Printf("    [[")
	for _, n := range nums {
		log.Printf("     %012b", n)
	}
	log.Printf("    ]]")
}

func solveDayThree(in *DayThreeInput) (*DayThreeOutput, error) {
	out := &DayThreeOutput{}

	// part 1
	bits := make([]int, in.BitLength)
	for i := range bits {
		bits[i] = numBitsInPosition(in.Diagnostics, i)
	}
	for i, bitsSeen := range bits {
		if bitsSeen > len(in.Diagnostics)/2 {
			out.GammaRate += uint64(math.Pow(2, float64(i)))
		}
	}
	out.EpsilonRate = ^out.GammaRate & (uint64(math.Pow(2, float64(in.BitLength))) - 1)
	out.PartOneAnswer = int(out.GammaRate * out.EpsilonRate)

	// part 2
	candidateOxygenRatings := make([]uint64, len(in.Diagnostics))
	candidateCO2Ratings := make([]uint64, len(in.Diagnostics))
	copy(candidateOxygenRatings, in.Diagnostics)
	copy(candidateCO2Ratings, in.Diagnostics)

	for bitPlace := in.BitLength - 1; bitPlace >= 0; bitPlace-- {
		bitValueToKeep := majorityBitInPosition(candidateOxygenRatings, bitPlace)
		for i := len(candidateOxygenRatings) - 1; i >= 0; i-- {
			d := candidateOxygenRatings[i]
			bitValue := isBitOn(d, bitPlace)
			if bitValue != bitValueToKeep {
				candidateOxygenRatings = append(candidateOxygenRatings[:i], candidateOxygenRatings[i+1:]...)
			}
		}
		if len(candidateOxygenRatings) == 1 && out.OxygenRating == 0 {
			out.OxygenRating = candidateOxygenRatings[0]
		}
	}

	for bitPlace := in.BitLength - 1; bitPlace >= 0; bitPlace-- {
		bitValueToKeep := minorityBitInPosition(candidateCO2Ratings, bitPlace)
		for i := len(candidateCO2Ratings) - 1; i >= 0; i-- {
			d := candidateCO2Ratings[i]
			bitValue := isBitOn(d, bitPlace)
			if bitValue != bitValueToKeep {
				candidateCO2Ratings = append(candidateCO2Ratings[:i], candidateCO2Ratings[i+1:]...)
			}
		}
		if len(candidateCO2Ratings) == 1 && out.ScrubberRating == 0 {
			out.ScrubberRating = candidateCO2Ratings[0]
		}
	}

	out.PartTwoAnswer = int(out.OxygenRating * out.ScrubberRating)

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
