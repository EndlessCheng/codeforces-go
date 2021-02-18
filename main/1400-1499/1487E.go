package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1487E(_r io.Reader, out io.Writer) {
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
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
	const inf int = 1e9
	type pair struct{ v, i int }

	n := [4]int{}
	for i := range n {
		n[i] = r()
	}
	a := [4][]pair{}
	for i := range a {
		a[i] = make([]pair, n[i])
		for j := range a[i] {
			a[i][j] = pair{r(), j}
		}
	}
	vis := make([]bool, 15e4)
	for i := 1; i < 4; i++ {
		pre, cur := a[i-1], a[i]
		sort.Slice(pre, func(i, j int) bool { return pre[i].v < pre[j].v })
		b := make([][]int, len(cur))
		for m := r(); m > 0; m-- {
			v, w := r()-1, r()-1
			b[w] = append(b[w], v)
		}
		for j, ps := range b {
			for _, p := range ps {
				vis[p] = true
			}
			mi := inf
			for _, p := range pre {
				if !vis[p.i] {
					mi = p.v
					break
				}
			}
			cur[j].v += mi
			for _, p := range ps {
				vis[p] = false
			}
		}
	}
	ans := inf
	for _, p := range a[3] {
		if p.v < ans {
			ans = p.v
		}
	}
	if ans == inf {
		ans = -1
	}
	Fprintln(out, ans)
}

//func main() { CF1487E(os.Stdin, os.Stdout) }
