package main

import (
	. "fmt"
	"sort"
	. "strings"
)

// https://space.bilibili.com/206214
func CF1838F() {
	var n, x, y int
	Scan(&n)
	U := Repeat("^", n)
	D := Repeat("v", n)
	L := Repeat("<", n-1)
	R := Repeat(">", n-1)
	r := sort.Search(n-1, func(m int) bool {
		Println("?", 1, 1)
		for i := 0; i < m; i++ {
			if i%2 == 0 {
				Println(R + "v")
			} else {
				Println("v" + L)
			}
		}
		Println(Repeat(string("><"[m%2]), n))
		for i := m + 1; i < n; i++ {
			Println(D)
		}
		Scan(&x, &y)
		if x != m+1 || y != []int{n + 1, 0}[m%2] {
			return true
		}

		Println("?", m+1, []int{n, 1}[m%2])
		Println(L + "<")
		for i := 1; i <= m; i++ {
			if i%2 == 0 {
				Println("^" + L)
			} else {
				Println(R + "^")
			}
		}
		for i := m + 1; i < n; i++ {
			Println(D)
		}
		Scan(&x, &y)
		return x != 1 || y != 0
	}) + 1

	Println("?", r, 1)
	for i := 1; i < r; i++ {
		Println(U)
	}
	Println(R + ">")
	for i := r + 1; i <= n; i++ {
		Println(D)
	}
	Scan(&x, &y)
	if x == 0 {
		Println("!", r, y, "^")
		return
	}
	if x > n {
		Println("!", r, y, "v")
		return
	}

	ans := '>'
	if y <= n {
		ans = '<'
	}
	c := sort.Search(n-1, func(m int) bool {
		if ans == '>' {
			m += 2
			Println("?", r, m)
		} else {
			Println("?", r, 1)
		}
		for i := 1; i < r; i++ {
			Println(U)
		}
		Println(Repeat(string(ans^2), m) + Repeat("^", n-m))
		for i := r + 1; i <= n; i++ {
			Println(D)
		}
		Scan(&x, &y)
		return ans == '>' && x == -1 || ans == '<' && x != 0
	}) + 1
	Println("!", r, c, string(ans))
}

//func main() { CF1838F() }
