package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type node24 struct {
	lr       [2]*node24
	priority uint
	v, c1    int
}

func (o *node24) cmp(b int) int {
	switch {
	case b < o.v:
		return 0
	case b > o.v:
		return 1
	default:
		return -1
	}
}

func (o *node24) rotate(d int) *node24 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap24 struct {
	rd   uint
	root *node24
}

func (t *treap24) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap24) _put(o *node24, v, c1 int) *node24 {
	if o == nil {
		return &node24{priority: t.fastRand(), v: v, c1: c1}
	}
	if d := o.cmp(v); d >= 0 {
		o.lr[d] = t._put(o.lr[d], v, c1)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap24) put(v, c1 int) { t.root = t._put(t.root, v, c1) }

func (t *treap24) prev(v int) (prev *node24) {
	for o := t.root; o != nil; {
		if o.cmp(v) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

func (t *treap24) lowerBound(v int) (lb *node24) {
	for o := t.root; o != nil; {
		switch c := o.cmp(v); {
		case c == 0:
			lb = o
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func CF424D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, tar, p, u, d, R1, C1, R2, C2 int
	Fscan(in, &n, &m, &tar, &p, &u, &d)
	f := func(from, to int) int {
		if from < to {
			return u
		}
		if from > to {
			return d
		}
		return p
	}
	a := make([][]int, n)
	lr := make([][]int, n)
	rl := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		lr[i] = make([]int, m)
		rl[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			if j > 0 {
				lr[i][j] = lr[i][j-1] + f(a[i][j-1], a[i][j])
				rl[i][j] = rl[i][j-1] + f(a[i][j], a[i][j-1])
			}
		}
	}
	ud := make([][]int, m)
	du := make([][]int, m)
	for j := 0; j < m; j++ {
		ud[j] = make([]int, n)
		du[j] = make([]int, n)
		for i := 1; i < n; i++ {
			ud[j][i] = ud[j][i-1] + f(a[i-1][j], a[i][j])
			du[j][i] = du[j][i-1] + f(a[i][j], a[i-1][j])
		}
	}

	minD := int(1e9)
	for r1, lr1 := range lr {
		for r2 := r1 + 2; r2 < n; r2++ {
			t := &treap24{rd: 1}
			for c2 := 2; c2 < m; c2++ {
				c1 := c2 - 2
				t.put(lr1[c1]+rl[r2][c1]-du[c1][r2]+du[c1][r1], c1)
				v := lr1[c2] + rl[r2][c2] + ud[c2][r2] - ud[c2][r1] - tar
				if o := t.lowerBound(v); o != nil {
					if d := o.v - v; d < minD {
						minD, R1, C1, R2, C2 = d, r1, o.c1, r2, c2
					}
				}
				if o := t.prev(v); o != nil {
					if d := v - o.v; d < minD {
						minD, R1, C1, R2, C2 = d, r1, o.c1, r2, c2
					}
				}
			}
		}
	}
	Fprint(out, R1+1, C1+1, R2+1, C2+1)
}

//func main() { CF424D(os.Stdin, os.Stdout) }
