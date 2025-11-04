package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
const mod52 = 1_000_000_007

type matrix52 [][]int

func newMatrix52(n, m int) matrix52 {
	a := make(matrix52, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix52) mul(b matrix52) matrix52 {
	c := newMatrix52(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod52
			}
		}
	}
	return c
}

func (a matrix52) powMul(n int, f0 matrix52) matrix52 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf1152F2(in io.Reader, out io.Writer) {
	var n, k, d, ans int
	Fscan(in, &n, &k, &d)
	f := func(i, j int) int { return i*(k+1) + j }

	sz := (k + 1) << d
	m := newMatrix52(sz, sz)
	for s := range 1 << d {
		t := s << 1 & (1<<d - 1)
		c := bits.OnesCount(uint(s))
		for j := c; j <= k; j++ {
			// 不访问星球 i
			m[f(t, j)][f(s, j)] = 1
			if j < k {
				// 访问星球 i
				// i 可以插在这 c 个星球的右边，或者作为第一个访问的星球
				m[f(t^1, j+1)][f(s, j)] = c + 1
			}
		}
	}

	f0 := newMatrix52(sz, 1)
	f0[0][0] = 1

	fn := m.powMul(n, f0)
	for s := range 1 << d {
		ans += fn[f(s, k)][0]
	}
	Fprint(out, ans%mod52)
}

//func main() { cf1152F2(bufio.NewReader(os.Stdin), os.Stdout) }
