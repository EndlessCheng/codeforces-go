package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod11 = 998244353

type matrix11 [][]int

func newMatrix11(n, m int) matrix11 {
	a := make(matrix11, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix11) mul(b matrix11) matrix11 {
	c := newMatrix11(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod11
			}
		}
	}
	return c
}

func (a matrix11) powMul(n int, f0 matrix11) matrix11 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf1511F(in io.Reader, out io.Writer) {
	var n, k int
	var s string
	Fscan(in, &n, &k)

	type node struct {
		son [26]*node
		end bool
	}
	root := &node{}
	for range n {
		Fscan(in, &s)
		o := root
		for _, b := range s {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{}
			}
			o = o.son[b]
		}
		o.end = true
	}

	m := newMatrix11(145, 145)
	type pair struct{ p, q *node }
	idx := map[pair]int{}
	var dfs func(*node, *node) int
	dfs = func(p, q *node) int {
		t := pair{p, q}
		if i, ok := idx[t]; ok {
			return i
		}
		v := len(idx)
		idx[t] = v
		for i := range 26 {
			a, b := p.son[i], q.son[i]
			if a == nil || b == nil {
				continue
			}
			m[dfs(a, b)][v] = 1
			if a.end {
				m[dfs(root, b)][v] = 1
			}
			if b.end {
				m[dfs(root, a)][v]++
			}
			if a.end && b.end {
				m[dfs(root, root)][v]++
			}
		}
		return v
	}
	dfs(root, root)

	size := len(idx)
	m = m[:size]
	for i, row := range m {
		m[i] = row[:size]
	}

	f0 := newMatrix11(size, 1)
	f0[0][0] = 1

	fn := m.powMul(k, f0)
	Fprint(out, fn[0][0]%mod11)
}

//func main() { cf1511F(os.Stdin, os.Stdout) }
