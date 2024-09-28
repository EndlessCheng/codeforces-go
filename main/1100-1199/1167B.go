package main

import (
	. "fmt"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1167B() {
	mask := map[int]int{}
	a := []int{4, 8, 15, 16, 23, 42}
	for i, v := range a {
		for j, w := range a[:i] {
			mask[v*w] = 1<<i | 1<<j
		}
	}
	q := func(i, j int) (v int) {
		Println("?", i, j)
		Scan(&v)
		v = mask[v]
		return
	}
	m12 := q(1, 2)
	m34 := q(3, 4)
	m13 := q(1, 3)
	m15 := q(1, 5)
	m1 := m12 & m13
	f := func(i int) int { return a[bits.TrailingZeros(uint(i))] }
	Println("!", f(m1), f(m12^m1), f(m13^m1), f(m34^m13^m1), f(m1^m15), f(63^m12^m34^m1^m15))
}

//func main() { cf1167B() }
