package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	g := make([][]int, 4*n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		g[v] = append(g[v], i*2, i*2+1)
	}
	ans := make([]int, 2*n+2)
	var dfs func(v, d int)
	dfs = func(v, d int) {
		ans[v] = d
		for _, w := range g[v] {
			dfs(w, d+1)
		}
	}
	dfs(1, 0)
	for _, v := range ans[1:] {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
