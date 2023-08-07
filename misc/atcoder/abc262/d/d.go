package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := make([][]int, n+1)
	for m := 1; m <= n; m++ {
		for i := 0; i <= m; i++ {
			f[i] = make([]int, m)
		}
		f[0][0] = 1
		for _, v := range a {
			for j := m; j > 0; j-- {
				for k := 0; k < m; k++ {
					f[j][(k+v)%m] = (f[j][(k+v)%m] + f[j-1][k]) % mod
				}
			}
		}
		ans += f[m][0]
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
