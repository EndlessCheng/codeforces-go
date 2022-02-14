package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1361C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][2]int, n)
	for i := range a {
		Fscan(in, &a[i][0], &a[i][1])
	}
	var ans []interface{}
	path := make([]interface{}, 0, n*2)
	Fprintln(out, sort.Search(21, func(b int) bool {
		mask := 1<<b - 1
		type nb struct{ to, eid int }
		g := make([][]nb, mask+1)
		for i, p := range a {
			v, w := p[0]&mask, p[1]&mask
			g[v] = append(g[v], nb{w, i})
			g[w] = append(g[w], nb{v, i})
		}
		for _, vs := range g {
			if len(vs)&1 > 0 {
				return true
			}
		}
		path = path[:0]
		vis := make([]bool, n)
		var f func(int)
		f = func(v int) {
			for len(g[v]) > 0 {
				e := g[v][0]
				g[v] = g[v][1:]
				i := e.eid
				if vis[i] {
					continue
				}
				vis[i] = true
				f(e.to)
				if a[i][0]&mask == v {
					path = append(path, i*2+2, i*2+1)
				} else {
					path = append(path, i*2+1, i*2+2)
				}
			}
		}
		for i := range g {
			f(i)
			if len(path) > 0 && len(path) < n*2 {
				return true
			}
		}
		ans = append([]interface{}{}, path...)
		return false
	})-1)
	Fprintln(out, ans...)
}

//func main() { CF1361C(os.Stdin, os.Stdout) }
