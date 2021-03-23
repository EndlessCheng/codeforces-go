package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF582A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n int
	Fscan(in, &n)
	m := n * n
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	if n == 1 {
		Fprint(out, a[0])
		return
	}
	sort.Ints(a)
	ans := []int{a[m-2], a[m-1]}
	del := map[int]int{gcd(a[m-2], a[m-1]): 2}
	for i := m - 3; i >= 0; i-- {
		if v := a[i]; del[v] > 0 {
			del[v]--
		} else {
			for _, w := range ans {
				del[gcd(v, w)] += 2
			}
			ans = append(ans, v)
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF582A(os.Stdin, os.Stdout) }
