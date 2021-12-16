package main

import (
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	inp := lines[0]

	m := make(map[string]string)

	lines = lines[2:]
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		from := parts[0]
		to := parts[1]
		m[from] = to
	}

	st := make(map[string]int)
	for i := range inp {
		if i+2 <= len(inp) {
			st[inp[i:i+2]]++
		}
	}

	for k := 0; k < 40; k++ {
		newSt := make(map[string]int)
		for p, c := range st {
			if x, ok := m[p]; ok {
				newSt[p[0:1]+x] += c
				newSt[x+p[1:2]] += c
			} else {
				newSt[p] += c
			}
		}
		st = newSt
	}

	cnts := make(map[string]int)
	for p, c := range st {
		cnts[p[0:1]] += c
		cnts[p[1:2]] += c
	}

	cnts[inp[0:1]]++
	cnts[inp[len(inp)-1:]]++

	max := -1
	min := math.MaxInt64
	for _, c := range cnts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	log.Print((max - min) / 2)
}
