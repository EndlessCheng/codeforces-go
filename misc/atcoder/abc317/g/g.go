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

	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			a[i][j]--
		}
	}

	g := make([]map[int]int, n)
	for i, r := range a {
		g[i] = map[int]int{}
		for j, v := range r {
			g[i][j] = v
		}
	}
	
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for k := 0; k < m; k++ {
		matchL := make([]int, n)
		matchR := make([]int, n)
		for i := range matchL {
			matchL[i] = -1
			matchR[i] = -1
		}
		var used []bool
		var f func(int) bool
		f = func(v int) bool {
			used[v] = true
			for id, w := range g[v] {
				if lv := matchR[w]; lv == -1 || !used[lv] && f(lv) {
					matchR[w] = v
					matchL[v] = id
					return true
				}
			}
			return false
		}
		for v := range g {
			used = make([]bool, n)
			f(v)
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
