package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type state struct {
	idx    int
	scores [2]int
	pos    [2]int
}

type ret struct {
	wins [2]int
}

var memo = make(map[state]ret)

var throws []int

func init() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				throws = append(throws, i+j+k)
			}
		}
	}
}

func count(idx int, scores [2]int, pos [2]int) ret {
	if scores[0] >= 21 {
		return ret{
			wins: [2]int{1, 0},
		}
	}

	if scores[1] >= 21 {
		return ret{
			wins: [2]int{0, 1},
		}
	}

	state := state{idx: idx, scores: scores, pos: pos}
	if known, ok := memo[state]; ok {
		return known
	}

	var ans ret
	for _, throw := range throws {
		nextPos := pos
		nextPos[idx] = (nextPos[idx] + throw) % 10
		nextScores := scores
		nextScores[idx] += nextPos[idx] + 1
		nextIdx := 1 - idx

		next := count(nextIdx, nextScores, nextPos)
		ans.wins[0] += next.wins[0]
		ans.wins[1] += next.wins[1]
	}

	memo[state] = ans
	return ans
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	a, _ := strconv.Atoi(strings.TrimPrefix(lines[0], "Player 1 starting position: "))
	b, _ := strconv.Atoi(strings.TrimPrefix(lines[1], "Player 2 starting position: "))

	idx := 0
	pos := []int{a - 1, b - 1}
	scores := []int{0, 0}
	next := 6
	rolls := 0

	for {
		pos[idx] = (pos[idx] + next) % 10
		next += 9
		rolls += 3

		scores[idx] += pos[idx] + 1

		if scores[idx] >= 1000 {
			log.Println(rolls, scores[1-idx], rolls*scores[1-idx])
			break
		}

		idx = 1 - idx
	}

	ans := count(0, [2]int{}, [2]int{a - 1, b - 1})
	log.Println(ans.wins[0], ans.wins[1])
	if ans.wins[0] > ans.wins[1] {
		log.Println(ans.wins[0])
	} else {
		log.Println(ans.wins[1])
	}
}
