package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rotation struct {
	Ord  []int
	Flip []int
}

func apply(r Rotation, p []int) []int {
	return []int{
		r.Flip[0] * p[r.Ord[0]],
		r.Flip[1] * p[r.Ord[1]],
		r.Flip[2] * p[r.Ord[2]],
	}
}

func sub(a, b []int) []int {
	return []int{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func add(a, b []int) []int {
	return []int{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
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

func inRange(a []int) bool {
	return a[0] <= 1000 && a[0] >= -1000 &&
		a[1] <= 1000 && a[1] >= -1000 &&
		a[2] <= 1000 && a[2] >= -1000
}

func lookupMap(a [][]int) map[Pt]bool {
	m := make(map[Pt]bool)
	for _, pt := range a {
		m[toPt(pt)] = true
	}
	return m
}

type Chain struct {
	Prev   *Chain
	Offset []int
	Rot    Rotation
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var rotations []Rotation

	for i := 0; i < 3*3*3; i++ {
		ord := []int{i % 3, (i / 3) % 3, (i / 9) % 3}
		if ord[0] == ord[1] || ord[1] == ord[2] || ord[0] == ord[2] {
			continue
		}
		for j := 0; j < 2*2*2; j++ {
			flip := []int{(j%2)*2 - 1, ((j/2)%2)*2 - 1, ((j/4)%2)*2 - 1}
			rotations = append(rotations, Rotation{
				Ord:  ord,
				Flip: flip,
			})
		}
	}

	lines := strings.Split(string(input), "\n")

	scannerIdx := -1

	var scanners [][][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[1] == '-' {
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

	chains := make(map[int]*Chain)
	chains[0] = nil

	queue := []int{0}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		ptsA := scanners[i]
		aMap := lookupMap(ptsA)

	outer:
		for j := range scanners {
			if i == j {
				continue
			}

			if _, ok := chains[j]; ok {
				continue
			}

			for _, rot := range rotations {
				var ptsB [][]int
				for _, pt := range scanners[j] {
					ptsB = append(ptsB, apply(rot, pt))
				}
				bMap := lookupMap(ptsB)

				for _, k := range ptsA {
					for _, l := range ptsB {
						// identify A with B
						offset := sub(k, l)
						good := true
						cnt := 0

						for _, m := range ptsB {
							mapped := add(m, offset)

							if inRange(mapped) {
								if !aMap[toPt(mapped)] {
									good = false
									break
								}
								cnt++
							}
						}
						if cnt < 12 {
							good = false
						}

						offset = sub(l, k)
						cnt = 0

						for _, m := range ptsA {
							mapped := add(m, offset)

							if inRange(mapped) {
								if !bMap[toPt(mapped)] {
									good = false
									break
								}
								cnt++
							}
						}
						if cnt < 12 {
							good = false
						}

						if good {
							log.Println(i, j, sub(k, l), rot)
							queue = append(queue, j)
							chains[j] = &Chain{
								Prev:   chains[i],
								Offset: sub(k, l),
								Rot:    rot,
							}
							continue outer
						}
					}
				}
			}
		}
	}

	all := make(map[Pt]bool)
	for i := range scanners {
		for _, pt := range scanners[i] {
			ch := chains[i]
			for ch != nil {
				pt = apply(ch.Rot, pt)
				pt = add(pt, ch.Offset)
				ch = ch.Prev
			}
			all[toPt(pt)] = true
		}
	}

	locs := [][]int{}
	for i := range scanners {
		pt := []int{0, 0, 0}
		ch := chains[i]
		for ch != nil {
			pt = apply(ch.Rot, pt)
			pt = add(pt, ch.Offset)
			ch = ch.Prev
		}
		locs = append(locs, pt)
	}

	mx := 0
	for _, a := range locs {
		for _, b := range locs {
			d := sub(b, a)
			if d[0] < 0 {
				d[0] = -d[0]
			}
			if d[1] < 0 {
				d[1] = -d[1]
			}
			if d[2] < 0 {
				d[2] = -d[2]
			}
			dist := d[0] + d[1] + d[2]
			if dist > mx {
				mx = dist
			}
		}
	}
	log.Println(mx)
}
