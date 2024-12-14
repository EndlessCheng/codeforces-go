package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2008F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		x %= mod
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans, s := 0, 0
		for range n {
			Fscan(in, &v)
			ans = (ans + v*s) % mod
			s = (s + v) % mod
		}
		Fprintln(out, ans*pow(n*(n-1)/2, mod-2)%mod)
	}
}

//func main() { cf2008F(bufio.NewReader(os.Stdin), os.Stdout) }
