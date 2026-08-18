package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf437E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
	Fscan(in, &n)
	x := make([]int, n)
	y := make([]int, n)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		Fscan(in, &x[i], &y[i])
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += x[i]*y[(i+1)%n] - y[i]*x[(i+1)%n]
	}
	if sum < 0 {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			x[i], x[j] = x[j], x[i]
			y[i], y[j] = y[j], y[i]
		}
	}
	for i := 0; i < n-1; i++ {
		f[i][i+1] = 1
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				if (x[i]-x[k])*(y[j]-y[k]) < (y[i]-y[k])*(x[j]-x[k]) {
					f[i][j] = (f[i][j] + f[i][k]*f[k][j]) % mod
				}
			}
		}
	}
	Fprint(out, f[0][n-1])
}

//func main() { cf437E(bufio.NewReader(os.Stdin), os.Stdout) }
