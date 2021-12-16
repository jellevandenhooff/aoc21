package main

import (
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

type State struct {
	bits []int
	idx  int
}

func (s *State) Read(bits int) int {
	v := 0
	for i := 0; i < bits; i++ {
		log.Println(s.bits[s.idx])
		v = (v << 1) | s.bits[s.idx]
		s.idx++
	}
	return v
}

func (s *State) ReadMany(bits int) []int {
	many := s.bits[s.idx : s.idx+bits]
	s.idx += bits
	return many
}

func Parse(s *State, d string) int {
	version := s.Read(3)
	log.Println(d, "version", version)
	typ := s.Read(3)
	log.Println(d, "typ", typ)

	ans := 0
	ans += version

	switch typ {
	// literal
	case 4:
		log.Println(d, "literal")
		for {
			c := s.Read(5)
			log.Println(d, "part", c)
			if c&(1<<4) == 0 {
				log.Println(d, "bailing")
				break
			}
		}

	default:
		log.Println(d, "operator")

		lenType := s.Read(1)
		log.Println(d, "lentype", lenType)
		switch lenType {
		case 0:
			numBits := s.Read(15)
			log.Println(d, "numbits", numBits)
			subBits := s.ReadMany(numBits)

			subS := &State{
				bits: subBits,
			}
			for subS.idx < len(subS.bits) {
				ans += Parse(subS, d+"  ")
			}
		case 1:
			numPackets := s.Read(11)
			log.Println(d, "numpackets", numPackets)
			for i := 0; i < numPackets; i++ {
				ans += Parse(s, d+"  ")
			}
		}
	}

	return ans
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	bytes, err := hex.DecodeString(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	bits := make([]int, 0) // len(bytes)*8)
	for _, b := range bytes {
		for i := 0; i < 8; i++ {
			bits = append(bits, (int(b)>>(7-i))&1)
		}
	}

	s := &State{
		bits: bits,
	}

	log.Println(Parse(s, ""))
}
