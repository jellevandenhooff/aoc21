package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	window := 3
	answer := 0
	for i := 0; i+1+window <= len(numbers); i++ {
		if sum(numbers[i:i+window]) < sum(numbers[i+1:i+1+window]) {
			answer++
		}
	}

	fmt.Println(answer)
}
