package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

type pnt struct {
	x, y int
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	cnt := make(map[pnt]int)

	for _, line := range lines {
		groups := re.FindStringSubmatch(line)
		x1, err := strconv.Atoi(groups[1])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(groups[2])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(groups[3])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(groups[4])
		if err != nil {
			log.Fatal(err)
		}

		dx := x2 - x1
		dy := y2 - y1

		if dx != 0 && dy != 0 {
			// continue
		}

		if dx > 0 {
			dy /= dx
			dx /= dx
		}
		if dx < 0 {
			dy /= -dx
			dx /= -dx
		}
		if dy > 0 {
			dx /= dy
			dy /= dy
		}
		if dy < 0 {
			dx /= -dy
			dy /= -dy
		}

		x, y := x1, y1
		for {
			cnt[pnt{x: x, y: y}]++

			if x == x2 && y == y2 {
				break
			}
			x += dx
			y += dy
		}
	}

	ans := 0
	for _, cnt := range cnt {
		if cnt >= 2 {
			ans++
		}
	}

	fmt.Println(ans)
}
