package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func apply(grid [][]int, table []int) [][]int {
	ret := make([][]int, len(grid))
	for r := range ret {
		ret[r] = make([]int, len(grid[0]))
	}

	for r := range grid {
		for c := range grid[0] {
			if r-1 < 0 || r+1 >= len(grid) || c-1 < 0 || c+1 >= len(grid[0]) {
				if grid[r][c] == 0 {
					ret[r][c] = table[0]
				} else {
					ret[r][c] = table[(1<<9)-1]
				}
				continue
			}

			idx := 0
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					idx = (idx << 1) | grid[r+dr][c+dc]
				}
			}
			ret[r][c] = table[idx]
		}
	}

	return ret
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	table := make([]int, 512)
	for i, v := range lines[0] {
		if v == '#' {
			table[i] = 1
		}
	}

	img := lines[2:]

	steps := 50
	offset := steps + 2
	sz := len(img) + 2*offset
	grid := make([][]int, sz)
	for r := range grid {
		grid[r] = make([]int, sz)
	}

	for r := 0; r < len(img); r++ {
		for c := 0; c < len(img); c++ {
			if img[r][c] == '#' {
				grid[r+offset][c+offset] = 1
			}
		}
	}

	for i := 0; i < steps; i++ {
		grid = apply(grid, table)
	}

	cnt := 0
	for r := range grid {
		for c := range grid[r] {
			cnt += grid[r][c]
		}
	}
	log.Println(cnt)
}
