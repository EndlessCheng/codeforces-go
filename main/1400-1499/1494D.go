package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1494D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ i, j, v int }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, v int
	Fscan(in, &n)
	fa := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	a := make([]pair, 0, n*(n-1)/2)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &v)
			if i < j {
				a = append(a, pair{i, j, v})
			} else if i == j {
				ans[i] = v
				fa[i] = i
			}
		}
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.v < b.v || a.v == b.v && a.i < b.i })
	pa := make([]int, n)
	for _, p := range a {
		if i, j := find(p.i), find(p.j); i == j {
		} else if max(ans[i], ans[j]) < p.v {
			ans = append(ans, p.v)
			f := len(fa)
			fa = append(fa, f)
			fa[f] = f
			fa[i] = f
			fa[j] = f
			pa = append(pa, f)
			pa[i] = f
			pa[j] = f
		} else if ans[i] == p.v {
			fa[j] = i
			pa[j] = i
		} else {
			fa[i] = j
			pa[i] = j
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	for i, f := range fa {
		if i == f {
			Fprintln(out, i+1)
			break
		}
	}
	for i, p := range pa {
		if i != p {
			Fprintln(out, i+1, p+1)
		}
	}
}

//func main() { CF1494D(os.Stdin, os.Stdout) }
