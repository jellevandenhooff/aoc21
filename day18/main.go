package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Pair struct {
	Value       int
	Left, Right *Pair
}

func parseActual(x interface{}) *Pair {
	if v, ok := x.(float64); ok {
		return &Pair{
			Value: int(v),
		}
	}

	if a, ok := x.([]interface{}); ok {
		return &Pair{
			Left:  parseActual(a[0]),
			Right: parseActual(a[1]),
		}
	}

	panic("x")
}

func parse(line string) *Pair {
	var x interface{}
	if err := json.Unmarshal([]byte(line), &x); err != nil {
		panic(err)
	}
	return parseActual(x)
}

func (p *Pair) String() string {
	if p.Left == nil {
		return fmt.Sprint(p.Value)
	} else {
		return fmt.Sprintf("[%s,%s]", p.Left, p.Right)
	}
}

func addLeft(p *Pair, d int) {
	if p.Left == nil {
		p.Value += d
		return
	}

	addLeft(p.Left, d)
}

func addRight(p *Pair, d int) {
	if p.Left == nil {
		p.Value += d
		return
	}
	addRight(p.Right, d)
}

func explode(p **Pair, d int) (*int, *int, bool) {
	o := *p

	if o.Left == nil {
		return nil, nil, false
	}

	if d >= 4 {
		*p = &Pair{Value: 0}
		return &o.Left.Value, &o.Right.Value, true
	}

	a, b, ok := explode(&o.Left, d+1)
	if ok {
		if b != nil {
			addLeft(o.Right, *b)
		}
		if a != nil {
			return a, nil, true
		}
		return nil, nil, true
	}

	a, b, ok = explode(&o.Right, d+1)
	if ok {
		if a != nil {
			addRight(o.Left, *a)
		}
		if b != nil {
			return nil, b, true
		}
		return nil, nil, true
	}

	return nil, nil, false
}

func split(p **Pair) bool {
	o := *p

	if o.Left == nil {
		if o.Value >= 10 {
			*p = &Pair{
				Left:  &Pair{Value: o.Value / 2},
				Right: &Pair{Value: (o.Value + 1) / 2},
			}
			return true
		}
		return false
	}

	if split(&o.Left) {
		return true
	}

	if split(&o.Right) {
		return true
	}

	return false
}

func reduce(p *Pair) *Pair {
	for {
		if _, _, ok := explode(&p, 0); ok {
			continue
		}

		if split(&p) {
			continue
		}

		break
	}

	return p
}

func add(l, r *Pair) *Pair {
	return &Pair{
		Left:  l,
		Right: r,
	}
}

func magnitude(p *Pair) int {
	if p.Left == nil {
		return p.Value
	}
	return 3*magnitude(p.Left) + 2*magnitude(p.Right)
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	a := parse(lines[0])

	for _, b := range lines[1:] {
		a = reduce(add(a, parse(b)))
	}

	log.Println(magnitude(a))
	// log.Println(reduce(add(parse(a), parse(b))))

	biggest := -1

	for i, a := range lines {
		for j, b := range lines {
			if i == j {
				continue
			}
			v := magnitude(reduce(add(parse(a), parse(b))))
			if v > biggest {
				biggest = v
			}
		}
	}
	log.Println(biggest)
}
