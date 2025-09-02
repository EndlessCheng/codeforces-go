package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf932E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	const inv2 = (mod + 1) / 2
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
	var n, k, ans int
	Fscan(in, &n, &k)
	s2 := make([]int, k+1)
	s2[0] = 1
	for i := 1; i <= k; i++ {
		for j := i; j > 0; j-- {
			s2[j] = (s2[j-1] + s2[j]*j) % mod
		}
		s2[0] = 0
	}

	pow2 := pow(2, n)
	for i := 1; i <= k; i++ {
		pow2 = pow2 * (n - i + 1) % mod * inv2 % mod
		ans = (ans + pow2*s2[i]) % mod
	}
	Fprint(out, ans)
}

//func main() { cf932E(os.Stdin, os.Stdout) }
