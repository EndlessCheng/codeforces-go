package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, ans, cntV, cntE int
	Fscan(in, &n)
	g := [4e5 + 1][]int{}
	for ; n > 0; n-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make([]bool, len(g))
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		cntV++
		cntE += len(g[v])
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}
	for i, b := range vis {
		if !b && g[i] != nil {
			cntV, cntE = 0, 0
			dfs(i)
			ans += min(cntV, cntE/2)
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if a > b { return b }; return a }
