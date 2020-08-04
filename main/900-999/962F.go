package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF962F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, v, w int
	Fscan(in, &n, &m)
	type nei struct{ to, eid int }
	g := make([][]nei, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], nei{w, i})
		g[w] = append(g[w], nei{v, i})
	}

	onSimpleCycle := make([]bool, m+1)
	bccIDs := make([]int, n)
	idCnt := 0
	dfn := make([]int, n)
	dfsClock := 0
	type edge struct{ v, w, eid int }
	stack := []edge{}
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		for _, ne := range g[v] {
			w := ne.to
			e := edge{v, w, ne.eid}
			if dfn[w] == 0 {
				stack = append(stack, e)
				lowW := f(w, v)
				if lowW >= dfn[v] {
					idCnt++
					var comp, eids []int
					for {
						e, stack = stack[len(stack)-1], stack[:len(stack)-1]
						if bccIDs[e.v] != idCnt {
							bccIDs[e.v] = idCnt
							comp = append(comp, e.v)
						}
						if bccIDs[e.w] != idCnt {
							bccIDs[e.w] = idCnt
							comp = append(comp, e.w)
						}
						eids = append(eids, e.eid)
						if e.v == v && e.w == w {
							break
						}
					}
					if len(comp) == len(eids) {
						for _, eid := range eids {
							onSimpleCycle[eid] = true
						}
					}
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] {
				stack = append(stack, e)
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for v, ts := range dfn {
		if ts == 0 && len(g[v]) > 0 {
			f(v, -1)
		}
	}

	ans := []interface{}{}
	for i, b := range onSimpleCycle {
		if b {
			ans = append(ans, i)
		}
	}
	Fprintln(out, len(ans))
	Fprint(out, ans...)
}

//func main() { CF962F(os.Stdin, os.Stdout) }
