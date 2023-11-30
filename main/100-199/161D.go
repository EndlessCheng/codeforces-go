package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF161D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, w ,ans int
	Fscan(in, &n, &k)
	k = min(k, n-1)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	markCentroid := make([]bool, len(g))
	size := make([]int, len(g))
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, w := range g[v] {
			if w != fa && !markCentroid[w] {
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

	_cnt := make([]int, n)
	_cnt[0] = 1
	tmp := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)
		markCentroid[ct] = true
		defer func() { markCentroid[ct] = false }()
		for _, w := range g[ct] {
			if !markCentroid[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}

		cnt := _cnt[:1]
		for _, w := range g[ct] {
			if markCentroid[w] {
				continue
			}
			t := tmp[:1]
			var f func(int, int, int)
			f = func(v, fa, d int) {
				if k-d < len(cnt) {
					ans += cnt[k-d]
				}
				if d == len(t) {
					t = append(t, 1)
				} else {
					t[d]++
				}
				if d >= k {
					return
				}
				for _, w := range g[v] {
					if w != fa && !markCentroid[w] {
						f(w, v, d+1)
					}
				}
			}
			f(w, ct, 1)
			for i, c := range t {
				if i >= len(cnt) {
					cnt = append(cnt, c)
				} else {
					cnt[i] += c
				}
			}
		}
	}
	dfs(0, -1, len(g))
	Fprint(out, ans)
}

//func main() { CF161D(os.Stdin, os.Stdout) }
