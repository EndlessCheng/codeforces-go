package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type matrix990 [][]int

func newMatrix990(n, m int) matrix990 {
	a := make(matrix990, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix990) mul(b matrix990) matrix990 {
	c := newMatrix990(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % 10000
			}
		}
	}
	return c
}

func (a matrix990) powMul(n int, f matrix990) matrix990 {
	res := f
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func p1990(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 1)
		return
	}
	f2 := matrix990{{2}, {1}, {1}}
	m := matrix990{
		{2, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	};
	Fprint(out, m.powMul(n-2, f2)[0][0])
}

//func main() { p1990(bufio.NewReader(os.Stdin), os.Stdout) }
