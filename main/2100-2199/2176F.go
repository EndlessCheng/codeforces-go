package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2176F(in io.Reader, out io.Writer) {
	const mod = 998244353
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

	const mx = 200001
	omega := [mx]int{}
	for i := 2; i < mx; i++ {
		if omega[i] == 0 {
			for j := i; j < mx; j += i {
				omega[j]++
			}
		}
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		cnt := make([]int, n+1)
		for range n {
			var v int
			Fscan(in, &v)
			cnt[v]++
		}

		const u = 7
		res := [u * 2]int{}
		f := make([][u * 2]int, n+1)
		for v := n; v > 0; v-- {
			cntOmega := [u]int{}
			for w := v; w <= n; w += v {
				cntOmega[omega[w]] += cnt[w]
				for j := range u * 2 {
					f[v][j] -= f[w][j]
				}
			}
			for i := range u {
				for j := range i {
					f[v][i+j] += cntOmega[i] * cntOmega[j]
				}
				f[v][i*2] += cntOmega[i] * (cntOmega[i] - 1) / 2
			}
			for i := omega[v]; i < u*2; i++ {
				res[i-omega[v]] += f[v][i]
			}
		}

		ans := 0
		for i, v := range res {
			if v != 0 {
				ans = (ans + v%mod*pow(i, k)) % mod
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2176F(bufio.NewReader(os.Stdin), os.Stdout) }
