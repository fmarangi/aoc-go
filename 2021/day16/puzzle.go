package day16

import (
	"fmt"
	"math"
	"strconv"
)

const (
	TYPE_SUM = iota
	TYPE_PRODUCT
	TYPE_MINIMUM
	TYPE_MAXIMUM
	TYPE_LITERAL
	TYPE_GREATER_THAN
	TYPE_LESS_THAN
	TYPE_EQUAL_TO
)

type Packet struct {
	version, typ, value int
	subs                []Packet
}

func ToBin(hex string) (result string) {
	for _, c := range hex {
		val, _ := strconv.ParseInt(string(c), 16, 32)
		result += fmt.Sprintf("%04b", val)
	}
	return
}

func toInt(bin string) int {
	num, _ := strconv.ParseInt(bin, 2, 64)
	return int(num)
}

func readPacket(bits string) (Packet, string) {
	packet := Packet{version: toInt(bits[0:3]), typ: toInt(bits[3:6])}

	switch {
	case packet.typ == TYPE_LITERAL:
		v, i := "", 6
		for ; ; i += 5 {
			v += bits[i+1 : i+5]
			if bits[i] == '0' {
				break
			}
		}
		packet.value = toInt(v)
		bits = bits[i+5:]

	case bits[6] == '0':
		length := toInt(bits[7:22]) // 15 positions
		subs := bits[22 : 22+length]
		for len(subs) > 6 {
			sub, rest := readPacket(subs)
			packet.subs = append(packet.subs, sub)
			subs = rest
		}
		bits = bits[22+length:]

	default:
		length := toInt(bits[7:18]) // 11 positions
		subs := bits[18:]
		for i := 0; i < length; i++ {
			sub, rest := readPacket(subs)
			packet.subs = append(packet.subs, sub)
			subs = rest
		}
		bits = subs
	}

	return packet, bits
}

func (p Packet) Version() (v int) {
	for _, s := range p.subs {
		v += s.Version()
	}
	return v + p.version
}

func (p Packet) Value() (v int) {
	switch p.typ {
	case TYPE_SUM:
		for _, s := range p.subs {
			v += s.Value()
		}
	case TYPE_PRODUCT:
		v = 1
		for _, s := range p.subs {
			v *= s.Value()
		}
	case TYPE_MINIMUM:
		v = math.MaxInt64
		for _, s := range p.subs {
			if val := s.Value(); val < v {
				v = val
			}
		}
	case TYPE_MAXIMUM:
		v = 0
		for _, s := range p.subs {
			if val := s.Value(); val > v {
				v = val
			}
		}
	case TYPE_LITERAL:
		v = p.value
	case TYPE_GREATER_THAN:
		if p.subs[0].Value() > p.subs[1].Value() {
			v = 1
		}
	case TYPE_LESS_THAN:
		if p.subs[0].Value() < p.subs[1].Value() {
			v = 1
		}
	case TYPE_EQUAL_TO:
		if p.subs[0].Value() == p.subs[1].Value() {
			v = 1
		}
	}
	return
}

func Part1(input string) int {
	p, _ := readPacket(ToBin(input))
	return p.Version()
}

func Part2(input string) int {
	p, _ := readPacket(ToBin(input))
	return p.Value()
}
