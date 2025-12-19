package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
var sz93 int
var b93 []int

type fenwick93 [][]int32

func newFenwick93(n, m int) fenwick93 {
	t := make(fenwick93, (n-1)/sz93+2)
	for i := range t {
		t[i] = make([]int32, m+1)
	}
	return t
}

func (t fenwick93) update(x, y int, val int32) {
	for i := x/sz93 + 1; i < len(t); i += i & -i {
		for j := y + 1; j < len(t[i]); j += j & -j {
			t[i][j] += val
		}
	}
}

func (t fenwick93) pre(x, y int) (res int32) {
	if x < 0 || y < 0 {
		return
	}
	for _, v := range b93[x-x%sz93 : x+1] {
		if v <= y {
			res++
		}
	}
	for i := x / sz93; i > 0; i &= i - 1 {
		for j := y + 1; j > 0; j &= j - 1 {
			res += t[i][j]
		}
	}
	return
}

func (t fenwick93) query(l1, r1, l2, r2 int) int32 {
	return t.pre(r1, r2) - t.pre(r1, l2-1) - t.pre(l1-1, r2) + t.pre(l1-1, l2-1)
}

func cf1093E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, x, y, l, r int
	Fscan(in, &n, &m)
	w := bits.Len(uint(n))
	sz93 = w * w * 3
	pos := make([]int, n+1)
	for i := range n {
		Fscan(in, &x)
		pos[x] = i
	}
	t := newFenwick93(n, n)
	b93 = make([]int, n)
	for i := range b93 {
		Fscan(in, &b93[i])
		b93[i] = pos[b93[i]]
		t.update(i, b93[i], 1)
	}

	for range m {
		Fscan(in, &op, &x, &y)
		x--
		y--
		if op == 1 {
			Fscan(in, &l, &r)
			l--
			r--
			Fprintln(out, t.query(l, r, x, y))
		} else {
			t.update(x, b93[x], -1)
			t.update(x, b93[y], 1)
			t.update(y, b93[y], -1)
			t.update(y, b93[x], 1)
			b93[x], b93[y] = b93[y], b93[x]
		}
	}
}

//func main() { cf1093E(bufio.NewReader(os.Stdin), os.Stdout) }
