package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func p3806(_r io.Reader, out io.Writer) {
	var n, m, v, w, wt int
	in := bufio.NewReader(_r)
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	qs := make([]int, m)
	for i := range qs {
		Fscan(in, &qs[i])
	}

	markCentroid := make([]bool, n)
	size := make([]int, n)
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, e := range g[v] {
			if w := e.to; w != fa && !markCentroid[w] {
				if minSizeW, ctW, faCtW := findCentroid(w, v, compSize); minSizeW < minSize {
					minSize, ct, faCt = minSizeW, ctW, faCtW
				}
				maxSubSize = max(maxSubSize, size[w])
				size[v] += size[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-size[v])
		if maxSubSize < minSize {
			minSize, ct, faCt = maxSubSize, v, fa
		}
		return
	}

	ans := make([]bool, m)
	tmp := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)
		markCentroid[ct] = true
		defer func() { markCentroid[ct] = false }()
		for _, e := range g[ct] {
			if w := e.to; !markCentroid[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}

		has := map[int]bool{0: true}
		for _, e := range g[ct] {
			w := e.to
			if markCentroid[w] {
				continue
			}
			t := tmp[:0]
			var f func(int, int, int)
			f = func(v, fa, d int) {
				for i, q := range qs {
					if !ans[i] && has[q-d] {
						ans[i] = true
					}
				}
				t = append(t, d)
				for _, e := range g[v] {
					if w := e.to; w != fa && !markCentroid[w] {
						f(w, v, d+e.wt)
					}
				}
			}
			f(w, ct, e.wt)
			for _, d := range t {
				has[d] = true
			}
		}
	}
	dfs(0, -1, n)
	for _, b := range ans {
		if b {
			Fprintln(out, "AYE")
		} else {
			Fprintln(out, "NAY")
		}
	}
}

//func main() { p3806(os.Stdin, os.Stdout) }
