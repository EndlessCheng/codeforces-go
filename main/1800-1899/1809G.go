package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1809G(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	if n > 1 && a[n]-a[n-1] <= k {
		Fprint(out, 0)
		return
	}

	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}

	f := make([]int, n+1)
	g := make([]int, n+1)
	for i, j := 1, 1; i <= n; i++ {
		for ; j <= n && a[i]-a[j] > k; j++ {
		}
		f[i] = g[j-1] + 1
		g[i] = (g[i-1] + f[i]*inv[n-j]) % mod
	}

	ans := f[n] * inv[n] % mod
	for i := 1; i <= n; i++ {
		ans = ans * i % mod
	}
	Fprint(out, ans)
}

//func main() { cf1809G(bufio.NewReader(os.Stdin), os.Stdout) }
