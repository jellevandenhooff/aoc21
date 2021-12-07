package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	pos := make([]int, 0)
	for _, s := range strings.Split(lines[0], ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		pos = append(pos, x)
	}

	best := 0
	bestCost := math.MaxInt64

	for i := 0; i < 1000000; i++ {
		cost := 0
		for _, p := range pos {
			d := p - i
			if d < 0 {
				d = -d
			}
			cost += d * (d + 1) / 2
		}
		if cost < bestCost {
			best = i
			bestCost = cost
		}
	}

	fmt.Println(best, bestCost)
}
