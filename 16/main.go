package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const (
	defaultGroupSize int = 1
)

var hexBinMap = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}
var proc *Processor
var allPackets = []*Packet{}

type Packet struct {
	bits       string
	subPackets []*Packet
	proc       *Processor
}

func NewPacket(bits string, proc *Processor) (packet *Packet) {
	// fmt.Println("new packet with", bits)

	packet = &Packet{
		bits:       bits,
		subPackets: []*Packet{},
		proc:       proc,
	}

	allPackets = append(allPackets, packet)

	return packet
}

func (p *Packet) version() int {
	v, _ := strconv.ParseInt(p.bits[0:3], 2, 64)
	return int(v)
}
func (p *Packet) typeId() int {
	v, _ := strconv.ParseInt(p.bits[3:6], 2, 64)
	return int(v)
}
func (p *Packet) process() {
	switch p.typeId() {
	case 4: // literal value
		p.loadLiteralGroups()
	default: // operator packet
		p.loadOperatorPackets()
	}
}
func (p *Packet) loadLiteralGroups() {
	// fmt.Println("literal packet")

	loaded := false
	pos := 6

	for !loaded {
		p.bits += p.proc.nextGroupOf(5)

		if len(p.bits) >= pos+5 {
			grp := p.bits[pos : pos+5]

			// fmt.Println("grp: ", grp)

			if grp[0] == '0' {
				loaded = true
			} else {
				pos += 5
			}
		}
	}
}
func (p *Packet) literalValue() int {
	if p.typeId() == 4 {
		bits := p.bits[6:len(p.bits)]
		binSum := ""
		for i := 0; i < len(bits); i += 5 {
			if i+5 <= len(bits) {
				grp := bits[i : i+5]
				binSum += grp[1:]
			}
		}
		dec, _ := strconv.ParseInt(binSum, 2, 64)
		return int(dec)
	} else {
		return -1
	}
}
func (p *Packet) loadOperatorPackets() {
	loaded := false
	for len(p.bits) < 7 {
		p.bits += p.proc.nextGroup()
	}
	lengthTypeId := p.bits[6]
	pos := 7

	switch lengthTypeId {
	case '0':
		var subPacketLength int64
		// then the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet
		for !loaded {
			if len(p.bits) > pos+15 {
				next15 := p.bits[pos : pos+15]
				subPacketLength, _ = strconv.ParseInt(next15, 2, 64)
				pos += 15
				loaded = true
				// fmt.Printf("operator packet 0: [%v]\n", subPacketLength)
			} else {
				p.bits += p.proc.nextGroup()
			}
		}

		loaded = false
		var subPacketBits string
		minPos := pos + int(subPacketLength)

		for !loaded {
			if len(p.bits) >= minPos {
				subPacketBits = p.bits[pos:minPos]
				loaded = true
			} else {
				p.bits += p.proc.nextGroup()
			}
		}

		// create a new processor for these bits
		tempProc := NewProcessor(subPacketBits)
		tempProc.process()
		p.subPackets = tempProc.packets
	case '1':
		var numSubPackets int64
		// then the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
		for !loaded {
			if len(p.bits) >= pos+11 {
				next11 := p.bits[pos : pos+11]
				numSubPackets, _ = strconv.ParseInt(next11, 2, 64)
				pos += 11
				loaded = true
				// fmt.Printf("operator packet 1: [%v]\n", numSubPackets)
			} else {
				p.bits += p.proc.nextGroup()
			}
		}

		startingPacketCount := len(allPackets)

		for len(allPackets) < startingPacketCount+int(numSubPackets) {
			tempPacket := NewPacket(p.proc.nextGroupOf(6), p.proc)
			// fmt.Println("o0", tempPacket.bits)
			tempPacket.process()
			p.subPackets = append(p.subPackets, tempPacket)
			// fmt.Printf("O1 %v [%v]\n", tempPacket.bits, tempPacket.literalValue())
		}
	}
}

type Processor struct {
	bits    string
	pos     int
	packets []*Packet
}

func NewProcessor(bits string) *Processor {
	// fmt.Println("new processor with", bits)
	return &Processor{
		bits:    bits,
		pos:     0,
		packets: []*Packet{},
	}
}

func (p *Processor) process() {
	for p.pos+6 < len(p.bits) {
		packet := NewPacket(p.nextGroupOf(6), p)
		p.packets = append(p.packets, packet)
		packet.process()
		// fmt.Printf("P %v [%v]\n", packet.bits, packet.literalValue())
	}
	// fmt.Println("end processor with", p.bits)
}
func (p *Processor) nextGroupOf(size int) string {
	grp := p.bits[p.pos:int(math.Min(float64(p.pos+size), float64(len(p.bits))))]
	p.pos += size
	return grp
}
func (p *Processor) nextGroup() string {
	return p.nextGroupOf(defaultGroupSize)
}

func main() {
	// TODO 5 and 6 arent working
	input, _ := ioutil.ReadFile("input.txt")
	bits := ""
	for _, hex := range input {
		bits += hexBinMap[rune(hex)]
	}

	proc = NewProcessor(bits)
	proc.process()

	versionSum := 0

	for _, packet := range allPackets {
		versionSum += packet.version()
	}

	fmt.Println(versionSum)
}
