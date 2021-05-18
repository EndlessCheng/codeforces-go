package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF337D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, d, v, w, ans int
	Fscan(in, &n, &m, &d)
	aff := make([]bool, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v)
		aff[v] = true
	}
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type pair struct{ fi, fv, se int }
	dis := make([]pair, n+1)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		fi, fv, se := int(-1e9), 0, int(-1e9)
		for _, w := range g[v] {
			if w != fa {
				if d := f(w, v) + 1; d > fi {
					fi, fv, se = d, w, fi
				} else if d > se {
					se = d
				}
			}
		}
		dis[v] = pair{fi, fv, se}
		if fi < 0 && aff[v] {
			return 0
		}
		return fi
	}
	f(1, 0)

	var f2 func(v, fa, dFa int)
	f2 = func(v, fa, dFa int) {
		if dFa > d {
			return
		}
		dv := dis[v]
		if dv.fi <= d {
			ans++
		}
		if aff[v] && dFa < 0 {
			dFa = 0
		}
		for _, w := range g[v] {
			if w != fa {
				if w == dv.fv {
					f2(w, v, max(dFa, dv.se)+1)
				} else {
					f2(w, v, max(dFa, dv.fi)+1)
				}
			}
		}
	}
	f2(1, 0, -1e9)
	Fprint(out, ans)
}

//func main() { CF337D(os.Stdin, os.Stdout) }
