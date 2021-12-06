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

func readNumbers(b []string) []int {
	a := make([]int, 0, len(b))
	for _, w := range b {
		n, err := strconv.Atoi(w)
		if err != nil {
			log.Println(b)
			log.Fatal(err)
		}
		a = append(a, n)
	}
	return a
}

func readBoard(b []string) [][]int {
	a := make([][]int, 0, len(b))
	for _, l := range b {
		a = append(a, readNumbers(regexp.MustCompile(` +`).Split(strings.Trim(l, " "), -1)))
	}
	return a
}

func score(board [][]int, numbers []int) (int, int) {
	done := make([][]bool, len(board))
	for i := range done {
		done[i] = make([]bool, len(board[0]))
	}

	for idx, number := range numbers {
		for r := range board {
			for c := range board[0] {
				if board[r][c] == number {
					done[r][c] = true
				}
			}
		}

		okok := false

		for r := range board {
			ok := true
			for c := range board[0] {
				if !done[r][c] {
					ok = false
					break
				}
			}
			if ok {
				okok = true
			}
		}

		for c := range board[0] {
			ok := true
			for r := range board {
				if !done[r][c] {
					ok = false
					break
				}
			}
			if ok {
				okok = true
			}
		}

		if okok {
			sum := 0
			for r := range board {
				for c := range board[0] {
					if !done[r][c] {
						sum += board[r][c]
					}
				}
			}
			return idx, sum
		}
	}

	return -1, -1
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	numbers := readNumbers(strings.Split(lines[0], ","))
	lines = lines[2:]

	var boards [][][]int
	for len(lines) >= 5 {
		boards = append(boards, readBoard(lines[:5]))
		lines = lines[6:]
	}

	best := 0 // len(boards[0])*len(boards[0][0]) + 1
	var bestScore int

	for _, board := range boards {
		time, score := score(board, numbers)
		if time > best {
			best = time
			bestScore = score
		}
	}

	fmt.Println(best, numbers[best], bestScore, numbers[best]*bestScore)
}
