package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod98 = 1_000_000_007

type matrix98 [][]int

func newMatrix98(n, m int) matrix98 {
	a := make(matrix98, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix98) mul(b matrix98) matrix98 {
	c := newMatrix98(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod98
			}
		}
	}
	return c
}

func (a matrix98) powMul(n int, f0 matrix98) matrix98 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf498E(in io.Reader, out io.Writer) {
	var w int
	f := matrix98{{1}}
	for h := 1; h <= 7; h++ {
		m := newMatrix98(1<<h, 1<<h)
		for right := range m {
			for left := range m[right] {
				f0, f1 := 0, 1
				for i := range h {
					s := f0 + f1
					if left>>i&1 > 0 && right>>i&1 > 0 {
						f1 = f0
					} else {
						f1 = s
					}
					f0 = s
				}
				m[right][left] = f1
			}
		}

		Fscan(in, &w)
		// 在 f 的前面插入一堆 {0}，旧的状态就自动相当于高位是 1 了
		f = append(newMatrix98(len(f), 1), f...)
		f = m.powMul(w, f)
	}
	Fprint(out, f[1<<7-1][0])
}

//func main() { cf498E(os.Stdin, os.Stdout) }
