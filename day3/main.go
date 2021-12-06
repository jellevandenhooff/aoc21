package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main0() {
	scanner := bufio.NewScanner(os.Stdin)

	first := true
	var counts [][2]int

	for scanner.Scan() {
		line := scanner.Text()

		if first {
			counts = make([][2]int, len(line))
			first = false
		}

		for i, c := range line {
			counts[i][c-'0'] += 1
		}
	}

	gamma := 0
	epsilon := 0

	for i := range counts {
		var common int
		if counts[i][0] > counts[i][1] {
			common = 0
		} else {
			common = 1
		}

		gamma = gamma*2 + common
		epsilon = epsilon*2 + (1 - common)
	}

	fmt.Println(gamma, epsilon, gamma*epsilon)
}

func count(a []string, i int) int {
	n := 0
	for _, s := range a {
		if s[i] == '0' {
			n++
		}
	}
	return n
}

func filter(a []string, i int, f bool) []string {
	var r []string
	for _, s := range a {
		if (s[i] == '0') == f {
			r = append(r, s)
		}
	}
	return r
}

func find(a []string, i int, f bool) string {
	if len(a) == 1 {
		return a[0]
	}
	a = filter(a, i, (count(a, i) > len(a)/2) != f)
	return find(a, i+1, f)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []string
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	a, _ := strconv.ParseInt(find(numbers, 0, false), 2, 64)
	b, _ := strconv.ParseInt(find(numbers, 0, true), 2, 64)

	fmt.Println(a, b, a*b)
}
