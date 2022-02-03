package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1626E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	color := make([]bool, n)
	for i := range color {
		Fscan(in, &color[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int8, n)
	cnt := make([]int, n)
	var f func(int, int)
	f = func(v, fa int) {
		if color[v] {
			ans[v] = 1
			cnt[v] = 1
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
				cnt[v] += cnt[w]
				if color[w] || ans[w] > 0 && cnt[w] > 1 {
					ans[v] = 1
				}
			}
		}
	}
	f(0, -1)

	f = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa { // 换根 DP，把 v 当成 w 的儿子
				if color[v] || ans[v] > 0 && cnt[0]-cnt[w] > 1 {
					ans[w] = 1
				}
				f(w, v)
			}
		}
	}
	f(0, -1)

	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1626E(os.Stdin, os.Stdout) }
