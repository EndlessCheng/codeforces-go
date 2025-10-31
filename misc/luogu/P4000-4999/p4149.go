package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func p4149(_r io.Reader, out io.Writer) {
	var n, k, v, w, wt int
	in := bufio.NewReader(_r)
	Fscan(in, &n, &k)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	markCentroid := make([]bool, len(g))
	size := make([]int, len(g))
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

	ans := int(1e9)
	type pair struct{ s, c int }
	tmp := make([]pair, n)
	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)

		minE := map[int]int{0: 0}
		for _, e := range g[ct] {
			w := e.to
			if markCentroid[w] {
				continue
			}
			t := tmp[:0]
			var f func(int, int, int, int)
			f = func(v, fa, s, c int) {
				if mn, ok := minE[k-s]; ok {
					ans = min(ans, mn+c)
				}
				t = append(t, pair{s, c})
				if s > k {
					return
				}
				for _, e := range g[v] {
					if w := e.to; w != fa && !markCentroid[w] {
						f(w, v, s+e.wt, c+1)
					}
				}
			}
			f(w, ct, e.wt, 1)
			for _, p := range t {
				if mn, ok := minE[p.s]; ok {
					if p.c < mn {
						minE[p.s] = p.c
					}
				} else {
					minE[p.s] = p.c
				}
			}
		}

		markCentroid[ct] = true
		for _, e := range g[ct] {
			if w := e.to; !markCentroid[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}
	}
	dfs(0, -1, len(g))
	if ans == 1e9 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { p4149(os.Stdin, os.Stdout) }
