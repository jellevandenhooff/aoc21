package main

import (
	"io"
	"log"
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
	line := lines[0]

	line = strings.TrimPrefix(line, "target area: x=")

	parts := strings.SplitN(line, "..", 2)
	x1, _ := strconv.Atoi(parts[0])
	line = parts[1]

	parts = strings.SplitN(line, ", y=", 2)
	x2, _ := strconv.Atoi(parts[0])
	line = parts[1]

	parts = strings.SplitN(line, "..", 2)
	y1, _ := strconv.Atoi(parts[0])
	y2, _ := strconv.Atoi(parts[1])

	log.Println(x1, x2, y1, y2)

	bestDYByStep := make(map[int]int)
	highestYByStep := make(map[int]int)
	allByStep := make(map[int][]int)

	for initDy := -10000; initDy <= 10000; initDy++ {
		dy := initDy
		y := 0

		highest := 0

		for step := 0; step < 10000; step++ {
			y += dy
			dy--

			if y > highest {
				highest = y
			}

			if y >= y1 && y <= y2 {
				if prev, ok := highestYByStep[step]; !ok || highest > prev {
					highestYByStep[step] = highest
					bestDYByStep[step] = initDy
				}
				allByStep[step] = append(allByStep[step], initDy)
			}
		}
	}

	total := 0
	best := 0

	for initDx := 1; initDx <= x2; initDx++ {
		dx := initDx
		x := 0

		works := make(map[int]bool)

		for step := 0; step < 10000; step++ {
			x += dx
			if dx > 0 {
				dx--
			}

			if x >= x1 && x <= x2 {
				if highest, ok := highestYByStep[step]; ok {
					if highest > best {
						best = highest
						log.Println(initDx, bestDYByStep[step], highest)
					}
				}
				for _, initDy := range allByStep[step] {
					works[initDy] = true
				}
			}
		}

		total += len(works)
	}

	log.Println(total)
}
