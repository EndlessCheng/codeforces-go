package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1228C(_r io.Reader, _w io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var x int
	var n int64
	Fscan(_r, &x, &n)
	primes := []int64{}
	for i := 2; i*i <= x; i++ {
		k := 0
		for ; x%i == 0; x /= i {
			k++
		}
		if k > 0 {
			primes = append(primes, int64(i))
		}
	}
	if x > 1 {
		primes = append(primes, int64(x))
	}
	ans := int64(1)
	for _, p := range primes {
		c := int64(0)
		for m := n; m > 0; m /= p {
			c += m / p
		}
		ans = ans * pow(p, c) % mod
	}
	Fprint(_w, ans)
}

//func main() { CF1228C(os.Stdin, os.Stdout) }
