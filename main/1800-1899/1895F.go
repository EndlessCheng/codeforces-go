package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod
			}
		}
	}
	return c
}

// a^n * f1
func (a matrix) powMul(n int, f1 matrix) matrix {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf1895F(in io.Reader, out io.Writer) {
	var T, n, x, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &k)
		ans := (x + k) * pow(k*2+1, n-1)

		if x > 0 {
			m := newMatrix(x, x)
			for j := range m {
				for v := max(j-k, 0); v <= min(j+k, x-1); v++ {
					m[j][v] = 1
				}
			}
			f1 := newMatrix(x, 1)
			for j := range f1 {
				f1[j][0] = 1
			}
			fn := m.powMul(n-1, f1)
			for _, row := range fn {
				ans -= row[0]
			}
		}

		Fprintln(out, (ans%mod+mod)%mod)
	}
}

//func main() { cf1895F(bufio.NewReader(os.Stdin), os.Stdout) }
