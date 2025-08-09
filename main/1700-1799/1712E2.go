package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick12 []int

func (f fenwick12) update(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

func (f fenwick12) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += f[i]
	}
	return
}

func cf1712E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var m, l, r int
	Fscan(in, &m)
	const mx int = 2e5 + 1
	type pair struct{ r, i int }
	qs := [mx][]pair{}
	for i := range m {
		Fscan(in, &l, &r)
		qs[l] = append(qs[l], pair{r, i})
	}

	ans := make([]int, m)
	divisors := [mx]int{}
	t := make(fenwick12, mx)
	for l := mx - 1; l > 0; l-- {
		for j := l * 2; j < mx; j += l {
			t.update(j, divisors[j])
			divisors[j]++
		}
		for _, p := range qs[l] {
			r := p.r
			ans[p.i] = (r-l+1)*(r-l)*(r-l-1)/6 - t.pre(r) - max(r/6-(l+2)/3+1, 0) - max(r/15-(l+5)/6+1, 0)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1712E2(bufio.NewReader(os.Stdin), os.Stdout) }
