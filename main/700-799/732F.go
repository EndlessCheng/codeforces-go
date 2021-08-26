package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF732F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, v, w int
	Fscan(in, &n, &m)
	type nb struct{ to, eid int }
	g := make([][]nb, n)
	type edge struct{ v, w int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], nb{w, i})
		g[w] = append(g[w], nb{v, i})
		es[i] = edge{v, w}
	}

	isBridge := make([]bool, m)
	dfn := make([]int, n)
	ts := 0
	var f func(int, int) int
	f = func(v, fa int) int {
		ts++
		dfn[v] = ts
		lowV := ts
		for _, e := range g[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := f(w, v)
				if lowW > dfn[v] {
					isBridge[e.eid] = true
				}
				lowV = min(lowV, lowW)
			} else if w != fa {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for v, ts := range dfn {
		if ts == 0 {
			f(v, -1)
		}
	}

	bccIDs := make([]int, n)
	idCnt := 0
	vs := []int{}
	vis := make([]bool, m)
	var f2 func(int)
	f2 = func(v int) {
		bccIDs[v] = idCnt
		vs = append(vs, v)
		for _, e := range g[v] {
			if w, i := e.to, e.eid; !isBridge[i] {
				if !vis[i] { // 避免重复定向
					vis[i] = true
					es[i] = edge{v, w} // 定向：BCC 内
				}
				if bccIDs[w] == 0 {
					f2(w)
				}
			}
		}
	}
	rt, maxBCC := 0, 0
	for v, id := range bccIDs {
		if id == 0 {
			idCnt++
			vs = []int{}
			f2(v)
			if len(vs) > maxBCC {
				rt, maxBCC = idCnt, len(vs)
			}
		}
	}

	g2 := make([][]nb, idCnt)
	for i, b := range isBridge {
		if b {
			v, w := bccIDs[es[i].v]-1, bccIDs[es[i].w]-1
			g2[v] = append(g2[v], nb{w, i})
			g2[w] = append(g2[w], nb{v, i})
		}
	}
	var f3 func(v, fa int)
	f3 = func(v, fa int) {
		for _, e2 := range g2[v] {
			if w, i := e2.to, e2.eid; w != fa {
				if e := es[i]; bccIDs[e.w]-1 == w { // 方向和输入相反
					es[i] = edge{e.w, e.v} // 定向：缩点后的 bridge tree 上
				}
				f3(w, v)
			}
		}
	}
	f3(rt-1, -1)
	Fprintln(out, maxBCC)
	for _, e := range es {
		Fprintln(out, e.v+1, e.w+1)
	}
}

//func main() { CF732F(os.Stdin, os.Stdout) }
