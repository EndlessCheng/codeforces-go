package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF891C(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m := r(), r()
	fa := make([]int, n+1)
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
	type edge struct{ v, w, wt, fv, fw int }
	es := make([]edge, m)
	ids := make([][]int, 5e5+1)
	for i := range es {
		es[i] = edge{r(), r(), r(), 0, 0}
		ids[es[i].wt] = append(ids[es[i].wt], i)
	}
	for _, ids := range ids {
		for _, id := range ids {
			// 记录下每条边两个端点所在的联通块
			es[id].fv = find(es[id].v)
			es[id].fw = find(es[id].w)
		}
		for _, id := range ids {
			fa[find(es[id].v)] = find(es[id].w)
		}
	}

	faMap := map[int]int{}
	find = func(x int) int {
		if fx, ok := faMap[x]; ok && fx != x {
			faMap[x] = find(fx)
			return faMap[x]
		}
		return x
	}
o:
	for q := r(); q > 0; q-- {
		faMap = map[int]int{}
		for k := r(); k > 0; k-- {
			i := r() - 1
			// 把询问边变成询问连通分量
			v, w := find(es[i].fv), find(es[i].fw)
			if v == w {
				Fprintln(out, "NO")
				for k--; k > 0; k-- {
					r()
				}
				continue o
			}
			faMap[v] = w
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF891C(os.Stdin, os.Stdout) }
