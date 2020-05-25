package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF219D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([]map[int]bool, n)
	for i := range g {
		g[i] = map[int]bool{}
	}
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v-1][w-1] = true
		g[w-1][v-1] = false
	}
	inv := make([]int, n)
	var f func(v, fa int)
	f = func(v, fa int) {
		for w, b := range g[v] {
			if w != fa {
				if !b {
					inv[0]++
				}
				f(w, v)
			}
		}
	}
	f(0, -1)
	f = func(v, fa int) {
		if g[v][fa] {
			inv[v] = inv[fa] - 1
		} else {
			inv[v] = inv[fa] + 1
		}
		for w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
	}
	for v := range g[0] {
		f(v, 0)
	}
	min := n
	for _, v := range inv {
		if v < min {
			min = v
		}
	}
	Fprintln(out, min)
	for i, v := range inv {
		if v == min {
			Fprint(out, i+1, " ")
		}
	}
}

//func main() { CF219D(os.Stdin, os.Stdout) }
