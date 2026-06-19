package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1874D(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	f := make([][]float64, n+1)
	for i := range f {
		f[i] = make([]float64, m+1)
		for j := range f[i] {
			f[i][j] = 1e99
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 1; k <= min(j, m/(n-i+1)); k++ {
				f[i][j] = min(f[i][j], f[i-1][j-k]+float64(j-k)/float64(k))
			}
		}
	}
	Fprintf(out, "%.9f", f[n][m]*2+float64(n))
}

//func main() { cf1874D(os.Stdin, os.Stdout) }
