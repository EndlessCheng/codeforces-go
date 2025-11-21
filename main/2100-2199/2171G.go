package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2171G(in io.Reader, out io.Writer) {
	const mod = 1_000_003
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
	f := [mod]int{1}
	for i := 1; i < mod; i++ {
		f[i] = f[i-1] * i % mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]uint, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		k := 99
		b := make([]uint, n)
		for i, v := range a {
			Fscan(in, &b[i])
			k = min(k, bits.Len(b[i]/v)-1)
		}

		cnt := make([]uint, k+1)
		ans := 1
		for i, v := range b {
			d := v>>k - a[i]
			cnt[0] += d
			ans = ans * pow(f[d], mod-2) % mod
			for v &= 1<<k - 1; v > 0; v &= v - 1 {
				cnt[k-bits.TrailingZeros(v)]++
			}
		}

		s := k
		for _, c := range cnt {
			s += int(c)
			if c < mod {
				ans = ans * f[c] % mod
			} else {
				ans = 0
			}
		}
		Fprintln(out, s, ans)
	}
}

//func main() { cf2171G(bufio.NewReader(os.Stdin), os.Stdout) }
