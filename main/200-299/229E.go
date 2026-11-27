package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf229E(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, m)
	b := []int{}
	mx := 0

	for i := 0; i < m; i++ {
		var k int
		Fscan(in, &k)
		mx = max(mx, k)

		for j := 0; j < k; j++ {
			var x int
			Fscan(in, &x)
			a[i] = append(a[i], x)
			b = append(b, x)
		}
	}

	c := make([][]float64, mx+1)
	for i := range c {
		c[i] = make([]float64, mx+1)
	}

	f := make([][]float64, m+1)
	g := make([][]float64, m+1)
	for i := range f {
		f[i] = make([]float64, n+1)
		g[i] = make([]float64, n+1)
	}

	for i := 0; i <= mx; i++ {
		for j := 0; j <= i; j++ {
			if i > 0 && j > 0 {
				c[i][j] = c[i-1][j] + c[i-1][j-1]
			} else {
				c[i][j] = 1
			}
		}
	}

	slices.SortFunc(b, func(a, b int) int { return b - a })

	f[0][0], g[0][0] = 1, 1
	for i := 0; i < m; i++ {
		z := len(a[i])
		x, y := 0, 0
		for j := 0; j < z; j++ {
			if a[i][j] > b[n-1] {
				x++
			} else if a[i][j] == b[n-1] {
				y++
			}
		}
		for j := 0; j <= n; j++ {
			for k := 0; k <= y; k++ {
				if j+x+k <= n {
					g[i+1][j+x+k] += g[i][j]
					f[i+1][j+x+k] += f[i][j] / c[z][x+k]
				}
			}
		}
	}

	Fprintf(out, "%.9f", f[m][n]/g[m][n])
}

//func main() { cf229E(bufio.NewReader(os.Stdin), os.Stdout) }
