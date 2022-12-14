package year2022

import (
	"bufio"
	"fmt"
	"log"
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
		packet := &Packet{
			Data:     &PacketList{},
			OrigLine: line,
		}
		var cur *PacketList
		for _, c := range line {
			switch c {
			case '[':
				newList := &PacketList{Parent: cur}
				if cur == nil {
					packet.Data = newList
				} else {
					cur.Items = append(cur.Items, newList)
				}
				cur = newList
			case ']':
				cur = cur.Parent
			case ',':
				// we don't care about commas
			default:
				num := must.Get(strconv.Atoi(string(c)))
				cur.Items = append(cur.Items, num)
			}
		}
		log.Printf("packet: %s", packet)
		in.Packets = append(in.Packets, packet)
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
		log.Printf("  right: %s", spew.Sdump(next))

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
