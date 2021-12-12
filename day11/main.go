package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	g := make([][]int, 10)
	for r := range g {
		g[r] = make([]int, 10)
		for c := range g[r] {
			g[r][c] = int(lines[r][c] - '0')
		}
	}

	type Pt struct{ r, c int }

	ans := 0
	ans2 := 0

	for step := 0; ; step++ {
		var q []Pt

		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				g[r][c]++

				if g[r][c] == 10 {
					q = append(q, Pt{r: r, c: c})
				}
			}
		}

		cnt := 0

		for len(q) > 0 {
			r, c := q[0].r, q[0].c
			q = q[1:]

			if step < 100 {
				ans++
			}
			cnt++

			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					nr, nc := r+dr, c+dc
					if nr < 0 || nr >= 10 || nc < 0 || nc >= 10 {
						continue
					}

					g[nr][nc]++
					if g[nr][nc] == 10 {
						q = append(q, Pt{r: nr, c: nc})
					}
				}
			}
		}

		if cnt == 100 {
			ans2 = step
			break
		}

		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				if g[r][c] >= 10 {
					g[r][c] = 0
				}
			}
		}
	}

	log.Println(ans)
	log.Println(ans2 + 1)
}
