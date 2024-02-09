package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type xorBasis59 struct {
	b   [20]int
	num int
}

func (b *xorBasis59) insert(v int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.num++
			return
		}
		v ^= b.b[i]
	}
}

func (b *xorBasis59) decompose(v int) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 > 0 {
			if b.b[i] == 0 {
				return false
			}
			v ^= b.b[i]
		}
	}
	return true
}

func cf959F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007

	var n, q, l, x int
	Fscan(in, &n, &q)
	a := make([]int, n)
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := range a {
		Fscan(in, &a[i])
		pow2[i+1] = pow2[i] * 2 % mod
	}
	type pair struct{ x, i int }
	qs := make([][]pair, n)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &x)
		qs[l-1] = append(qs[l-1], pair{x, i})
	}

	ans := make([]int, q)
	b := &xorBasis59{}
	for i, v := range a {
		b.insert(v)
		for _, p := range qs[i] {
			if b.decompose(p.x) {
				ans[p.i] = pow2[i+1-b.num]
			}
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf959F(os.Stdin, os.Stdout) }
