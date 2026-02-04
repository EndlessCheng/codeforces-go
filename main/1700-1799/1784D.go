package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1784D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var n int
	Fscan(in, &n)
	m := 1 << n
	F := make([]int, m+2)
	F[0] = 1
	for i := 1; i <= m; i++ {
		F[i] = F[i-1] * i % mod
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	f[0][1] = m
	for i := 1; i <= n; i++ {
		mx := m - 1<<(n-i) + 1
		for j := i + 1; j <= mx; j++ {
			f[i][j] = (f[i-1][j-1]<<(n-i) + f[i][j-1]*(mx-j+1)) % mod
		}
	}

	for i := 1; i <= m; i++ {
		Fprintln(out, f[n][i]*F[m-i]%mod)
	}
}

//func main() { cf1784D(bufio.NewReader(os.Stdin), os.Stdout) }
