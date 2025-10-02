package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1750D(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, m, pre, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &pre)
		primes := []int{}
		t := pre
		for p := 2; p*p <= t; p++ {
			if t%p > 0 {
				continue
			}
			for t /= p; t%p == 0; t /= p {
			}
			primes = append(primes, p)
		}
		if t > 1 {
			primes = append(primes, t)
		}

		ans := 1
		for range n - 1 {
			Fscan(in, &v)
			if v == pre {
				ans = ans * (m / v) % mod
				continue
			}
			if pre%v > 0 {
				ans = 0
				continue
			}
			m := m / v
			t := pre / v
			ps := []int{}
			for _, p := range primes {
				if t%p == 0 {
					ps = append(ps, p)
				}
			}
			res := 0
			for i := range 1 << len(ps) {
				mul := 1
				for j := uint(i); j > 0; j &= j - 1 {
					mul *= ps[bits.TrailingZeros(j)]
				}
				res += m / mul * (1 - bits.OnesCount(uint(i))%2*2)
			}
			ans = ans * res % mod
			pre = v
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1750D(bufio.NewReader(os.Stdin), os.Stdout) }
