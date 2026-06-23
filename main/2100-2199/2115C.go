package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2115C(in io.Reader, out io.Writer) {
	var T, n, lim int
	var pp float64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &lim, &pp)
		p := pp / 100
		bc := 4000
		bd := 0
		for i := 0; i < n; i++ {
			var x int
			Fscan(in, &x)
			x--
			bc = min(bc, x)
			bd += x
		}
		m := bc * n
		bd -= m

		f := make([][]float64, lim+1)
		for i := range f {
			f[i] = make([]float64, bd+1)
		}
		f[0][0] = 1
		for i := 1; i <= lim; i++ {
			for j := 1; j <= min(i, bd); j++ {
				f[i][j] = f[i-1][j-1]*(1-p) + f[i-1][j]*p
			}
		}

		g := make([]float64, m+1)
		g[0] = 1
		ans := 0.
		for i := lim; i >= bd; i-- {
			ans += f[i][bd] * g[max(bc-(i-bd), 0)*n]
			for j := m; j >= 1; j-- {
				v1 := g[j]
				if j >= n {
					v1 = max(v1, g[j-n])
				}
				g[j] = v1*p + max(g[j], g[j-1])*(1-p)
			}
		}
		Fprintf(out, "%.6f\n", ans)
	}
}

//func main() { cf2115C(bufio.NewReader(os.Stdin), os.Stdout) }
