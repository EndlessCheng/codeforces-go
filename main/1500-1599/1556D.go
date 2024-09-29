package main

import (
	. "fmt"
	"slices"
)

// https://github.com/EndlessCheng
func cf1556D() {
	var n, k, v, w int
	Scan(&n, &k)
	q := func(i, j int) int {
		Println("or", i, j)
		Scan(&v)
		Println("and", i, j)
		Scan(&w)
		return v + w
	}
	a := make([]int, n)
	s12 := q(1, 2)
	s13 := q(1, 3)
	a[0] = (s12 + s13 - q(2, 3)) / 2
	a[1] = s12 - a[0]
	a[2] = s13 - a[0]
	for i := 3; i < n; i++ {
		a[i] = q(1, i+1) - a[0]
	}
	slices.Sort(a)
	Println("finish", a[k-1])
}

//func main() { cf1556D() }
