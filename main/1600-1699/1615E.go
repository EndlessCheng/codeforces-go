package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1615E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ls := []int{}
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		maxL := 0
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			l := dfs(w, v)
			if maxL > 0 {
				ls = append(ls, min(maxL, l))
			}
			maxL = max(maxL, l)
		}
		return maxL + 1
	}
	ls = append(ls, dfs(0, -1))

	maxR := len(ls)
	if maxR <= k {
		// b = 0，最大化 (n-r)*r，其中 r∈[maxR,k]
		// 目标：让 r 靠近 n/2
		r := maxR // n/2 < maxR
		if n/2 >= maxR {
			r = min(k, n/2)
		}
		Fprint(out, (n-r)*r)
		return
	}

	slices.Sort(ls)
	b := n
	for _, v := range ls[maxR-k:] {
		b -= v
	}
	b = min(b, n/2)
	Fprint(out, (n-k-b)*(k-b))
}

//func main() { cf1615E(bufio.NewReader(os.Stdin), os.Stdout) }
