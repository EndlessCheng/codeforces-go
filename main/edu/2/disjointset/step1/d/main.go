package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type query struct {
		ask, same bool
		v, w      int
	}

	var n, m, q, v int
	var s string
	Fscan(in, &n, &m, &q)
	for ; m > 0; m-- {
		Fscan(in, &v, &v)
	}
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	qs := make([]query, q)
	for i := range qs {
		Fscan(in, &s, &qs[i].v, &qs[i].w)
		qs[i].ask = s[0] == 'a'
	}
	for i := len(qs) - 1; i >= 0; i-- {
		q := qs[i]
		if q.ask {
			qs[i].same = f(q.v) == f(q.w)
		} else {
			fa[f(q.v)] = fa[f(q.w)]
		}
	}
	for _, q := range qs {
		if q.ask {
			if q.same {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
