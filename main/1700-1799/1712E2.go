package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
type fenwick []int

func (f fenwick) update(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

func (f fenwick) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += f[i]
	}
	return
}

func cf1712E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var m int
	Fscan(in, &m)
	type query struct{ l, r, i int }
	qs := make([]query, m)
	for i := range qs {
		Fscan(in, &qs[i].l, &qs[i].r)
		qs[i].i = i
	}
	slices.SortFunc(qs, func(a, b query) int { return a.l - b.l })

	ans := make([]int, m)
	const mx int = 2e5 + 1
	divisors := [mx]int{}
	t := make(fenwick, mx)
	curL := mx
	for i := m - 1; i >= 0; i-- {
		l, r := qs[i].l, qs[i].r
		for curL > l {
			curL--
			for j := curL * 2; j < mx; j += curL {
				t.update(j, divisors[j])
				divisors[j]++
			}
		}
		ans[qs[i].i] = (r-l+1)*(r-l)*(r-l-1)/6 - t.pre(r) - max(r/6-(l+2)/3+1, 0) - max(r/15-(l+5)/6+1, 0)
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1712E2(bufio.NewReader(os.Stdin), os.Stdout) }
