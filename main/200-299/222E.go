package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod22 = 1_000_000_007

type matrix22 [][]int

func newMatrix22(n, m int) matrix22 {
	a := make(matrix22, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix22) mul(b matrix22) matrix22 {
	c := newMatrix22(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod22
			}
		}
	}
	return c
}

// a^n * f1
func (a matrix22) powMul(n int, f1 matrix22) matrix22 {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf222E(in io.Reader, out io.Writer) {
	var n, sz, k, ans int
	var s string
	Fscan(in, &n, &sz, &k)
	f := func(b byte) byte {
		if b >= 'a' {
			return b - 'a'
		}
		return b - 'A' + 26
	}

	m := newMatrix22(sz, sz)
	for i := range m {
		for j := range m[i] {
			m[i][j] = 1
		}
	}
	for range k {
		Fscan(in, &s)
		m[f(s[1])][f(s[0])] = 0
	}

	f1 := newMatrix22(sz, 1)
	for i := range f1 {
		f1[i][0] = 1
	}

	fn := m.powMul(n-1, f1)
	for _, row := range fn {
		ans += row[0]
	}
	Fprint(out, ans%mod22)
}

//func main() { cf222E(bufio.NewReader(os.Stdin), os.Stdout) }
