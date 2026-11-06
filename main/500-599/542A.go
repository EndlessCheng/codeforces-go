package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf542A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, L, R, K int
	Fscan(in, &n, &m)
	type tuple struct{ l, r, i int }
	a := make([]tuple, n+1)
	b := make([]tuple, n+1)
	l := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].i = i
	}
	sort.Slice(a[1:], func(i, j int) bool { return a[i+1].l < a[j+1].l })

	t := 1
	l[1] = a[1].l
	b[1] = a[1]
	for i := 2; i <= n; i++ {
		if a[i].r > b[t].r {
			t++
			b[t] = a[i]
			l[t] = a[i].l
		}
	}

	var ans, p, q int
	for j := 1; j <= m; j++ {
		Fscan(in, &L, &R, &K)
		for k := max(sort.SearchInts(l[1:t+1], L), 1); k <= t && l[k] <= R; k++ {
			v := K * (min(R, b[k].r) - max(L, b[k].l))
			if v > ans {
				ans = v
				p = b[k].i
				q = j
			}
		}
	}

	Fprintln(out, ans)
	if ans != 0 {
		Fprintln(out, p, q)
	}
}

//func main() { cf542A(bufio.NewReader(os.Stdin), os.Stdout) }
