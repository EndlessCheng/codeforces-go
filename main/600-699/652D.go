package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF652D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ps := make([]struct{ v, i int }, n*2)
	for i := range ps {
		Fscan(in, &ps[i].v)
		ps[i].i = i
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })
	kth := make([]int, n*2)
	for i, p := range ps {
		kth[p.i] = i + 1
	}

	type pair struct{ l, r, i int }
	a := make([]pair, n)
	for i := range a {
		a[i] = pair{kth[i*2], kth[i*2+1], i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	ans := make([]int, n)
	tree := make([]int, n*2+1)
	for i, p := range a {
		c := i
		for i := p.l; i > 0; i &= i - 1 {
			c -= tree[i]
		}
		ans[p.i] = c
		for i := p.l; i <= n*2; i += i & -i {
			tree[i]++
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF652D(os.Stdin, os.Stdout) }
