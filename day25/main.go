package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func step(grid [][]byte) ([][]byte, bool) {
	h := len(grid)
	w := len(grid[0])

	moved := false

	newGrid := make([][]byte, h)
	for r := range grid {
		newGrid[r] = make([]byte, w)
		for c := range newGrid[r] {
			newGrid[r][c] = '.'
		}
	}

	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == '>' && grid[r][(c+1)%w] == '.' {
				newGrid[r][(c+1)%w] = '>'
				moved = true
			} else if grid[r][c] == '>' {
				newGrid[r][c] = '>'
			} else if grid[r][c] == 'v' {
				newGrid[r][c] = 'v'
			}
		}
	}

	grid = newGrid
	newGrid = make([][]byte, h)
	for r := range grid {
		newGrid[r] = make([]byte, w)
		for c := range newGrid[r] {
			newGrid[r][c] = '.'
		}
	}

	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 'v' && grid[(r+1)%h][c] == '.' {
				newGrid[(r+1)%h][c] = 'v'
				moved = true
			} else if grid[r][c] == 'v' {
				newGrid[r][c] = 'v'
			} else if grid[r][c] == '>' {
				newGrid[r][c] = '>'
			}
		}
	}

	return newGrid, moved
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	grid := make([][]byte, len(lines))
	for r, line := range lines {
		grid[r] = []byte(line)
	}

	for idx := 1; ; idx++ {
		newGrid, moved := step(grid)
		if !moved {
			log.Println(idx)
			break
		}
		grid = newGrid
	}
}
