package main

import (
	. "fmt"
	"sort"
)

// https://space.bilibili.com/206214
func cf835E() {
	var n, x, y, res, diff, p int
	Scan(&n, &x, &y)
	query := func(a []int) bool {
		k := len(a)
		Print("? ", k)
		for _, i := range a {
			Print(" ", i)
		}
		Println()
		Scan(&res)
		return k%2 == 0 && res == x^y || k%2 > 0 && res == y
	}
	for bit := 1; bit <= n; bit <<= 1 {
		a := []int{}
		for j := 1; j <= n; j++ {
			if j&bit > 0 {
				a = append(a, j)
			}
		}
		if !query(a) {
			continue
		}
		diff |= bit
		if p > 0 {
			continue
		}
		j := sort.Search(len(a)-1, func(i int) bool { return query(a[:i+1]) })
		p = a[j]
	}
	q := p ^ diff
	if p > q {
		p, q = q, p
	}
	Println("!", p, q)
}

//func main() { cf835E() }
