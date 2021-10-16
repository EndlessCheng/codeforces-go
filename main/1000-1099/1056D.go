package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1056D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		g[v-1] = append(g[v-1], w)
	}
	a := make([]int, 0, n)
	var f func(int) int
	f = func(v int) (l int) {
		for _, w := range g[v] {
			l += f(w)
		}
		if l == 0 {
			l = 1
		}
		a = append(a, l)
		return
	}
	f(0)
	sort.Ints(a)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF1056D(os.Stdin, os.Stdout) }
