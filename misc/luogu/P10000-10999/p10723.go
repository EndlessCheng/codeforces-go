package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p10723(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	ans := n
	for i := range a {
		Fscan(in, &a[i])
		ans -= a[i]
	}

	g := make([][]int, n)
	deg := make([]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 && a[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		ans--
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if a[w] > 0 {
				continue
			}
			deg[w]--
			if deg[w] == 1 {
				q = append(q, w)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { p10723(bufio.NewReader(os.Stdin), os.Stdout) }
