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
	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	cnt := [2]int{}
	var dfs func(v, fa, d int)
	dfs = func(v, fa, d int) {
		cnt[d]++
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d^1)
			}
		}
	}
	dfs(1, 0, 0)
	Fprint(out, cnt[0]*cnt[1]-n+1)
}

func main() { run(os.Stdin, os.Stdout) }
