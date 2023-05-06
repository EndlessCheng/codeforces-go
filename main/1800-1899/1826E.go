package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
const _w26 = bits.UintSize

func NewBitset26(n int) Bitset26 { return make(Bitset26, n/_w26+1) } // (n+_w-1)/_w

type Bitset26 []uint

func (b Bitset26) Has(p int) bool { return b[p/_w26]&(1<<(p%_w26)) != 0 } // get
func (b Bitset26) Set(p int)      { b[p/_w26] |= 1 << (p % _w26) }        // 置 1
func (b Bitset26) SetAll1() {
	for i := range b {
		b[i] = math.MaxUint
	}
}
func (b Bitset26) IntersectionFrom(c Bitset26) {
	for i, v := range c {
		b[i] &= v
	}
}

func CF1826E(_r io.Reader, out io.Writer) {
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	ri := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	m, n := ri(), ri()
	type pair struct {
		p int
		r []int
	}
	a := make([]pair, n)
	for i := range a {
		a[i].p = ri()
		a[i].r = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for _, p := range a {
			p.r[i] = ri()
		}
	}

	sort.Slice(a, func(i, j int) bool { return a[i].r[0] < a[j].r[0] })

	from := make([]Bitset26, n)
	for i := range from {
		from[i] = NewBitset26(n)
		from[i].SetAll1()
	}

	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	for city := 0; city < m; city++ {
		sort.Slice(ids, func(i, j int) bool { return a[ids[i]].r[city] < a[ids[j]].r[city] })
		cur := NewBitset26(n)
		j := 0
		for _, i := range ids {
			for a[ids[j]].r[city] < a[i].r[city] {
				cur.Set(ids[j])
				j++
			}
			from[i].IntersectionFrom(cur)
		}
	}

	ans := int64(0)
	f := make([]int64, n)
	for i, p := range a {
		f[i] = 0
		for j := i - 1; j >= 0; j-- {
			if f[j] > f[i] && from[i].Has(j) {
				f[i] = f[j]
			}
		}
		f[i] += int64(p.p)
		ans = max(ans, f[i])
	}
	Fprint(out, ans)
}

//func main() { CF1826E(os.Stdin, os.Stdout) }
