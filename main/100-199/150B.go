package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf150B(in io.Reader, out io.Writer) {
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
	var n, m, k int
	Fscan(in, &n, &m, &k)
	if k == 1 || k > n {
		Fprint(out, pow(m, n))
	} else if k == n {
		Fprint(out, pow(m, (n+1)/2))
	} else if k%2 == 0 {
		Fprint(out, m)
	} else {
		Fprint(out, m*m)
	}
}

//func main() { cf150B(os.Stdin, os.Stdout) }
