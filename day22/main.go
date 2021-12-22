package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Cube struct {
	On                     bool
	x1, x2, y1, y2, z1, z2 int
}

type Compressor struct {
	Values []int
	Lookup map[int]int
}

func (c *Compressor) Add(x int) {
	c.Values = append(c.Values, x)
	c.Values = append(c.Values, x+1)
}

func (c *Compressor) Compress() {
	sort.Ints(c.Values)
	c.Lookup = make(map[int]int)
	for i, v := range c.Values {
		c.Lookup[v] = i
	}
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	var cubes []*Cube

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

		cube := &Cube{
			On: on,
			x1: x1,
			x2: x2,
			y1: y1,
			y2: y2,
			z1: z1,
			z2: z2,
		}

		/*
			good := false
			if cube.x1 >= -50 && cube.x2 <= 50 && cube.y1 >= -50 && cube.y2 <= 50 && cube.z1 >= -50 && cube.z2 <= 50 {
				good = true
			}
			if !good {
				continue
			}
		*/

		cubes = append(cubes, cube)
	}

	xC := &Compressor{}
	yC := &Compressor{}
	zC := &Compressor{}
	for _, cube := range cubes {
		xC.Add(cube.x1)
		xC.Add(cube.x2)
		yC.Add(cube.y1)
		yC.Add(cube.y2)
		zC.Add(cube.z1)
		zC.Add(cube.z2)
	}
	xC.Compress()
	yC.Compress()
	zC.Compress()
	for _, cube := range cubes {
		cube.x1 = xC.Lookup[cube.x1]
		cube.x2 = xC.Lookup[cube.x2]
		cube.y1 = yC.Lookup[cube.y1]
		cube.y2 = yC.Lookup[cube.y2]
		cube.z1 = zC.Lookup[cube.z1]
		cube.z2 = zC.Lookup[cube.z2]
	}

	m := make([][][]bool, len(xC.Values))
	for x := range m {
		m[x] = make([][]bool, len(yC.Values))
		for y := range m[x] {
			m[x][y] = make([]bool, len(zC.Values))
		}
	}

	for _, cube := range cubes {
		for x := cube.x1; x <= cube.x2; x++ {
			for y := cube.y1; y <= cube.y2; y++ {
				for z := cube.z1; z <= cube.z2; z++ {
					m[x][y][z] = cube.On
				}
			}
		}
	}

	count := 0
	for x := range m {
		for y := range m[x] {
			for z := range m[x][y] {
				if m[x][y][z] {
					count += (xC.Values[x+1] - xC.Values[x]) *
						(yC.Values[y+1] - yC.Values[y]) *
						(zC.Values[z+1] - zC.Values[z])
				}
			}
		}
	}
	log.Println(count)
}
