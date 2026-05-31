package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf891E(in io.Reader, out io.Writer) {
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

	var n, k, v int
	Fscan(in, &n, &k)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		for j := i; j > 0; j-- {
			f[j] = (v*f[j]%mod - f[j-1] + mod) % mod
		}
		f[0] = f[0] * v % mod
	}

	ans := f[0]
	m := 1
	invN := pow(n, mod-2)
	for i := 0; i <= n; i++ {
		ans = (ans - f[i]*m) % mod
		m = m * invN % mod * k % mod
		k--
	}
	Fprint(out, (ans+mod)%mod)
}

//func main() { cf891E(bufio.NewReader(os.Stdin), os.Stdout) }
