package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
const mod91 = 1_000_000_007

type matrix91 [][]int

func newMatrix91(n, m int) matrix91 {
	a := make(matrix91, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix91) mul(b matrix91) matrix91 {
	c := newMatrix91(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod91
			}
		}
	}
	return c
}

func (a matrix91) powMul(n int, f0 matrix91) matrix91 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf691E(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	m := newMatrix91(n, n)
	for i, v := range a {
		for j, w := range a[:i+1] {
			if bits.OnesCount(uint(v^w))%3 == 0 {
				m[i][j] = 1
				m[j][i] = 1
			}
		}
	}

	f0 := newMatrix91(n, 1)
	for i := range f0 {
		f0[i][0] = 1
	}

	fk := m.powMul(k-1, f0)
	for _, row := range fk {
		ans += row[0]
	}
	Fprint(out, ans%mod91)
}

//func main() { cf691E(os.Stdin, os.Stdout) }
