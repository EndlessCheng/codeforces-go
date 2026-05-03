package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1987F2(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		b := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			if i < a[i] || (i-a[i])%2 > 0 {
				b[i] = 1e9
			} else {
				b[i] = i - a[i]
			}
		}

		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n+1)
		}
		for l := n; l >= 1; l-- {
			for r := l + 1; r <= n; r += 2 {
				f[l][r] = 1e9
				for p := l + 1; p < r; p += 2 {
					f[l][r] = min(f[l][r], max(f[l][p], f[p+1][r]-(p-l+1)))
				}
				if b[l] >= f[l+1][r-1] {
					f[l][r] = min(f[l][r], b[l])
				}
			}
		}

		g := make([]int, n+1)
		for i := 1; i <= n; i++ {
			g[i] = g[i-1]
			for j := i - 1; j >= 1; j -= 2 {
				if g[j-1] >= f[j][i] {
					g[i] = max(g[i], g[j-1]+(i-j+1))
				}
			}
		}
		Fprintln(out, g[n]/2)
	}
}

//func main() { cf1987F2(bufio.NewReader(os.Stdin), os.Stdout) }
