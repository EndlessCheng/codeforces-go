package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF629D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct {
		v int64
		i int
	}

	var n int
	var r, h, ans int64
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &r, &h)
		a[i] = pair{r * r * h, i}
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.v < b.v || a.v == b.v && a.i > b.i })

	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	tree := make([]int64, n+1)
	add := func(i int, val int64) {
		for ; i <= n; i += i & -i {
			tree[i] = max(tree[i], val)
		}
	}
	query := func(i int) (res int64) {
		for ; i > 0; i &= i - 1 {
			res = max(res, tree[i])
		}
		return
	}
	for _, p := range a {
		v := p.v + query(p.i)
		ans = max(ans, v)
		add(p.i+1, v)
	}
	Fprintf(_w, "%.9f", float64(ans)*math.Pi)
}

//func main() { CF629D(os.Stdin, os.Stdout) }
