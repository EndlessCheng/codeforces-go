package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func p2868(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, t int
	Fscan(in, &n, &m)
	fun := make([]float64, n)
	for i := range fun {
		Fscan(in, &fun[i])
	}
	type nb struct{ to, t int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &t)
		g[v-1] = append(g[v-1], nb{w - 1, t})
	}

	f := func(avg float64) bool {
		dist := make([]float64, n)
		q := make([]int, n)
		inQ := make([]bool, n)
		for i := range q {
			q[i] = i
			inQ[i] = true
		}
		cnt := make([]int, n)
		for len(q) > 0 {
			v, q = q[0], q[1:]
			inQ[v] = false
			for _, e := range g[v] {
				w := e.to
				if newD := dist[v] + avg*float64(e.t) - fun[v]; newD < dist[w] { // 写成 fun[w] 也行，反正是个环
					dist[w] = newD
					cnt[w] = cnt[v] + 1
					if cnt[w] >= n {
						return true
					}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		return false
	}
	l, r := 0.0, 1e3
	for step := int(math.Log2((r - l) / 1e-4)); step > 0; step-- {
		mid := (l + r) / 2
		if f(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	Fprintf(out, "%.2f", (l+r)/2)
}

//func main() { p2868(os.Stdin, os.Stdout) }
