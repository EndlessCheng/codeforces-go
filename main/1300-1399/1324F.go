package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1324F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	color := make([]int, n)
	for i := range color {
		Fscan(in, &color[i])
		if color[i] == 0 {
			color[i] = -1
		}
	}
	g := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	g[1] = append(g[1], 0)

	ans := make([]int, n+1)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		sum := color[v-1]
		for _, w := range g[v] {
			if w != fa {
				sum += f(w, v)
			}
		}
		ans[v] = sum
		if sum < 0 {
			return 0
		}
		return sum
	}
	f(1, 0)
	var f2 func(v, fa int)
	f2 = func(v, fa int) {
		if ans[v] >= 0 {
			ans[v] = max(ans[v], ans[fa])
		} else {
			ans[v] = max(ans[v], ans[v]+ans[fa])
		}
		for _, w := range g[v] {
			if w != fa {
				f2(w, v)
			}
		}
	}
	f2(1, 0)
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF1324F(os.Stdin, os.Stdout) }
