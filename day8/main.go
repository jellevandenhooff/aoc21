package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

var truth = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

var inverted map[string]int

func init() {
	inverted = make(map[string]int)
	for k, v := range truth {
		inverted[v] = k
	}
}

var all = []string{"a", "b", "c", "d", "e", "f", "g"}

func apply(s string, m map[string]string) string {
	var r []string
	for i := range s {
		r = append(r, m[s[i:i+1]])
	}
	sort.Strings(r)
	return strings.Join(r, "")
}

type searcher struct {
	known   []string
	current map[string]string
	used    map[string]struct{}
}

func (s *searcher) search() map[string]string {
	if len(s.current) == len(all) {
		for _, w := range s.known {
			if _, ok := inverted[apply(w, s.current)]; !ok {
				return nil
			}
		}
		return s.current
	}

	from := all[len(s.current)]
	for _, to := range all {
		if _, ok := s.used[to]; ok {
			continue
		}

		s.used[to] = struct{}{}
		s.current[from] = to

		if r := s.search(); r != nil {
			return r
		}

		delete(s.current, from)
		delete(s.used, to)
	}

	return nil
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	ans1 := 0
	ans2 := 0

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		known := strings.Split(parts[0], " ")
		unknown := strings.Split(parts[1], " ")

		s := searcher{
			known:   known,
			current: make(map[string]string),
			used:    make(map[string]struct{}),
		}
		m := s.search()
		if m == nil {
			log.Fatal(m)
		}

		x := 0
		for _, w := range unknown {
			d := inverted[apply(w, m)]
			if d == 1 || d == 4 || d == 7 || d == 8 {
				ans1++
			}
			x = x*10 + d
		}
		ans2 += x
	}

	log.Println(ans1, ans2)
}
