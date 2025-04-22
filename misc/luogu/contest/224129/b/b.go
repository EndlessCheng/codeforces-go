package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
const _w = bits.UintSize

func NewBitset(n int) Bitset { return make(Bitset, (n+_w-1)/_w) }

type Bitset []uint

func (b Bitset) Has(p int) bool { return b[p/_w]&(1<<(p%_w)) != 0 }
func (b Bitset) Set(p int)      { b[p/_w] |= 1 << (p % _w) }
func (b Bitset) Reset(p int)    { b[p/_w] &^= 1 << (p % _w) }
func (b Bitset) Flip(p int)     { b[p/_w] ^= 1 << (p % _w) }

func (b Bitset) And(c Bitset) {
	for i, v := range c {
		b[i] &= v
	}
}
func (b Bitset) OnesCount() (c int) {
	for _, v := range b {
		c += bits.OnesCount(v)
	}
	return
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]Bitset, n)
	type pair struct{ mx, cnt int }
	res := make([]pair, n)
	add := func(i, v int) {
		if v > res[i].mx {
			res[i].mx = v
			res[i].cnt = 1
		} else if v == res[i].mx {
			res[i].cnt++
		}
	}

	and := func(a, b Bitset) int {
		s := 0
		for k, v := range b {
			s += bits.OnesCount(v & a[k])
		}
		return s
	}
	
	for i := range a {
		a[i] = NewBitset(m)
		for j := 0; j < m; j++ {
			var v int
			Fscan(in, &v)
			if v == 1 {
				a[i].Set(j)
			}
		}

		for j, b := range a[:i] {
			s := and(b, a[i])
			add(j, s)
			add(i, s)
		}
	}

	ans := 0
	check := func() {
		c := 0
		for i := 1; i < n; i++ {
			s := and(a[i], a[0])
			if s >= res[i].mx {
				c++
			}
		}
		ans = max(ans, c)
	}
	
	check()
	for j := 0; j < m; j++ {
		if !a[0].Has(j) {
			a[0].Flip(j)
			check()
			a[0].Flip(j)
		}
	}
	Fprintln(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
