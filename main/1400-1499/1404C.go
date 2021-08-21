package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1404C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] = i - a[i]
	}
	qs := make([]struct{ l, r, i int }, q)
	for i := range qs {
		Fscan(in, &qs[i].l, &qs[i].r)
		qs[i].l++
		qs[i].r = n - qs[i].r
		qs[i].i = i
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].r < qs[j].r })

	ans := make([]int, q)
	tree := make([]int, n+2)
	kth := func(k int) (res int) {
		for b := 1 << 18; b > 0; b >>= 1 {
			if next := res | b; next <= n && k > tree[next] {
				k -= tree[next]
				res = next
			}
		}
		return res + 1
	}
	cur := 1
	for _, q := range qs {
		for ; cur <= q.r; cur++ {
			k := 1
			if a[cur] >= 0 {
				k = min(kth(cur-a[cur]), cur+1)
			}
			for ; k <= n; k += k & -k {
				tree[k]++
			}
		}
		s := 0
		for i := q.l; i > 0; i &= i - 1 {
			s += tree[i]
		}
		ans[q.i] = q.r - s
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF1404C(os.Stdin, os.Stdout) }
