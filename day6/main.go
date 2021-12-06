package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func sim(a map[int]int) map[int]int {
	b := make(map[int]int)
	for x, n := range a {
		if x == 0 {
			b[6] += n
			b[8] += n
		} else {
			b[x-1] += n
		}
	}
	return b
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	counts := make(map[int]int)
	for _, s := range strings.Split(lines[0], ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		counts[x]++
	}

	for i := 0; i < 256; i++ {
		counts = sim(counts)
	}
	ans := 0
	for _, n := range counts {
		ans += n
	}

	fmt.Println(ans)
}
