package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type matrix1939 [][]int

func newMatrix1939(n, m int) matrix1939 {
	a := make(matrix1939, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix1939) mul(b matrix1939) matrix1939 {
	c := newMatrix1939(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % 1_000_000_007
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix1939) powMul(n int, f0 matrix1939) matrix1939 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func p1939(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n <= 3 {
			Fprintln(out, 1)
			continue
		}
		m := matrix1939{
			{1, 0, 1},
			{1, 0, 0},
			{0, 1, 0},
		}
		f0 := matrix1939{
			{1},
			{1},
			{1},
		}
		Fprintln(out, m.powMul(n-3, f0)[0][0])
	}
}

//func main() { p1939(bufio.NewReader(os.Stdin), os.Stdout) }
