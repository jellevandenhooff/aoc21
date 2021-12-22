package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	value                  int
	x1, x2, y1, y2, z1, z2 int
}

func intersect(a1, a2, b1, b2 int) (int, int, bool) {
	c1 := a1
	if b1 > c1 {
		c1 = b1
	}
	c2 := a2
	if b2 < c2 {
		c2 = b2
	}
	if c1 > c2 {
		return 0, 0, false
	}
	return c1, c2, true
}

func Intersect(a, b Cube) (Cube, bool) {
	c := Cube{}
	ok := true

	c.x1, c.x2, ok = intersect(a.x1, a.x2, b.x1, b.x2)
	if !ok {
		return Cube{}, false
	}

	c.y1, c.y2, ok = intersect(a.y1, a.y2, b.y1, b.y2)
	if !ok {
		return Cube{}, false
	}

	c.z1, c.z2, ok = intersect(a.z1, a.z2, b.z1, b.z2)
	if !ok {
		return Cube{}, false
	}

	return c, true
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	var cubes []Cube

	for _, line := range lines {
		parts := strings.SplitN(line, " x=", 2)
		on := parts[0] == "on"
		line = parts[1]

		parts = strings.SplitN(line, "..", 2)
		x1, _ := strconv.Atoi(parts[0])
		line = parts[1]

		parts = strings.SplitN(line, ",y=", 2)
		x2, _ := strconv.Atoi(parts[0])
		line = parts[1]

		parts = strings.SplitN(line, "..", 2)
		y1, _ := strconv.Atoi(parts[0])
		line = parts[1]

		parts = strings.SplitN(line, ",z=", 2)
		y2, _ := strconv.Atoi(parts[0])
		line = parts[1]

		parts = strings.SplitN(line, "..", 2)
		z1, _ := strconv.Atoi(parts[0])

		z2, _ := strconv.Atoi(parts[1])

		value := 0
		if on {
			value = 1
		}

		cube := Cube{
			value: value,
			x1:    x1,
			x2:    x2,
			y1:    y1,
			y2:    y2,
			z1:    z1,
			z2:    z2,
		}

		good := false
		if cube.x1 >= -50 && cube.x2 <= 50 && cube.y1 >= -50 && cube.y2 <= 50 && cube.z1 >= -50 && cube.z2 <= 50 {
			good = true
		}
		if !good {
			continue
		}

		cubes = append(cubes, cube)
	}

	current := []Cube{}

	for _, cube := range cubes {
		for _, other := range current {
			c, ok := Intersect(cube, other)
			if ok {
				c.value = -other.value
				current = append(current, c)
			}
		}
		if cube.value != 0 {
			current = append(current, cube)
		}
	}

	answer := 0
	for _, cube := range current {
		answer += cube.value *
			(1 + cube.x2 - cube.x1) *
			(1 + cube.y2 - cube.y1) *
			(1 + cube.z2 - cube.z1)
	}
	log.Println(answer)
}
