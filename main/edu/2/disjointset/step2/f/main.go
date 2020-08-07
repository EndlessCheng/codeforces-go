package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	fa := make([]int, n)
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}

	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		es[i] = edge{v - 1, w - 1, wt}
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })

	ans := math.MaxInt32
o:
	for i, ei := range es {
		min, c := ei.wt, 0
		for j := range fa {
			fa[j] = j
		}
		for _, e := range es[i:] {
			if fv, fw := f(e.v), f(e.w); fv != fw {
				fa[fv] = fw
				c++
				if c == n-1 {
					if e.wt-min < ans {
						ans = e.wt - min
					}
					continue o
				}
			}
		}
	}
	if ans < math.MaxInt32 {
		Fprint(out, "YES\n", ans)
	} else {
		Fprint(out, "NO")
	}
}

func main() { run(os.Stdin, os.Stdout) }
