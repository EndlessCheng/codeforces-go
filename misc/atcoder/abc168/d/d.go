package main

import (
	"bufio"
	. "fmt"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func main() {
	r := bufio.NewReader(os.Stdin)
	o := bufio.NewWriter(os.Stdout)
	var n, m, v, w int
	Fscan(r, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(r, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	a := make([]int, n+1)
	c := 1
	s := make([]bool, n+1)
	s[1] = true
	q := []int{1}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, w := range g[v] {
			if !s[w] {
				s[w] = true
				c++
				a[w] = v
				q = append(q, w)
			}
		}
	}
	if c < n {
		Fprint(o, "No")
		return
	}
	Fprintln(o, "Yes")
	for _, v := range a[2:] {
		Fprintln(o, v)
	}
	o.Flush()
}
