package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1842G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n, m, v, x, ans int
	Fscan(in, &n, &m, &v)
	k := min(m, n)
	f := make([]int, k+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		Fscan(in, &x)
		w := v * i % mod
		for j := k; j > 0; j-- {
			f[j] = (f[j]*(x+v*j%mod) + f[j-1]*w%mod*(m-j+1)) % mod
		}
		f[0] = f[0] * x % mod
	}
	for j, fj := range f {
		ans = (ans + fj*pow(n, mod-1-j)) % mod
	}
	Fprint(out, ans)
}

//func main() { cf1842G(bufio.NewReader(os.Stdin), os.Stdout) }
