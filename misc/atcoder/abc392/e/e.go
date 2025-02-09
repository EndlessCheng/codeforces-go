package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	type pair struct{ i, v int }
	ps := []pair{}
	for i := 1; i <= m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		fv, fw := find(v), find(w)
		if fv != fw {
			fa[fv] = fw
		} else {
			ps = append(ps, pair{i, v})
		}
	}

	roots := []int{}
	for i := 1; i <= n; i++ {
		if find(i) == i {
			roots = append(roots, i)
		}
	}
	Fprintln(out, len(roots)-1)

	ex := map[int][]pair{}
	for _, p := range ps {
		rt := fa[p.v]
		ex[rt] = append(ex[rt], p)
	}
	exRt := make([]int, 0, len(ex))
	for k := range ex {
		exRt = append(exRt, k)
	}
	sort.Ints(exRt)

	j := 0
	for _, rt := range exRt {
		for _, p := range ex[rt] {
			if roots[j] == rt {
				j++
			}
			if j == len(roots) {
				return
			}
			Fprintln(out, p.i, p.v, roots[j])
			j++
			if j == len(roots) {
				return
			}
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
