package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1004(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	for {
		var x, y, v int
		Fscan(in, &x, &y, &v)
		if x == 0 {
			break
		}
		a[x-1][y-1] = v
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = -1e9
		}
	}
	f[1][1] = a[0][0]
	for t := 1; t < n*2-1; t++ {
		for j := min(t, n-1); j >= max(t-n+1, 0); j-- {
			for k := min(t, n-1); k >= j; k-- {
				f[j+1][k+1] = max(f[j+1][k+1], f[j+1][k], f[j][k+1], f[j][k]) + a[t-j][j]
				if k != j {
					f[j+1][k+1] += a[t-k][k]
				}
			}
		}
	}
	Fprint(out, f[n][n])
}

//func main() { p1004(bufio.NewReader(os.Stdin), os.Stdout) }
