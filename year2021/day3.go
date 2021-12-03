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

func solveDayThree(in *DayThreeInput) (*DayThreeOutput, error) {
	out := &DayThreeOutput{}

	bits := make([]int, in.BitLength)
	for _, d := range in.Diagnostics {
		log.Printf("checking: %012b", d)
		for i := range bits {
			pow := uint64(math.Pow(2, float64(i)))
			mod := d & pow
			log.Printf("d: %d, pow: %d, mod: %d", d, pow, mod)
			log.Printf("  bit %d is...", i)
			if mod > 0 {
				log.Printf("on!")
				bits[i]++
			} else {
				log.Printf("off.")
			}
		}
	}
	log.Printf("%#v", bits)
	for i, bitsSeen := range bits {
		if bitsSeen > len(in.Diagnostics)/2 {
			out.GammaRate += uint64(math.Pow(2, float64(i)))
		}
	}

	log.Printf("gamma %b", out.GammaRate)
	log.Printf("not gamma %b", (^out.GammaRate & (uint64(math.Pow(2, float64(in.BitLength))) - 1)))
	out.EpsilonRate = ^out.GammaRate & (uint64(math.Pow(2, float64(in.BitLength))) - 1)
	out.PartOneAnswer = int(out.GammaRate * out.EpsilonRate)

	candidateOxygenRatings := make([]uint64, len(in.Diagnostics))
	candidateCO2Ratings := make([]uint64, len(in.Diagnostics))
	copy(candidateOxygenRatings, in.Diagnostics)
	copy(candidateCO2Ratings, in.Diagnostics)

	for bitPlace := len(bits) - 1; bitPlace >= 0; bitPlace-- {
		log.Printf("searching for ratings bit place: %d, freqency: %d...", bitPlace, bits[bitPlace])
		oxygenBitValueToKeep := (bits[bitPlace] >= len(in.Diagnostics)/2)
		co2BitValueToKeep := (bits[bitPlace] < len(in.Diagnostics)/2)
		for i := len(candidateOxygenRatings) - 1; i >= 0; i-- {
			d := candidateOxygenRatings[i]
			pow := uint64(math.Pow(2, float64(bitPlace)))
			bitIsOn := (d & pow) > 0
			if bitIsOn != oxygenBitValueToKeep {
				candidateOxygenRatings = append(candidateOxygenRatings[:i], candidateOxygenRatings[i+1:]...)
			}
		}
		for i := len(candidateCO2Ratings) - 1; i >= 0; i-- {
			d := candidateCO2Ratings[i]
			pow := uint64(math.Pow(2, float64(bitPlace)))
			bitIsOn := (d & pow) > 0
			if bitIsOn != co2BitValueToKeep {
				candidateCO2Ratings = append(candidateCO2Ratings[:i], candidateCO2Ratings[i+1:]...)
			}
		}
		log.Printf("for bit %d, got %d oxygen ratings", bitPlace, len(candidateOxygenRatings))
		log.Printf("for bit %d, got %d co2 ratings", bitPlace, len(candidateCO2Ratings))
		if len(candidateCO2Ratings) == 1 && out.ScrubberRating == 0 {
			out.ScrubberRating = candidateCO2Ratings[0]
		}
		if len(candidateOxygenRatings) == 1 && out.OxygenRating == 0 {
			out.OxygenRating = candidateOxygenRatings[0]
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
