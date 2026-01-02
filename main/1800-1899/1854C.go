package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1854C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	const inv2 = (mod + 1) / 2
	var n, m, ans int
	Fscan(in, &n, &m)

	f := make([][]int, m+2)
	for i := range f {
		f[i] = make([]int, m+2)
	}
	for i := m; i >= 1; i-- {
		f[i][m+1] = m + 1 - i
		for j := m; j > i; j-- {
			f[i][j] = (f[i][j+1] + f[i+1][j] + 1) * inv2 % mod
		}
	}

	a := make([]int, n+1)
	for i := range n {
		Fscan(in, &a[i])
	}
	a[n] = m + 1

	for i, v := range a[:n] {
		ans += f[v][a[i+1]]
	}
	Fprint(out, ans%mod)
}

//func main() { cf1854C(os.Stdin, os.Stdout) }
