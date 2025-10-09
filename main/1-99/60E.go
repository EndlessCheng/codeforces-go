package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
var mod60 int

type matrix60 [][]int

func newMatrix60(n, m int) matrix60 {
	a := make(matrix60, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix60) mul(b matrix60) matrix60 {
	c := newMatrix60(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod60
			}
		}
	}
	return c
}

func (a matrix60) powMul(n int, f0 matrix60) matrix60 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf60E(in io.Reader, out io.Writer) {
	var n, x, y, s0 int
	Fscan(in, &n, &x, &y, &mod60)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s0 += a[i]
	}
	s0 %= mod60
	if n == 1 {
		Fprint(out, s0)
		return
	}

	m := matrix60{{1, 1}, {1, 0}}
	f0 := matrix60{{a[n-1]}, {a[n-2]}}
	big := m.powMul(x, f0)[0][0]

	f := func(s0, c, x int) int {
		m := matrix60{{3, 1}, {0, 1}}
		f0 := matrix60{{s0}, {c}}
		return m.powMul(x, f0)[0][0]
	}
	sx := f(s0, -a[0]-a[n-1], x)
	Fprint(out, (f(sx, -a[0]-big, y)+mod60)%mod60)
}

//func main() { cf60E(bufio.NewReader(os.Stdin), os.Stdout) }
