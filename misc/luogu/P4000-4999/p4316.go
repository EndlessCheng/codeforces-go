package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p4316(in io.Reader, out io.Writer) {
	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v-1] = append(g[v-1], nb{w - 1, wt})
	}

	memo := make([]float64, n)
	var f func(int) float64
	f = func(v int) (res float64) {
		if g[v] == nil {
			return 0
		}
		p := &memo[v]
		if *p != 0 {
			return *p
		}
		for _, e := range g[v] {
			res += f(e.to) + float64(e.wt)
		}
		res /= float64(len(g[v]))
		*p = res
		return
	}
	Fprintf(out, "%.2f", f(0))
}

//func main() { p4316(bufio.NewReader(os.Stdin), os.Stdout) }
