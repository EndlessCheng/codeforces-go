package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf367E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m, x int
	Fscan(in, &n, &m, &x)
	if n > m {
		Fprint(out, 0)
		return
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, i+2)
	}
	f[0][0] = 1
	for i := 1; i <= m; i++ {
		for l := min(i, n); l > 0; l-- {
			for r := l; r > 0; r-- {
				if i == x {
					f[l][r] = (f[l-1][r-1] + f[l-1][r]) % mod
				} else {
					f[l][r] = (f[l-1][r-1] + f[l-1][r] + f[l][r-1] + f[l][r]) % mod
				}
			}
			if i == x {
				f[l][0] = f[l-1][0]
			} else {
				f[l][0] = (f[l-1][0] + f[l][0]) % mod
			}
		}
		if i == x {
			f[0][0] = 0
		}
	}

	ans := f[n][n]
	for i := 2; i <= n; i++ {
		ans = ans * i % mod
	}
	Fprint(out, ans)
}

//func main() { cf367E(os.Stdin, os.Stdout) }
