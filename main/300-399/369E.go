package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF369E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e6
	type seg struct{ l, r, id int }

	var n, q, m, p int
	Fscan(in, &n, &q)
	a := make([]seg, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].id = -1
	}
	for i := 0; i < q; i++ {
		pre := 0
		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &p)
			if pre+1 < p {
				a = append(a, seg{pre + 1, p - 1, i})
			}
			pre = p
		}
		if pre < mx {
			a = append(a, seg{pre + 1, mx, i})
		}
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.l > b.l || a.l == b.l && (a.r < b.r || a.r == b.r && a.id < b.id) })

	ans := make([]int, q)
	tree := [mx + 1]int{}
	for _, p := range a {
		if p.id < 0 {
			for i := p.r; i <= mx; i += i & -i {
				tree[i]++
			}
		} else {
			for i := p.r; i > 0; i &= i - 1 {
				ans[p.id] += tree[i]
			}
		}
	}
	for _, v := range ans {
		Fprintln(out, n-v)
	}
}

//func main() { CF369E(os.Stdin, os.Stdout) }
