package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod21 = 1_000_000_007

type matrix21 [][]int

func newMatrix21(n, m int) matrix21 {
	a := make(matrix21, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix21) mul(b matrix21) matrix21 {
	c := newMatrix21(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod21
			}
		}
	}
	return c
}

func (a matrix21) powMul(n int, f1 matrix21) matrix21 {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf821E(in io.Reader, out io.Writer) {
	var n, k, l, r, c int
	Fscan(in, &n, &k)
	f := newMatrix21(16, 1)
	f[0][0] = 1
	for p := range n {
		Fscan(in, &l, &r, &c)
		m := newMatrix21(16, 16)
		for i := range c + 1 {
			for j := max(i-1, 0); j <= min(i+1, c); j++ {
				m[i][j] = 1
			}
		}
		if p == n-1 {
			r = k
		}
		f = m.powMul(r-l, f)
	}
	Fprint(out, f[0][0])
}

//func main() { cf821E(bufio.NewReader(os.Stdin), os.Stdout) }
