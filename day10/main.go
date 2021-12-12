package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

/*
): 3 points.
]: 57 points.
}: 1197 points.
>: 25137 points.
*/

var m map[rune]rune = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var s map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var s2 map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	ans := 0
	var nums []int

	for _, line := range lines {
		var stack []rune
		good := true
		for _, c := range line {
			ok := true
			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')', ']', '}', '>':
				if len(stack) == 0 {
					ok = false
					break
				}
				if m[stack[len(stack)-1]] != c {
					ok = false
					break
				}
				stack = stack[:len(stack)-1]
			default:
				log.Fatal(c)
			}
			if !ok {
				ans += s[c]
				good = false
				break
			}
		}

		num := 0
		if good {
			for i := len(stack) - 1; i >= 0; i-- {
				num *= 5
				num += s2[stack[i]]
				// fmt.Printf("%c %d\n", stack[i], num)
			}
			// fmt.Println()
			nums = append(nums, num)
			// fmt.Println(num)
		}
	}

	sort.Ints(nums)

	log.Println(ans)
	log.Println(nums[len(nums)/2])
}
