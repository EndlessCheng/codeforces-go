package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1997F(in io.Reader, out io.Writer) {
	const mod = 998244353
	fib := [26]int{0, 1, 1}
	for i := 3; i <= 25; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	var n, x, m int
	Fscan(in, &n, &x, &m)
	maxW := n * fib[x]

	g := make([]int, maxW+1)
	for i := range g {
		g[i] = 1e9
	}
	g[0] = 0
	for i := 1; i <= maxW; i++ {
		for j := 1; fib[j] <= i; j++ {
			g[i] = min(g[i], g[i-fib[j]]+1)
		}
	}

	f := make([][]int32, n+1)
	for i := range f {
		f[i] = make([]int32, maxW+1)
	}
	f[0][0] = 1
	for _, w := range fib[1 : x+1] {
		for j := 1; j <= n; j++ {
			up := j * fib[x]
			for k := w; k <= up; k++ {
				f[j][k] = (f[j][k] + f[j-1][k-w]) % mod
			}
		}
	}

	ans := 0
	for i := n; i <= maxW; i++ {
		if g[i] == m {
			ans += int(f[n][i])
		}
	}
	Fprint(out, ans%mod)
}

//func main() { cf1997F(bufio.NewReader(os.Stdin), os.Stdout) }
