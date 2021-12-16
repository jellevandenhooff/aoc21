package main

import (
	"container/heap"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type q struct {
	r, c int
	d    int
}

type pq []q

func (pq pq) Len() int {
	return len(pq)
}

func (pq pq) Less(a, b int) bool {
	return pq[a].d < pq[b].d
}

func (pq pq) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq *pq) Push(x interface{}) {
	*pq = append(*pq, x.(q))
}

func (pq *pq) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	ih := len(lines)
	iw := len(lines[0])

	h := ih * 5
	w := iw * 5

	g := make([][]int, h)
	seen := make([][]bool, h)
	cost := make([][]int, h)

	for r := range g {
		g[r] = make([]int, w)
		seen[r] = make([]bool, w)
		cost[r] = make([]int, w)
		for c := range g[r] {
			g[r][c] = (int(lines[r%ih][c%iw]-'0') + (r / ih) + (c / iw))
			for g[r][c] >= 10 {
				g[r][c] -= 9
			}
			cost[r][c] = math.MaxInt64
		}
	}

	drs := []int{1, -1, 0, 0}
	dcs := []int{0, 0, 1, -1}

	var pq pq
	heap.Init(&pq)

	cost[0][0] = 0
	heap.Push(&pq, q{r: 0, c: 0, d: 0})

	for !seen[h-1][w-1] {
		cur := heap.Pop(&pq).(q)
		if seen[cur.r][cur.c] {
			continue
		}

		best := cur.d
		br, bc := cur.r, cur.c

		seen[br][bc] = true
		for i := range drs {
			dr, dc := drs[i], dcs[i]
			r, c := br+dr, bc+dc
			if r < 0 || r >= h || c < 0 || c >= w {
				continue
			}

			if best+g[r][c] < cost[r][c] {
				cost[r][c] = best + g[r][c]
				heap.Push(&pq, q{r: r, c: c, d: cost[r][c]})
			}
		}
	}

	log.Print(cost[h-1][w-1])
}
