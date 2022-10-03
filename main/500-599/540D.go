package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF540D(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var a, b, c int
	Fscan(in, &a, &b, &c)
	n := max(max(a, b), c)
	f := make([][][]float64, n+1)
	for i := range f {
		f[i] = make([][]float64, n+1)
		for j := range f[i] {
			f[i][j] = make([]float64, n+1)
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			f[i][j][0] = 1
			for k := 1; k <= n; k++ {
				f[i][j][k] += (float64(i*j)*f[i][j-1][k] + float64(j*k)*f[i][j][k-1] + float64(i*k)*f[i-1][j][k]) / float64(i*j+j*k+i*k)
			}
		}
	}
	Fprintf(out, "%.12f %.12f %.12f", f[a][b][c], f[b][c][a], f[c][a][b])
}

//func main() { CF540D(os.Stdin, os.Stdout) }
