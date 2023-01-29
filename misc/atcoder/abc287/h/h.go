package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
const U = bits.UintSize

type Bitset []uint

func (b Bitset) Has(p int) bool { return b[p/U]&(1<<(p%U)) > 0 }
func (b Bitset) Set(p int)      { b[p/U] |= 1 << (p % U) }
func (b Bitset) MergeFrom(c Bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, q int
	Fscan(in, &n, &m)
	vs := make([]Bitset, n)
	for i := range vs {
		vs[i] = make(Bitset, n/U+1)
		vs[i].Set(i)
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		vs[v-1].Set(w - 1)
	}
	Fscan(in, &q)
	qs := make([]struct{ x, y int }, q)
	ans := make([]int, q)
	for i := range qs {
		Fscan(in, &qs[i].x, &qs[i].y)
		qs[i].x--
		qs[i].y--
		ans[i] = n
	}

	for k := range vs {
		for i := range vs {
			if vs[i].Has(k) {
				vs[i].MergeFrom(vs[k])
			}
		}
		for i, q := range qs {
			if ans[i] == n && vs[q.x].Has(q.y) {
				ans[i] = k
			}
		}
	}
	for i, v := range ans {
		if v == n {
			Fprintln(out, -1)
		} else {
			Fprintln(out, max(v, max(qs[i].x, qs[i].y))+1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
