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

func (a matrix21) powMul(n int, f0 matrix21) matrix21 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf621E(in io.Reader, out io.Writer) {
	var n, b, k, x, v int
	Fscan(in, &n, &b, &k, &x)
	cnt := [10]int{}
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}

	m := newMatrix21(x, x)
	for v, c := range cnt {
		for r := range x {
			m[(r*10+v)%x][r] += c
		}
	}

	f0 := newMatrix21(x, 1)
	f0[0][0] = 1
	fn := m.powMul(b, f0)
	Fprint(out, fn[k][0])
}

//func main() { cf621E(bufio.NewReader(os.Stdin), os.Stdout) }
