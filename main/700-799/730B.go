package main

import (
	. "fmt"
	"math/bits"
)

// https://github.com/EndlessCheng
func less30(i, j int) bool {
	Println("?", i, j)
	c := ""
	Scan(&c)
	return c == "<"
}

func dfs30(l, r int) (mn, mx int) {
	sz := r - l
	if sz == 1 {
		return l, l
	}
	if sz == 2 {
		if less30(l, l+1) {
			return l, l + 1
		}
		return l + 1, l
	}
	m := l + sz/2
	if sz&(sz-1) > 0 {
		m = l + 1<<(bits.Len(uint(sz))-1)
	}
	mn, mx = dfs30(l, m)
	mn2, mx2 := dfs30(m, r)
	if less30(mn2, mn) {
		mn = mn2
	}
	if less30(mx, mx2) {
		mx = mx2
	}
	return
}

func cf730B() {
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		mn, mx := dfs30(1, n+1)
		Println("!", mn, mx)
	}
}

//func main() { cf730B() }
