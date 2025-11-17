package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2140E2(in io.Reader, out io.Writer) {
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

	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([]int, k)
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
		}

		f := make([]byte, 1<<n)
		f[1] = 1
		for sz := 2; sz <= n; sz++ {
			if (n-sz)%2 == 0 {
			o:
				for mask := 1<<sz - 1; mask > 0; mask-- {
					for _, i := range a {
						if i >= sz {
							break
						}
						if f[mask>>(i+1)<<i|mask&(1<<i-1)] > 0 {
							f[mask] = 1
							continue o
						}
					}
					f[mask] = 0
				}
			} else {
			o2:
				for mask := 1<<sz - 1; mask > 0; mask-- {
					for _, i := range a {
						if i >= sz {
							break
						}
						if f[mask>>(i+1)<<i|mask&(1<<i-1)] == 0 {
							f[mask] = 0
							continue o2
						}
					}
					f[mask] = 1
				}
			}
		}

		cnt := make([]int, n+1)
		for mask, v := range f {
			cnt[bits.OnesCount(uint(mask))] += int(v)
		}

		ans := 0
		for i, c := range cnt {
			s := 0
			for low := 1; low <= m; low++ {
				s = (s + pow(m-low+1, i)*pow(low-1, n-i)) % mod
			}
			ans += s * c
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { cf2140E2(bufio.NewReader(os.Stdin), os.Stdout) }
