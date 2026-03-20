package main

import (
	. "fmt"
	"slices"
)

// https://github.com/EndlessCheng
func cf1779E() {
	var n, x int
	Scan(&n)
	p := make([][2]int, n)
	for i := 1; i <= n; i++ {
		Print("? ", i, " ")
		for j := 1; j <= n; j++ {
			if i != j {
				Print(1)
			} else {
				Print(0)
			}
		}
		Println()
		Scan(&x)
		p[i-1] = [2]int{-x, i}
	}
	slices.SortFunc(p, func(a, b [2]int) int { return a[0] - b[0] })

	ans := make([]int, n+1)
	s := 0
	for i := 1; i <= n; i++ {
		ans[p[i-1][1]] = 1
		s -= p[i-1][0]
		if s == i*(i-1)/2+i*(n-i) {
			i = n
		}
	}
	Print("! ")
	for _, v := range ans[1:] {
		Print(v)
	}
	Println()
}

//func main() { cf1779E() }
