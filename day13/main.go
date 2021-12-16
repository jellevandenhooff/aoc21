package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pt struct {
	X, Y int
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	var pts []*Pt

	coords := true
	for _, line := range lines {
		if line == "" {
			coords = false
			continue
		}

		if coords {
			nums := strings.Split(line, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			pts = append(pts, &Pt{X: x, Y: y})
		} else {
			line = strings.TrimPrefix(line, "fold along ")
			prts := strings.Split(line, "=")
			idx, _ := strconv.Atoi(prts[1])
			switch prts[0] {
			case "x":
				for _, pt := range pts {
					if pt.X >= idx {
						pt.X = idx - (pt.X - idx)
					}
				}
			case "y":
				for _, pt := range pts {
					if pt.Y >= idx {
						pt.Y = idx - (pt.Y - idx)
					}
				}
			}
			// break
		}
	}

	/*
		m := make(map[Pt]struct{})
		for _, pt := range pts {
			m[*pt] = struct{}{}
		}
		log.Println(len(m))
	*/

	sz := 0
	for _, pt := range pts {
		if pt.X > sz {
			sz = pt.X
		}
		if pt.Y > sz {
			sz = pt.Y
		}
	}

	g := make([][]byte, sz+1)
	for r := range g {
		g[r] = make([]byte, sz+1)
		for c := range g[r] {
			g[r][c] = byte('.')
		}
	}

	for _, pt := range pts {
		g[pt.Y][pt.X] = '#'
	}

	for r := range g {
		log.Println(string(g[r]))
	}
}
