package main

import (
	. "fmt"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1010B() {
	var m, n, t int
	Scan(&m, &n)
	a := make([]bool, n)
	for i := range a {
		Println(1)
		Scan(&t)
		if t == 0 {
			return
		}
		a[i] = t == 1
	}
	cnt := 0
	sort.Search(m+1, func(x int) bool {
		if x == 0 {
			return false
		}
		Println(x)
		Scan(&t)
		if t == 0 {
			os.Exit(0)
		}
		if !a[cnt%n] {
			t = -t
		}
		cnt++
		return t == -1
	})
}

//func main() { CF1010B() }
