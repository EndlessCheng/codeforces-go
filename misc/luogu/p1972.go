package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p1972(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			in.Read(buf)
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

	n := r()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = r()
	}
	q := r()
	type pair struct{ l, r, i int }
	qs := make([]pair, q)
	for i := range qs {
		qs[i] = pair{r(), r(), i}
	}
	slices.SortFunc(qs, func(a, b pair) int { return a.r - b.r })

	tree := make([]int, n+1)
	add := func(i int, val int) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	query := func(l, r int) int { return sum(r) - sum(l-1) }

	ans := make([]int, q)
	posR := [1e6 + 1]int{}
	i := 1
	for _, q := range qs {
		for ; i <= q.r; i++ {
			if p := posR[a[i]]; p > 0 {
				add(p, -1)
			}
			add(i, 1)
			posR[a[i]] = i
		}
		ans[q.i] = query(q.l, q.r)
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p1972(bufio.NewReader(os.Stdin), os.Stdout) }
