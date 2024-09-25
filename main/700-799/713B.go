package main

import (
	. "fmt"
	"sort"
)

// https://github.com/EndlessCheng
func cf713B() {
	q := func(x1, y1, x2, y2 int) (c int) {
		Println("?", x1+1, y1+1, x2+1, y2+1)
		Scan(&c)
		return
	}
	var n int
	Scan(&n)

	d := sort.Search(n-1, func(i int) bool { return q(0, 0, i, n-1) == 2 })
	u := sort.Search(d, func(i int) bool { return q(i+1, 0, d, n-1) < 2 })
	r2 := sort.Search(n-1, func(j int) bool { return q(u, 0, d, j) == 2 })
	r1 := sort.Search(r2, func(j int) bool { return q(u, 0, d, j) > 0 })

	if r1 == r2 {
		d1 := sort.Search(d, func(i int) bool { return q(0, 0, i, r1) > 0 })
		l1 := sort.Search(r1, func(j int) bool { return q(u, j+1, d1, r1) == 0 })

		u2 := sort.Search(d, func(i int) bool { return q(i+1, 0, d, r2) == 0 })
		l2 := sort.Search(r2, func(j int) bool { return q(u2, j+1, d, r2) == 0 })

		Println("!", u+1, l1+1, d1+1, r1+1, u2+1, l2+1, d+1, r2+1)
		return
	}

	l1 := sort.Search(r1, func(j int) bool { return q(u, j+1, d, r1) == 0 })
	d1 := sort.Search(d, func(i int) bool { return q(0, l1, i, r1) > 0 })
	u1 := sort.Search(d1, func(i int) bool { return q(i+1, l1, d1, r1) == 0 })

	l2 := sort.Search(r2, func(j int) bool { return q(u, j+1, d, r2) == 0 })
	u2 := sort.Search(d, func(i int) bool { return q(i+1, l2, d, r2) == 0 })
	d2 := u2 + sort.Search(d-u2, func(i int) bool { return q(u2, l2, u2+i, r2) > 0 })

	Println("!", u1+1, l1+1, d1+1, r1+1, u2+1, l2+1, d2+1, r2+1)
}

//func main() { cf713B() }
