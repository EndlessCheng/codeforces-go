package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go

// todo 下面这种写法为什么 TLE 了？？
// https://atcoder.jp/contests/abc199/submissions/22685300

func run(in io.Reader, out io.Writer) {
	var n, m, v, w, cnt int
	Fscan(in, &n, &m)
	color := make([]int, n)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	comp := []int{}
	vis := make([]bool, n)
	var dfs func(v int)
	dfs = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}

	var f func(int)
	f = func(p int) {
		if p == len(comp) {
			cnt++
			return
		}
		v := comp[p]
	o:
		for i := 1; i <= 3; i++ {
			for _, w := range g[v] {
				if color[w] == i {
					continue o
				}
			}
			color[v] = i
			f(p + 1)
			color[v] = 0
		}
	}

	ans := 1
	for i, b := range vis {
		if !b {
			comp = []int{}
			dfs(i)
			cnt = 0
			f(0)
			ans *= cnt
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
