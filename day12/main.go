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

	edges := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			log.Fatal(line)
		}
		a, b := parts[0], parts[1]
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}

	seen := make(map[string]bool)
	double := false

	var count func(cur string) int
	count = func(cur string) int {
		if strings.ToLower(cur) == cur {
			if seen[cur] {
				double = true
				defer func() {
					double = false
				}()
			} else {
				seen[cur] = true
				defer func() {
					seen[cur] = false
				}()
			}
		}

		if cur == "end" {
			return 1
		}

		ans := 0
		for _, next := range edges[cur] {
			if strings.ToLower(next) == next && seen[next] && (next == "start" || double) {
				continue
			}
			ans += count(next)
		}
		return ans
	}

	log.Println(count("start"))
}
