package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF977E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, ans, c int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n+1)
	var f func(int)
	f = func(v int) {
		vis[v] = true
		if len(g[v]) != 2 {
			c = 0
		}
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			c = 1
			f(i)
			ans += c
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF977E(os.Stdin, os.Stdout)
//}
