package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type mat46 [][]int

func newMat46(n int) mat46 {
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	return a
}

func mul46(a, b mat46) mat46 {
	const mod = 1_000_000_007
	c := newMat46(len(a))
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

func cf946F(in io.Reader, out io.Writer) {
	var n, x int
	var s string
	Fscan(in, &n, &x, &s)
	f := make([]mat46, x+2)
	for i := range f {
		f[i] = newMat46(n + 1)
	}
	for i, b := range s {
		f[b&1][i+1][i] = 1
		f[0][i][i] = 1
		f[1][i][i] = 1
	}
	f[0][0][0] = 2
	f[1][0][0] = 2
	f[0][n][n] = 2
	f[1][n][n] = 2
	for i := 2; i <= x; i++ {
		f[i] = mul46(f[i-2], f[i-1])
	}
	Fprint(out, f[x][n][0])
}

//func main() { cf946F(bufio.NewReader(os.Stdin), os.Stdout) }
