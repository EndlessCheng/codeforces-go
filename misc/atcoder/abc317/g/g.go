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

	var n, m, v int
	Fscan(in, &n, &m)
	g := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]int, m)
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			g[i][j] = v - 1
		}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	matchL := make([]int, n)
	matchR := make([]int, n)
	for k := 0; k < m; k++ {
		for i := range matchL {
			matchL[i] = -1
			matchR[i] = -1
		}
		vis := make([]int, n)
		for ts := range g {
			var f func(int) bool
			f = func(v int) bool {
				vis[v] = ts + 1
				for id, w := range g[v] {
					lv := matchR[w]
					if lv == -1 || vis[lv] != ts+1 && f(lv) {
						matchR[w] = v
						matchL[v] = id
						return true
					}
				}
				return false
			}
			f(ts)
		}
		for i, id := range matchL {
			ans[i][k] = g[i][id]
			delete(g[i], id)
		}
	}
	Fprintln(out, "Yes")
	for _, v := range ans {
		for _, v := range v {
			Fprint(out, v+1, " ")
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
