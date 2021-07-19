package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF936B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, sz, st int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := range g {
		Fscan(in, &sz)
		g[i] = make([]int, sz)
		for j := range g[i] {
			Fscan(in, &g[i][j])
			g[i][j]--
		}
	}
	Fscan(in, &st)
	st--

	path := []int{}
	vis := make([][2]bool, n)
	var f func(int, int8) bool
	f = func(v int, step int8) bool {
		path = append(path, v)
		if step > 0 && len(g[v]) == 0 {
			return true
		}
		vis[v][step] = true
		for _, w := range g[v] {
			if !vis[w][step^1] && f(w, step^1) {
				return true
			}
		}
		path = path[:len(path)-1]
		return false
	}
	if f(st, 0) {
		Fprintln(out, "Win")
		for _, v := range path {
			Fprint(out, v+1, " ")
		}
		return
	}

	vis2 := make([]int8, n)
	var f2 func(int) bool
	f2 = func(v int) bool {
		vis2[v] = 1
		for _, w := range g[v] {
			if tp := vis2[w]; tp == 1 || tp == 0 && f2(w) {
				return true
			}
		}
		vis2[v] = 2
		return false
	}
	if f2(st) {
		Fprint(out, "Draw")
	} else {
		Fprint(out, "Lose")
	}
}

//func main() { CF936B(os.Stdin, os.Stdout) }
