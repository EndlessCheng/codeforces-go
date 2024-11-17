package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2027D2(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		f := make([]int, n+1)
		for j := range n {
			f[j] = 1e18
		}
		cnt := make([]int, n+1)
		cnt[n] = 1
		for i := m - 1; i >= 0; i-- {
			minK, maxK, sumC := n-1, n, 1
			for j := n - 1; j >= 0; j-- {
				for s[maxK] > s[j]+b[i] {
					sumC -= cnt[maxK]
					maxK--
				}
				for minK > j && f[minK] <= f[maxK] {
					sumC += cnt[minK]
					minK--
				}
				res := f[maxK] + m - 1 - i
				if res < f[j] {
					f[j] = res
					cnt[j] = sumC % mod
				} else if res == f[j] {
					cnt[j] = (cnt[j] + sumC) % mod
				}
			}
		}
		if f[0] >= 1e18 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, f[0], (cnt[0]+mod)%mod)
		}
	}
}

//func main() { cf2027D2(bufio.NewReader(os.Stdin), os.Stdout) }
