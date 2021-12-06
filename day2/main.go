package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mainA() {
	scanner := bufio.NewScanner(os.Stdin)

	x := 0
	y := 0

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")

		op := words[0]
		val, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}

		switch op {
		case "forward":
			x += val
		case "down":
			y += val
		case "up":
			y -= val
		default:
			log.Fatal(op)
		}
	}

	fmt.Println(x, y, x*y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	x := 0
	aim := 0
	y := 0

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")

		op := words[0]
		val, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}

		switch op {
		case "forward":
			x += val
			y += val * aim
		case "down":
			aim += val
		case "up":
			aim -= val
		default:
			log.Fatal(op)
		}
	}

	fmt.Println(x, y, x*y)
}
