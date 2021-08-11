package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1225E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7

	var n, m int
	Fscan(in, &n, &m)
	if n == 1 && m == 1 {
		Fprint(out, 1)
		return
	}
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	r := make([][]int, n+1)
	c := make([][]int, n+1)
	dpr := make([][]int64, n+1) // dp 前缀和
	dpc := make([][]int64, n+1)
	for i := range r {
		r[i] = make([]int, m+1)
		c[i] = make([]int, m+1)
		dpr[i] = make([]int64, m+1)
		dpc[i] = make([]int64, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			r[i][j] = r[i][j+1]
			c[i][j] = c[i+1][j]
			if a[i][j] == 'R' {
				r[i][j]++
				c[i][j]++
			}
			if i == n-1 && j == m-1 {
				dpc[i][j] = 1
				dpr[i][j] = 1
			} else {
				dpc[i][j] = (dpc[i+1][j] + dpr[i][j+1] - dpr[i][m-r[i][j+1]]) % mod
				dpr[i][j] = (dpr[i][j+1] + dpc[i+1][j] - dpc[n-c[i+1][j]][j]) % mod
			}
		}
	}
	Fprint(out, ((dpc[0][0]-dpc[1][0]+dpr[0][0]-dpr[0][1])%mod+mod)%mod)
}

//func main() { CF1225E(os.Stdin, os.Stdout) }
