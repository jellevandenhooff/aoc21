package main

import (
	"encoding/hex"
	"io"
	"log"
	"math"
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

	switch typ {
	// literal
	case 4:
		log.Println(d, "literal")
		v := 0
		for {
			c := s.Read(5)
			log.Println(d, "part", c)
			v = (v << 4) | (c & ((1 << 4) - 1))
			if c&(1<<4) == 0 {
				log.Println(d, "bailing")
				break
			}
			log.Println(d, "v", v)
		}
		return v

	default:
		log.Println(d, "operator")
		var vs []int

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
				vs = append(vs, Parse(subS, d+"  "))
			}
		case 1:
			numPackets := s.Read(11)
			log.Println(d, "numpackets", numPackets)
			for i := 0; i < numPackets; i++ {
				vs = append(vs, Parse(s, d+"  "))
			}
		}

		switch typ {
		case 0:
			ans := 0
			for _, v := range vs {
				ans += v
			}
			return ans
		case 1:
			ans := 1
			for _, v := range vs {
				ans *= v
			}
			return ans
		case 2:
			ans := math.MaxInt64
			for _, v := range vs {
				if v < ans {
					ans = v
				}
			}
			return ans
		case 3:
			ans := math.MinInt64
			for _, v := range vs {
				if v > ans {
					ans = v
				}
			}
			return ans
		case 5:
			if vs[0] > vs[1] {
				return 1
			}
			return 0
		case 6:
			if vs[0] < vs[1] {
				return 1
			}
			return 0
		case 7:
			if vs[0] == vs[1] {
				return 1
			}
			return 0
		}
	}

	panic("help")
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
