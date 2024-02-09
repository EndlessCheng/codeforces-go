package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
type xorBasis struct{ b, pos [60]int }

func (b *xorBasis) insertRightMost(idx, v int) {
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

func (b *xorBasis) decompose(l, v int) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 || b.pos[i] < l {
			return false
		}
		v ^= b.b[i]
	}
	return true
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, x int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ l, v, i int }
	qs := make([][]pair, n)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &r, &x)
		qs[r-1] = append(qs[r-1], pair{l - 1, x, i})
	}

	ans := make([]bool, q)
	b := &xorBasis{}
	for r, v := range a {
		b.insertRightMost(r, v)
		for _, p := range qs[r] {
			ans[p.i] = b.decompose(p.l, p.v)
		}
	}

	for _, v := range ans {
		if v {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
