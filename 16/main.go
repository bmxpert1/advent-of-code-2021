package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/davecgh/go-spew/spew"
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

type Packet struct {
	bits       string
	subPackets []*Packet
	proc       *Processor
}

func NewPacket(bits string, proc *Processor) *Packet {
	fmt.Println("new packet with", bits)
	return &Packet{
		bits:       bits,
		subPackets: []*Packet{},
		proc:       proc,
	}
}

func (p *Packet) version() int {
	v, _ := strconv.ParseInt(p.bits[0:3], 2, 64)
	return int(v)
}
func (p *Packet) typeId() int {
	v, _ := strconv.ParseInt(p.bits[3:6], 2, 64)
	return int(v)
}
func (p *Packet) start() {
	switch p.typeId() {
	case 4: // literal value
		p.loadLiteralGroups()
	default: // operator packet
		p.loadOperatorPackets()
	}
}
func (p *Packet) loadLiteralGroups() {
	loaded := false
	pos := 6

	for !loaded {
		p.bits += p.proc.nextGroup()

		if len(p.bits) > pos+5 {
			grp := p.bits[pos : pos+5]

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
			if i+5 < len(bits) {
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
			} else {
				p.bits += p.proc.nextGroup()
			}
		}

		loaded = false
		var subPacketBits string
		minPos := pos + int(subPacketLength)

		for !loaded {
			if len(p.bits) > minPos {
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

		spew.Dump(p.subPackets)
	case '1':
		// then the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
	}
}

type Processor struct {
	bits    string
	pos     int
	packets []*Packet
}

func NewProcessor(bits string) *Processor {
	fmt.Println("new processor with", bits)
	return &Processor{
		bits:    bits,
		pos:     0,
		packets: []*Packet{},
	}
}

func (p *Processor) process() {
	for p.pos+6 < len(p.bits) {
		// start with first 2 groups of 4 bits because we need at least 6 for version+type
		packet := NewPacket(p.nextGroupOf(6), p)
		p.packets = append(p.packets, packet)
		packet.start()
		fmt.Println(packet.literalValue())
	}
}
func (p *Processor) nextGroupOf(size int) string {
	grp := p.bits[p.pos : p.pos+size]
	p.pos += size
	return grp
}
func (p *Processor) nextGroup() string {
	return p.nextGroupOf(defaultGroupSize)
}

func main() {
	input, _ := ioutil.ReadFile("example_input3.txt")
	bits := ""
	for _, hex := range input {
		bits += hexBinMap[rune(hex)]
	}

	proc = NewProcessor(bits)
	proc.process()

	fmt.Println(proc.packets[0].literalValue())

	// fmt.Println(proc.currentPacket.version(), proc.currentPacket.typeId(), proc.currentPacket.bits)
}
