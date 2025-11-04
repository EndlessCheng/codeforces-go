package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod93 = 1_000_000_007

type matrix93 [][]int

func newMatrix93(n, m int) matrix93 {
	a := make(matrix93, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix93) mul(b matrix93) matrix93 {
	c := newMatrix93(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod93
			}
		}
	}
	return c
}

func (a matrix93) powMul(n int, f0 matrix93) matrix93 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf593E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, tp, x, y, t int
	Fscan(in, &n, &m, &q)

	size := n * m
	M := newMatrix93(size, size)
	set := func(x, y, val int) {
		i := x*m + y
		M[i][i] = val
		if y > 0 {
			M[i][i-1] = val
		}
		if y < m-1 {
			M[i][i+1] = val
		}
		if x > 0 {
			M[i][i-m] = val
		}
		if x < n-1 {
			M[i][i+m] = val
		}
	}

	for i := range n {
		for j := range m {
			set(i, j, 1)
		}
	}

	f := newMatrix93(size, 1)
	f[0][0] = 1

	pre := 1
	for range q {
		Fscan(in, &tp, &x, &y, &t)
		x--
		y--

		f = M.powMul(t-pre, f)
		pre = t

		i := x*m + y
		if tp == 1 {
			Fprintln(out, f[i][0])
		} else if tp == 2 {
			f[i][0] = 0
			set(x, y, 0)
		} else {
			set(x, y, 1)
		}
	}
}

//func main() { cf593E(bufio.NewReader(os.Stdin), os.Stdout) }
