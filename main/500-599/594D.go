package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod94 = 1_000_000_007

type fenwick94 []int

func (f fenwick94) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = f[i] * val % mod94
	}
}

func (f fenwick94) pre(i int) int {
	res := 1
	for ; i > 0; i &= i - 1 {
		res = res * f[i] % mod94
	}
	return res
}

func cf594D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e6 + 1
	lpf := [mx]int{1: 1}
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	inv := func(x int) int {
		res := 1
		for n := mod94 - 2; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod94
			}
			x = x * x % mod94
		}
		return res
	}

	var n, q, l, r int
	Fscan(in, &n)
	a := make([]int, n)
	mul := 1
	for i := range a {
		Fscan(in, &a[i])
		mul = mul * a[i] % mod94
	}
	invM := make([]int, n+1)
	invM[n] = inv(mul)
	for i := n - 1; i >= 0; i-- {
		invM[i] = invM[i+1] * a[i] % mod94
	}

	Fscan(in, &q)
	type pair struct{ l, i int }
	qs := make([][]pair, n+1)
	for i := range q {
		Fscan(in, &l, &r)
		qs[r] = append(qs[r], pair{l - 1, i})
	}

	ans := make([]int, q)
	t := make(fenwick94, n+1)
	for i := range t {
		t[i] = 1
	}
	last := [mx]int{}
	mul = 1
	for i, x := range a {
		i++
		mul = mul * x % mod94
		for x > 1 {
			p := lpf[x]
			for x /= p; x%p == 0; x /= p {
			}
			if last[p] > 0 {
				t.update(last[p], 1+inv(p-1))
			}
			t.update(i, 1-inv(p)+mod94)
			last[p] = i
		}
		for _, p := range qs[i] {
			ans[p.i] = mul * invM[p.l] % mod94 * t.pre(i) % mod94 * inv(t.pre(p.l)) % mod94
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf594D(bufio.NewReader(os.Stdin), os.Stdout) }
