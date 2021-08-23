package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF242D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	q := []int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] == 0 {
			q = append(q, i)
		}
	}

	ans := []int{}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		ans = append(ans, v+1)
		for _, w := range g[v] {
			if a[w]--; a[w] == 0 {
				q = append(q, w)
			}
		}
	}
	sort.Ints(ans)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF242D(os.Stdin, os.Stdout) }
