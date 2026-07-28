package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1608F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] -= k
	}

	C := make([][]int, n+1)
	for i := range C {
		C[i] = make([]int, n+1)
	}
	C[0][0] = 1
	for i := 1; i <= n; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	f := make([][][]int, 2)
	for i := range f {
		f[i] = make([][]int, 2*k+1)
		for j := range f[i] {
			f[i][j] = make([]int, n+2)
		}
	}

	g := make([]int, n+2)
	f[0][0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= 2*k; j++ {
			for p := 0; p <= i; p++ {
				f[i&1][j][p] = 0
			}
		}
		for p := 0; p <= i; p++ {
			g[p] = 0
		}
		for j := 0; j <= 2*k && a[i-1]+j < a[i]-1; j++ {
			for p := 0; p+a[i]-2-(a[i-1]+j) <= i; p++ {
				g[p] = (g[p] + f[i&1^1][j][p+a[i]-2-(a[i-1]+j)]) % mod
			}
		}
		for j := 0; j <= 2*k; j++ {
			if a[i]+j < 0 {
				continue
			}
			for p := 0; p <= i; p++ {
				g[p] = g[p+1]
			}
			cur := a[i] + j - a[i-1]
			if cur-1 >= 0 && cur-1 <= 2*k {
				for p := 0; p <= i; p++ {
					g[p] = (g[p] + f[i&1^1][cur-1][p]) % mod
				}
			}
			for p := 0; p <= i; p++ {
				f[i&1][j][p] = g[p]
			}
			if cur >= 0 && cur <= 2*k {
				for p := 0; p <= i; p++ {
					f[i&1][j][p] = (f[i&1][j][p] + f[i&1^1][cur][p]*(a[i]+j+p)) % mod
					if p > 0 {
						f[i&1][j][p] = (f[i&1][j][p] + f[i&1^1][cur][p-1]*p) % mod
					}
				}
			}
		}
	}

	ans := 0
	for j := 0; j <= 2*k; j++ {
		if a[n]+j >= 0 {
			for p := 0; p <= n-(a[n]+j); p++ {
				ans = (ans + C[n-(a[n]+j)][p]*f[n&1][j][p]) % mod
			}
		}
	}
	Fprintln(out, ans)
}

//func main() { cf1608F(bufio.NewReader(os.Stdin), os.Stdout) }
