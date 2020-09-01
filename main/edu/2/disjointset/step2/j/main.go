package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	dis := make([]int, n)
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			ffx := f(fa[x])
			dis[x] += dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	merge := func(from, to int) bool {
		if ff, ft := f(from), f(to); ff != ft {
			dis[ff] = 1 + dis[to] - dis[from]
			fa[ff] = ft
			return true
		}
		return dis[from]&1 != dis[to]&1
	}
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w)
		if !merge(v-1, w-1) {
			Fprint(out, i)
			return
		}
	}
	Fprint(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
