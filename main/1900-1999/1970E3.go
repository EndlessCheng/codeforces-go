package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod70 = 1_000_000_007

type matrix70 [][]int

func newMatrix70(n, m int) matrix70 {
	a := make(matrix70, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix70) mul(b matrix70) matrix70 {
	c := newMatrix70(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod70
			}
		}
	}
	return c
}

func (a matrix70) powMul(n int, f0 matrix70) matrix70 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf1970E3(in io.Reader, out io.Writer) {
	var k, n, l, ss, sl int
	Fscan(in, &k, &n)
	a := make([]int, k)
	for i := range a {
		Fscan(in, &a[i])
		ss += a[i]
	}

	m := newMatrix70(2, 2)
	f1 := newMatrix70(2, 1)
	f1[0][0] = a[0]
	for i, s := range a {
		Fscan(in, &l)
		if i == 0 {
			f1[1][0] = l
		}
		sl += l
		m[0][0] = (m[0][0] + s*(s+l)) % mod70
		m[0][1] = (m[0][1] + s*s) % mod70
		m[1][0] = (m[1][0] + l*(s+l)) % mod70
		m[1][1] = (m[1][1] + l*s) % mod70
	}

	fn := m.powMul(n-1, f1)
	Fprint(out, (fn[0][0]*(ss+sl)+fn[1][0]*ss)%mod70)
}

//func main() { cf1970E3(bufio.NewReader(os.Stdin), os.Stdout) }
