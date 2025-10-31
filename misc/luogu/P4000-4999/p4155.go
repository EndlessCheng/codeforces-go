package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p4155(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	if n == 1 {
		Fprint(out, 0)
		return
	}
	type tuple struct{ l, r, i int }
	a := make([]tuple, n, n*2)
	for i := range a {
		a[i].i = i
		Fscan(in, &a[i].l, &a[i].r)
		if a[i].r < a[i].l {
			a[i].r += m
		}
	}

	slices.SortFunc(a, func(a, b tuple) int { return a.l - b.l })
	for _, t := range a {
		t.l += m
		t.r += m
		a = append(a, t)
	}

	const mx = 19
	pa := make([][mx]int, n*2)
	r := n*2 - 1
	for i := n*2 - 1; i >= 0; i-- {
		for a[r].l > a[i].r {
			r--
		}
		pa[i][0] = r
	}

	for i := 0; i < mx-1; i++ {
		for x := range pa {
			p := pa[x][i]
			pa[x][i+1] = pa[p][i]
		}
	}

	ans := make([]any, n)
	for i, p := range a[:n] {
		res := 0
		cur := i
		for k := mx - 1; k >= 0; k-- {
			x := pa[cur][k]
			if a[x].r < p.l+m {
				cur = x
				res |= 1 << k
			}
		}
		ans[p.i] = res + 2
	}
	Fprintln(out, ans...)
}

//func main() { p4155(bufio.NewReader(os.Stdin), os.Stdout) }
