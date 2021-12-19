package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mat [4][4]int

var Ident = Mat{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func applyMat(m Mat, p []int) []int {
	p = append(p, 1)
	q := make([]int, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			q[i] += m[i][j] * p[j]
		}
	}
	return q[:3]
}

func multMat(a, b Mat) Mat {
	var c Mat
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func translateMat(offset []int) Mat {
	return Mat{
		{1, 0, 0, offset[0]},
		{0, 1, 0, offset[1]},
		{0, 0, 1, offset[2]},
		{0, 0, 0, 1},
	}
}

func makeRotations() []Mat {
	var rotations []Mat
	for i := 0; i < 3*3*3; i++ {
		ord := []int{i % 3, (i / 3) % 3, (i / 9) % 3}
		if ord[0] == ord[1] || ord[1] == ord[2] || ord[0] == ord[2] {
			continue
		}
		for j := 0; j < 2*2*2; j++ {
			flip := []int{(j%2)*2 - 1, ((j/2)%2)*2 - 1, ((j/4)%2)*2 - 1}
			var m Mat
			for k := 0; k < 3; k++ {
				m[k][ord[k]] = flip[k]
			}
			m[3][3] = 1
			rotations = append(rotations, m)
		}
	}
	return rotations
}

func sub(a, b []int) []int {
	return []int{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func add(a, b []int) []int {
	return []int{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sz(x []int) int {
	return abs(x[0]) + abs(x[1]) + abs(x[2])
}

type Pt struct {
	X, Y, Z int
}

func toPt(a []int) Pt {
	return Pt{
		X: a[0],
		Y: a[1],
		Z: a[2],
	}
}

func lookupMap(a [][]int) map[Pt]bool {
	m := make(map[Pt]bool)
	for _, pt := range a {
		m[toPt(pt)] = true
	}
	return m
}

func inRange(a []int) bool {
	return a[0] <= 1000 && a[0] >= -1000 &&
		a[1] <= 1000 && a[1] >= -1000 &&
		a[2] <= 1000 && a[2] >= -1000
}

func scanWorks(offset []int, ptsB [][]int, aMap map[Pt]bool) bool {
	cnt := 0
	for _, m := range ptsB {
		mapped := add(m, offset)
		if inRange(mapped) {
			if !aMap[toPt(mapped)] {
				return false
			}
			cnt++
		}
	}
	return cnt >= 12
}

func overlap(ptsA, ptsB [][]int) ([]int, bool) {
	aMap := lookupMap(ptsA)
	bMap := lookupMap(ptsB)

	for _, ptA := range ptsA {
		for _, ptB := range ptsB {
			if !scanWorks(sub(ptA, ptB), ptsB, aMap) {
				continue
			}
			if !scanWorks(sub(ptB, ptA), ptsA, bMap) {
				continue
			}
			return sub(ptA, ptB), true
		}
	}

	return nil, false
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	scannerIdx := -1
	var scanners [][][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0:2] == "--" {
			scannerIdx++
			scanners = append(scanners, nil)
			continue
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		scanners[scannerIdx] = append(scanners[scannerIdx], []int{x, y, z})
	}

	rotations := makeRotations()

	transforms := make(map[int]Mat)
	transforms[0] = Ident

	queue := []int{0}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		ptsA := scanners[i]

		for j := range scanners {
			if i == j {
				continue
			}
			if _, ok := transforms[j]; ok {
				continue
			}

			for _, rot := range rotations {
				var ptsB [][]int
				for _, pt := range scanners[j] {
					ptsB = append(ptsB, applyMat(rot, pt))
				}
				if offset, ok := overlap(ptsA, ptsB); ok {
					transforms[j] = multMat(transforms[i], multMat(translateMat(offset), rot))
					queue = append(queue, j)
					break
				}
			}
		}
	}

	all := make(map[Pt]bool)
	for i := range scanners {
		for _, pt := range scanners[i] {
			all[toPt(applyMat(transforms[i], pt))] = true
		}
	}
	log.Println(len(all))

	locs := [][]int{}
	for i := range scanners {
		pt := applyMat(transforms[i], []int{0, 0, 0})
		locs = append(locs, pt)
	}

	mx := 0
	for _, a := range locs {
		for _, b := range locs {
			dist := sz(sub(b, a))
			if dist > mx {
				mx = dist
			}
		}
	}
	log.Println(mx)
}
