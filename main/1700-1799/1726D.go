package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1726D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := bytes.Repeat([]byte{'0'}, m)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}
		es := make([]struct{ x, y int }, m)
		rem := map[int]bool{}
		id := 0
		for i := range es {
			Fscan(in, &v, &w)
			v--
			w--
			es[i].x, es[i].y = v, w
			if find(v) != find(w) {
				fa[fa[v]] = fa[w]
			} else {
				rem[v] = true
				rem[w] = true
				id = i
				ans[i] = '1'
			}
		}

		if len(rem) == 3 {
			for i := range fa {
				fa[i] = i
			}
			fa[es[id].x] = es[id].y
			for i, e := range es {
				if ans[i] == '0' {
					if find(e.x) == find(e.y) {
						ans[i] = '1'
						break
					}
					fa[fa[e.x]] = fa[e.y]
				}
			}
			ans[id] = '0'
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1726D(os.Stdin, os.Stdout) }
