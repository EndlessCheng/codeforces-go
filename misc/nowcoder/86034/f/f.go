package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	type pair struct{ l, i int }
	qs := make([][]pair, n+1)
	Fscan(in, &q)
	for i := 0; i < q; i++ {
		var l, r int
		Fscan(in, &l, &r)
		qs[r] = append(qs[r], pair{l, i})
	}

	ans := make([]int, q)
	var mx, mn []int
	for r, v := range s {
		for len(mx) > 0 && v >= s[mx[len(mx)-1]] {
			mx = mx[:len(mx)-1]
		}
		mx = append(mx, r)
		for len(mn) > 0 && v <= s[mn[len(mn)-1]] {
			mn = mn[:len(mn)-1]
		}
		mn = append(mn, r)
		for _, p := range qs[r] {
			ans[p.i] = s[mx[sort.SearchInts(mx, p.l-1)]] - s[mn[sort.SearchInts(mn, p.l-1)]]
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
