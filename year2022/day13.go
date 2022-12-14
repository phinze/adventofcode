package year2022

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"tailscale.com/util/must"
)

type DayThirteenInput struct {
	Packets []*Packet
}
type DayThirteenOutput struct {
	PartOneAnswer      int
	PartOneComparisons []bool
	PartTwoAnswer      int
}

// first return is answer, second return is done
func inOrder(left, right PacketItem) (bool, bool) {
	log.Printf("  [inOrder] left : %s", left)
	log.Printf("  [inOrder] right: %s", right)
	switch l := left.(type) {
	case int:
		switch r := right.(type) {
		case int:
			log.Printf("    checking ints %d <= %d", l, r)
			if l < r {
				log.Printf("     left less than right, in order!")
				return true, true
			} else if l > r {
				log.Printf("     left greater than right, not in order!")
				return false, true
			} else {
				log.Printf("     left equal to right, keep checking...")
				return true, false
			}
		case *PacketList:
			log.Printf("    mixed types; left is %d and right is %s, wrapping left in list", l, r)
			return inOrder(NewPacketList(l), r)
		}
	case *PacketList:
		switch r := right.(type) {
		case int:
			log.Printf("    mixed types; left is %s and right is %d, wrapping right in list", l, r)
			return inOrder(l, NewPacketList(r))
		case *PacketList:
			ln := len(l.Items)
			rn := len(r.Items)
			if ln == 0 && rn == 0 {
				log.Printf("    two empty lists, keep checking")
				return true, false
			} else if ln == 0 && rn > 0 {
				log.Printf("    left ran out before right, so in order!")
				return true, true
			} else if ln > 0 && rn == 0 {
				log.Printf("    right ran out before left, not in order!")
				return false, true
			} else if ln > 0 && rn > 0 {
				// both lists have items, compare first elements and recurse
				newL := NewPacketList(l.Items[1:]...)
				newR := NewPacketList(r.Items[1:]...)
				log.Printf("    have two nonempty lists, comparing first element")
				firstComparison, done := inOrder(l.Items[0], r.Items[0])
				if done {
					return firstComparison, done
				}
				log.Printf("first element not conclusive, comparing rest")
				return inOrder(newL, newR)
			}
		}
	}
	panic("unknown types")
}

func InOrder(left, right PacketItem) bool {
	answer, done := inOrder(left, right)
	if !done {
		panic("not done how")
	}
	return answer
}

type PacketItem interface{}

type PacketNumber int

type PacketList struct {
	Items  []PacketItem
	Parent *PacketList
}

func (pl *PacketList) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i, item := range pl.Items {
		switch v := item.(type) {
		case *PacketList:
			b.WriteString(v.String())
		case int:
			b.WriteString(fmt.Sprintf("%d", v))
		}
		if i != len(pl.Items)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("]")
	return b.String()
}

func NewPacketList(items ...PacketItem) *PacketList {
	return &PacketList{Items: items}
}

type Packet struct {
	Data     *PacketList
	OrigLine string
}

func ParsePacket(line string) *Packet {
	packet := &Packet{
		Data:     &PacketList{},
		OrigLine: line,
	}
	var cur *PacketList
	for i := 0; i < len(line); {
		switch line[i] {
		case '[':
			newList := &PacketList{Parent: cur}
			if cur == nil {
				packet.Data = newList
			} else {
				cur.Items = append(cur.Items, newList)
			}
			cur = newList
			i++
		case ']':
			cur = cur.Parent
			i++
		case ',':
			// we don't care about commas
			i++
		default:
			// we have a number, it might be more than one char so we need
			// to continue until there's a non-digit
			numStr := ""
			for line[i] >= '0' && line[i] <= '9' {
				numStr = numStr + string(line[i])
				i++
			}
			num := must.Get(strconv.Atoi(numStr))
			cur.Items = append(cur.Items, num)
		}
	}
	return packet
}

func (p *Packet) String() string {
	return p.Data.String()
}

func parseDayThirteen(rawInput string) (*DayThirteenInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayThirteenInput{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		in.Packets = append(in.Packets, ParsePacket(line))
	}
	log.Printf("got %d packets", len(in.Packets))

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayThirteen(in *DayThirteenInput) (*DayThirteenOutput, error) {
	out := &DayThirteenOutput{}

	// part one
	pairIndex := 1
	for i := 0; i < len(in.Packets)-1; i += 2 {
		this := in.Packets[i]
		next := in.Packets[i+1]

		log.Printf("checking pair %d (%d and %d)", pairIndex, i, i+1)
		log.Printf("  left : %s", spew.Sdump(this))
		log.Printf("  origL: %s", this.OrigLine)
		log.Printf("  right: %s", spew.Sdump(next))
		log.Printf("  origR: %s", next.OrigLine)

		result := InOrder(this.Data, next.Data)
		out.PartOneComparisons = append(out.PartOneComparisons, result)
		if result {
			out.PartOneAnswer += pairIndex
			log.Printf("  in order!")
		} else {
			log.Printf("  not in order!")
		}
		pairIndex++
	}

	// part two
	in.Packets = append(in.Packets, ParsePacket("[[2]]"), ParsePacket("[[6]]"))

	sort.Slice(in.Packets, func(i, j int) bool {
		return InOrder(in.Packets[i].Data, in.Packets[j].Data)
	})

	out.PartTwoAnswer = 1
	for i := 0; i < len(in.Packets); i++ {
		p := in.Packets[i].Data.String()
		// log.Printf("packet %d - %s", i+1, p)
		if p == "[[6]]" || p == "[[2]]" {
			log.Printf("divider at index %d", i+1)
			out.PartTwoAnswer *= i + 1
		}
	}

	return out, nil
}

func DayThirteen(rawInput string) (*DayThirteenOutput, error) {
	in, err := parseDayThirteen(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayThirteen(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
