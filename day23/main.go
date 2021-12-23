package main

import (
	"container/heap"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const types = 4
const clones = 4

var (
	w, h int

	adjacent  [][]int
	goals     [types][clones]int
	stepCosts [types]int

	intermediate []bool

	initState State
)

func idx(r, c int) int {
	return r*w + c
}

func setup(grid []string) {
	h = len(grid)
	w = len(grid[0])

	for h := range grid {
		grid[h] += strings.Repeat(" ", w-len(grid[h]))
	}

	adjacent = make([][]int, h*w)

	intermediate = make([]bool, w*h)
	intermediate[idx(1, 1)] = true
	intermediate[idx(1, 2)] = true
	intermediate[idx(1, 4)] = true
	intermediate[idx(1, 6)] = true
	intermediate[idx(1, 8)] = true
	intermediate[idx(1, 10)] = true
	intermediate[idx(1, 11)] = true

	drs := []int{1, -1, 0, 0}
	dcs := []int{0, 0, 1, -1}

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] == '#' {
				continue
			}

			for i := range drs {
				dr, dc := drs[i], dcs[i]
				nr, nc := r+dr, c+dc
				if nr < 0 || nc < 0 || nr >= h || nc >= w {
					continue
				}

				if grid[nr][nc] == '#' {
					continue
				}

				adjacent[idx(r, c)] = append(adjacent[idx(r, c)], idx(nr, nc))
			}
		}
	}

	var locs [types][]int
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] >= 'A' && grid[r][c] <= 'D' {
				i := int(grid[r][c] - 'A')
				locs[i] = append(locs[i], idx(r, c))
			}
		}
	}

	for i := 0; i < types; i++ {
		for j := 0; j < clones; j++ {
			initState[i][j] = locs[i][j]
		}
	}

	for i := 0; i < types; i++ {
		for j := 0; j < clones; j++ {
			goals[i][j] = idx(2+j, 3+2*i)
		}
	}

	stepCosts[0] = 1
	for i := 1; i < types; i++ {
		stepCosts[i] = stepCosts[i-1] * 10
	}
}

type State [types][clones]int

type Move struct {
	state State
	dist  int
}

func moves(state State) []Move {
	occ := make([]bool, w*h)
	for _, locs := range state {
		for _, loc := range locs {
			occ[loc] = true
		}
	}

	var moves []Move

	for i := range state {
		settled := 0
		improved := true
		for improved && settled < clones {
			improved = false
			for _, loc := range state[i] {
				if loc == goals[i][clones-1-settled] {
					settled++
					improved = true
					break
				}
			}
		}
		if settled == clones {
			continue
		}
		nextGoal := goals[i][clones-1-settled]

		for j := range state[i] {
			start := state[i][j]

			best := make(map[int]int)
			best[start] = 0
			queue := []int{start}

			for len(queue) > 0 {
				cur := queue[0]
				dist := best[cur]
				queue = queue[1:]

				for _, nex := range adjacent[cur] {
					if occ[nex] {
						continue
					}
					if _, ok := best[nex]; ok {
						continue
					}
					best[nex] = dist + 1
					queue = append(queue, nex)
				}
			}

			for dst, dist := range best {
				if start == dst {
					continue
				}
				if intermediate[start] && dst != nextGoal {
					continue
				}
				if !intermediate[dst] && dst != nextGoal {
					continue
				}
				newState := state
				newState[i][j] = dst

				sort.Ints(newState[i][:])

				moves = append(moves, Move{
					state: newState,
					dist:  dist * stepCosts[i],
				})
			}
		}
	}

	return moves
}

func done(state State) bool {
	for i, locs := range state {
		for j, loc := range locs {
			if loc != goals[i][j] {
				return false
			}
		}
	}
	return true
}

type q struct {
	state State
	d     int
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
	setup(lines)

	seen := make(map[State]bool, h)
	dist := make(map[State]int, h)

	var pq pq
	heap.Init(&pq)

	heap.Push(&pq, q{state: initState, d: 0})
	seen[initState] = false
	dist[initState] = 0

	for len(pq) > 0 {
		cur := heap.Pop(&pq).(q)
		if seen[cur.state] {
			continue
		}
		seen[cur.state] = true

		if len(seen)%1000 == 0 {
			log.Println(len(seen), cur.d)
		}

		if done(cur.state) {
			log.Println(cur.d)
			break
		}

		for _, move := range moves(cur.state) {
			newState := move.state
			newDist := cur.d + move.dist

			if known, ok := dist[newState]; !ok || newDist < known {
				dist[newState] = newDist
				heap.Push(&pq, q{state: newState, d: newDist})
			}
		}
	}
}
