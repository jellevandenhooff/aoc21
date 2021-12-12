package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	grid := lines
	w := len(grid[0])
	h := len(grid)

	drs := []int{-1, 1, 0, 0}
	dcs := []int{0, 0, -1, 1}

	seen := make([][]bool, h)
	for r := range seen {
		seen[r] = make([]bool, w)
	}

	var find func(r, c int) int
	find = func(r, c int) int {
		if seen[r][c] {
			return 0
		}
		seen[r][c] = true

		n := 1
		for i := range drs {
			dr, dc := drs[i], dcs[i]
			nr, nc := r+dr, c+dc
			if nr < 0 || nr >= h || nc < 0 || nc >= w {
				continue
			}

			if grid[r][c] <= grid[nr][nc] && grid[nr][nc] != '9' {
				n += find(nr, nc)
			}
		}
		return n
	}

	ans1 := 0
	sizes := make([]int, 0)

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			ok := true

			for i := range drs {
				dr, dc := drs[i], dcs[i]
				nr, nc := r+dr, c+dc
				if nr < 0 || nr >= h || nc < 0 || nc >= w {
					continue
				}

				if grid[r][c] >= grid[nr][nc] {
					ok = false
					break
				}
			}

			if ok {
				ans1 += int(grid[r][c]-'0') + 1
				sizes = append(sizes, find(r, c))
			}
		}
	}

	log.Println(ans1)
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	log.Println(sizes[0] * sizes[1] * sizes[2])
}
