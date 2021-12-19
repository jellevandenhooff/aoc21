package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rot struct {
	Ord  []int
	Flip []int
}

func applyRot(r Rot, p []int) []int {
	return []int{
		r.Flip[0] * p[r.Ord[0]],
		r.Flip[1] * p[r.Ord[1]],
		r.Flip[2] * p[r.Ord[2]],
	}
}

func chainRot(a, b Rot) Rot {
	return Rot{
		Flip: []int{
			a.Flip[0] * b.Flip[a.Ord[0]],
			a.Flip[1] * b.Flip[a.Ord[1]],
			a.Flip[2] * b.Flip[a.Ord[2]],
		},
		Ord: []int{
			b.Ord[a.Ord[0]],
			b.Ord[a.Ord[1]],
			b.Ord[a.Ord[2]],
		},
	}
}

func makeRotations() []Rot {
	var rotations []Rot
	for i := 0; i < 3*3*3; i++ {
		ord := []int{i % 3, (i / 3) % 3, (i / 9) % 3}
		if ord[0] == ord[1] || ord[1] == ord[2] || ord[0] == ord[2] {
			continue
		}
		for j := 0; j < 2*2*2; j++ {
			flip := []int{(j%2)*2 - 1, ((j/2)%2)*2 - 1, ((j/4)%2)*2 - 1}
			rotations = append(rotations, Rot{
				Ord:  ord,
				Flip: flip,
			})
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

type Transform struct {
	Offset []int
	Rot    Rot
}

func applyTransform(t Transform, pt []int) []int {
	pt = applyRot(t.Rot, pt)
	pt = add(pt, t.Offset)
	return pt
}

func chainTransforms(a, b Transform) Transform {
	return Transform{
		Offset: add(applyRot(a.Rot, b.Offset), a.Offset),
		Rot:    chainRot(a.Rot, b.Rot),
	}
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

	transforms := make(map[int]Transform)
	transforms[0] = Transform{
		Offset: []int{0, 0, 0},
		Rot: Rot{
			Ord:  []int{0, 1, 2},
			Flip: []int{1, 1, 1},
		},
	}

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
					ptsB = append(ptsB, applyRot(rot, pt))
				}
				if offset, ok := overlap(ptsA, ptsB); ok {
					queue = append(queue, j)
					transforms[j] = chainTransforms(transforms[i], Transform{
						Offset: offset,
						Rot:    rot,
					})
					break
				}
			}
		}
	}

	all := make(map[Pt]bool)
	for i := range scanners {
		for _, pt := range scanners[i] {
			all[toPt(applyTransform(transforms[i], pt))] = true
		}
	}
	log.Println(len(all))

	locs := [][]int{}
	for i := range scanners {
		locs = append(locs, transforms[i].Offset)
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
