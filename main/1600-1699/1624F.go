package main

import (
	. "fmt"
	"sort"
)

// https://github.com/EndlessCheng
func cf1624F() {
	var n, resp int
	Scan(&n)
	add := 1
	ans := sort.Search(n-2, func(m int) bool {
		// 二分过程保证 m2 不会和 n-1 同余
		m2 := m + add
		c := n - 1 - m2%n
		add += c
		Println("+", c)
		Scan(&resp)
		return resp == m2/n
	})
	Println("!", ans+add)
}

//func main() { cf1624F() }
