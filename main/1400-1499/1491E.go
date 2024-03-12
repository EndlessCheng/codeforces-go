package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1491E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	Fscan(in, &n)
	f := []int{1, 1}
	for f[len(f)-1]+f[len(f)-2] <= n {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}
	if f[len(f)-1] != n {
		Fprint(out, "NO")
		return
	}

	type edge struct{ to, eid int }
	g := make([][]edge, n+1)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], edge{w, i})
		g[w] = append(g[w], edge{v, i})
	}

	vis := make([]bool, n-1)
	var check func(int, int) bool
	check = func(rt, tar int) bool {
		if tar == 1 {
			return true
		}
		t := f[sort.SearchInts(f, tar)-1]
		var x, y, z int
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			size := 1
			for _, e := range g[v] {
				w := e.to
				if w == fa || vis[e.eid] {
					continue
				}
				sz := dfs(w, v)
				if z > 0 {
					return 0
				}
				if sz == t || sz == tar-t {
					vis[e.eid] = true
					x, y, z = w, v, sz
				}
				size += sz
			}
			return size
		}
		dfs(rt, 0)
		return z > 0 && check(x, z) && check(y, tar-z)
	}
	if check(1, n) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { cf1491E(os.Stdin, os.Stdout) }
