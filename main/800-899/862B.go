package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf862B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	cnt := [2]int{}
	var dfs func(int, int, int)
	dfs = func(v, fa, d int) {
		cnt[d]++
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d^1)
			}
		}
	}
	dfs(0, -1, 0)
	Fprint(out, cnt[0]*cnt[1]-n+1)
}

//func main() { cf862B(os.Stdin, os.Stdout) }
