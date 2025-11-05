package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod14 = 1_000_000_007

type matrix14 [][]int

func newMatrix14(n, m int) matrix14 {
	a := make(matrix14, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix14) mul(b matrix14) matrix14 {
	c := newMatrix14(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod14
			}
		}
	}
	return c
}

func (a matrix14) powMul(n int, f1 matrix14) matrix14 {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf514E(in io.Reader, out io.Writer) {
	var n, x, v int
	Fscan(in, &n, &x)
	const mx = 100
	cnt := [mx + 1]int{}
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}

	m := newMatrix14(mx+1, mx+1)
	m[0] = append(cnt[1:], 1)
	for j := 1; j < mx; j++ {
		m[j][j-1] = 1
	}
	m[mx][mx] = 1

	f0 := newMatrix14(mx+1, 1)
	f0[0][0] = 1
	f0[mx][0] = 1

	fn := m.powMul(x, f0)
	Fprint(out, fn[0][0])
}

//func main() { cf514E(bufio.NewReader(os.Stdin), os.Stdout) }
