package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf838D(in io.Reader, out io.Writer) {
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
	var n, m int
	Fscan(in, &n, &m)
	Fprint(out, pow(2*(n+1), m)*(n-m+1)%mod*pow(n+1, mod-2)%mod)
}

//func main() { cf838D(os.Stdin, os.Stdout) }
