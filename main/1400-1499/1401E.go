package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick01 []int

func (t fenwick01) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t fenwick01) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += t[i]
	}
	return
}

func cf1401E(in io.Reader, out io.Writer) {
	const mx int = 1e6
	var n, m, x, l, r int
	Fscan(in, &n, &m)

	ans := 1
	g := [mx + 2][]int{}
	for range n {
		Fscan(in, &x, &l, &r)
		g[l] = append(g[l], x<<1|1)
		g[r+1] = append(g[r+1], x<<1)
		if l == 0 && r == mx {
			ans++
		}
	}
	type pair struct{ l, r int }
	qs := [mx][]pair{}
	for range m {
		Fscan(in, &x, &l, &r)
		qs[x] = append(qs[x], pair{l, r})
	}

	t := make(fenwick01, mx+1)
	for i, qs := range qs[:] {
		for _, j := range g[i] {
			t.update(j>>1, j&1*2-1)
		}
		for _, q := range qs {
			ans += t.pre(q.r) - t.pre(q.l-1)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1401E(bufio.NewReader(os.Stdin), os.Stdout) }
