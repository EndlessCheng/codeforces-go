package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1528D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e9 + 1e3
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v][w] = wt
	}
	for st, row := range g {
		dis := append([]int{}, row...)
		vis := make([]bool, n)
		for {
			v := -1
			for w, b := range vis {
				if !b && (v < 0 || dis[w] < dis[v]) {
					v = w
				}
			}
			if v < 0 {
				break
			}
			vis[v] = true
			d := dis[v]
			dis[(v+1)%n] = min(dis[(v+1)%n], d+1)
			for j, wt := range g[v] {
				dis[(j+d)%n] = min(dis[(j+d)%n], d+wt)
			}
		}
		for i, d := range dis {
			if i == st {
				d = 0
			}
			Fprint(out, d, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1528D(os.Stdin, os.Stdout) }
