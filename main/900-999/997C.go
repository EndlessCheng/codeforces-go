package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf997C(in io.Reader, out io.Writer) {
	const M = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}

	const mx int = 1e6
	inv := [mx + 1]int{}
	inv[1] = 1
	for i := 2; i <= mx; i++ {
		inv[i] = (M - M/i) * inv[M%i] % M
	}

	var n int
	Fscan(in, &n)
	p3 := pow(3, n)
	ans := pow(p3-1, n)*3 - pow(p3-3, n)*2 - pow(p3, n)
	comb := n%2*2 - 1
	p3 = 1
	for i := range n {
		ans = (ans + comb*pow(p3-1, n)*3) % M
		comb = -comb * (n - i) % M * inv[i+1] % M
		p3 = p3 * 3 % M
	}
	Fprint(out, (ans+M)%M)
}

//func main() { cf997C(os.Stdin, os.Stdout) }
