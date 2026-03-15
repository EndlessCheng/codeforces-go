package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1372E(in io.Reader, out io.Writer) {
	var n, m, K, t int
	Fscan(in, &n, &m)
	var l, r [10001]int
	var c, f [102][102]int
	for i := 1; i <= n; i++ {
		Fscan(in, &K)
		for range K {
			t++
			Fscan(in, &l[t], &r[t])
			for j := l[t]; j <= r[t]; j++ {
				c[i][j] = t
			}
		}
	}

	for i := 1; i <= m; i++ {
		for j, k := 1, i; k <= m; j, k = j+1, k+1 {
			for p := j; p <= k; p++ {
				u := 0
				for q := 1; q <= n; q++ {
					if l[c[q][p]] >= j && r[c[q][p]] <= k {
						u++
					}
				}
				f[j][k] = max(f[j][k], f[j][p-1]+f[p+1][k]+u*u)
			}
		}
	}
	Fprint(out, f[1][m])
}

//func main() { cf1372E(bufio.NewReader(os.Stdin), os.Stdout) }
