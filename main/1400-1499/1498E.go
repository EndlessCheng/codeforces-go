package main

import (
	. "fmt"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1498E() {
	var n, s, l int
	Scan(&n)
	a := make([]struct{ inDeg, i int }, n)
	for i := range a {
		Scan(&a[i].inDeg)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].inDeg < a[j].inDeg })
	ans, v, w := -1, -1, -1
	for r, p := range a {
		s += p.inDeg - r
		if s == 0 { // 拆出了一个 SCC
			if l != r && p.inDeg-a[l].inDeg > ans {
				ans, v, w = p.inDeg-a[l].inDeg, p.i, a[l].i
			}
			l = r + 1
		}
	}
	Println("!", v+1, w+1)
}

//func main() { CF1498E() }
