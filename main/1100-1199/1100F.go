package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type xorBasis00 struct{ b, pos [20]int }

func (b *xorBasis00) insertRightMost(idx, v int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.pos[i] = idx
			return
		}
		if idx >= b.pos[i] {
			idx, b.pos[i] = b.pos[i], idx
			v, b.b[i] = b.b[i], v
		}
		v ^= b.b[i]
	}
}

func (b *xorBasis00) maxXor(l int) (xor int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if xor>>i&1 == 0 && b.pos[i] >= l && xor^b.b[i] > xor {
			xor ^= b.b[i]
		}
	}
	return
}

func cf1100F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ l, i int }
	qs := make([][]pair, n)
	Fscan(in, &q)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &r)
		qs[r-1] = append(qs[r-1], pair{l - 1, i})
	}

	ans := make([]int, q)
	b := &xorBasis00{}
	for r, v := range a {
		b.insertRightMost(r, v)
		for _, p := range qs[r] {
			ans[p.i] = b.maxXor(p.l)
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1100F(os.Stdin, os.Stdout) }
