package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick92 []int

func (t fenwick92) update(i, x int) {
	for ; i < len(t); i += i & -i {
		t[i] += x
	}
}

func (t fenwick92) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += t[i]
	}
	return
}

func cf992E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, p, x int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	t := make(fenwick92, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		t.update(i, a[i])
	}

	var dfs func(int, int) int
	dfs = func(l, r int) int {
		if l+1 == r {
			if t.pre(l)*2 == t.pre(r) {
				return r
			}
			return -1
		}
		res := -1
		m := (l + r) / 2
		if t.pre(l)*2 <= t.pre(m) {
			res = max(res, dfs(l, m))
		}
		if t.pre(m)*2 <= t.pre(r) {
			res = max(res, dfs(m, r))
		}
		return res
	}

	for range q {
		Fscan(in, &p, &x)
		t.update(p, x-a[p])
		a[p] = x
		Fprintln(out, dfs(0, n))
	}
}

//func main() { cf992E(bufio.NewReader(os.Stdin), os.Stdout) }
