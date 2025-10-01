package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
var mod26 int

type matrix26 [][]int

func newMatrix26(n, m int) matrix26 {
	a := make(matrix26, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix26) mul(b matrix26) matrix26 {
	c := newMatrix26(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod26
			}
		}
	}
	return c
}

func (a matrix26) powMul(n int, f0 matrix26) matrix26 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf226C(in io.Reader, out io.Writer) {
	var l, r, k, n int
	Fscan(in, &mod26, &l, &r, &k)
	l--
	for i := 1; i*i <= r; i++ {
		if j := r / i; r/j-l/j >= k {
			n = j
			break
		}
		if r/i-l/i >= k {
			n = i
		}
	}
	m := matrix26{
		{1, 1},
		{1, 0},
	}
	f1 := matrix26{{1}, {0}}
	fn := m.powMul(n-1, f1)
	Fprint(out, fn[0][0]%mod26)
}

//func main() { cf226C(os.Stdin, os.Stdout) }
