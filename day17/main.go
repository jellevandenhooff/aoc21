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

	total := 0
	allHighest := 0

	for initDx := 1; initDx <= x2; initDx++ {
		for initDy := -200; initDy <= 200; initDy++ {
			dx := initDx
			dy := initDy
			x := 0
			y := 0

			highest := 0
			counted := false

			for step := 0; step < 10000; step++ {
				y += dy
				dy--

				x += dx
				if dx > 0 {
					dx--
				}

				if y > highest {
					highest = y
				}

				if x > x2 {
					break
				}
				if x < x1 && dx <= 0 {
					break
				}
				if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
					if highest > allHighest {
						allHighest = highest
					}
					if !counted {
						log.Println(initDy)
						counted = true
						total++
					}
				}
			}
		}
	}

	log.Println(total, allHighest)
}
