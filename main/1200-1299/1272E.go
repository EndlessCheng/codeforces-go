package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1272E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := make([]int, n)
	g := make([][]int, n)
	q := [2][]int{}
	vis := make([]bool, n)
	for i, v := range a {
		ans[i] = -1
		p := v & 1
		w1, w2 := i-v, i+v
		if w1 >= 0 && a[w1]&1 != p || w2 < n && a[w2]&1 != p {
			q[p] = append(q[p], i)
			vis[i] = true
		}
		if w1 >= 0 && a[w1]&1 == p {
			g[w1] = append(g[w1], i)
		}
		if w2 < n && a[w2]&1 == p {
			g[w2] = append(g[w2], i)
		}
	}

	for _, q := range q {
		for d := 1; len(q) > 0; d++ {
			tmp := q
			q = nil
			for _, v := range tmp {
				ans[v] = d
				for _, w := range g[v] {
					if !vis[w] {
						vis[w] = true
						q = append(q, w)
					}
				}
			}
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1272E(os.Stdin, os.Stdout) }
