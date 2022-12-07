package year2022

import (
	"strings"
)

type DaySixInput struct {
	DataStream string
}

type DaySixOutput struct {
	PacketPosition  int
	MessagePosition int
}

func parseDaySix(rawInput string) (*DaySixInput, error) {
	in := &DaySixInput{
		DataStream: rawInput,
	}

	return in, nil
}

func foundMarker(window []byte) bool {
	valCounts := map[byte]int{}
	for _, b := range window {
		// bail if window is not full
		if b == 0 {
			return false
		}
		valCounts[b]++
	}
	if len(valCounts) < len(window) {
		// bail if not enough uniques seen
		return false
	}
	// verify all uniques seen only once
	for _, count := range valCounts {
		if count > 1 {
			return false
		}
	}
	return true
}

func solveDaySix(in *DaySixInput) (*DaySixOutput, error) {
	out := &DaySixOutput{}

	// part one
	reader := strings.NewReader(in.DataStream)
	window := make([]byte, 4)

	for i := 0; ; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		window[i%len(window)] = b
		if foundMarker(window) {
			out.PacketPosition = i + 1
			break
		}
	}

	// part two
	reader = strings.NewReader(in.DataStream)
	window = make([]byte, 14)

	for i := 0; ; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		window[i%len(window)] = b
		if foundMarker(window) {
			out.MessagePosition = i + 1
			break
		}
	}

	return out, nil
}

func DaySix(rawInput string) (*DaySixOutput, error) {
	in, err := parseDaySix(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDaySix(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
