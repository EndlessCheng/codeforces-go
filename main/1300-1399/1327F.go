package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1327F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var n, k, m int
	Fscan(in, &n, &k, &m)
	a := make([]struct{ l, r, x int }, m)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r, &a[i].x)
	}

	ans := 1
	f := make([]int, n+2)
	f[0] = 1
	for b := 0; b < k; b++ {
		maxL := make([]int, n+1)
		d := make([]int, n+2)
		for _, p := range a {
			if p.x>>b&1 == 0 {
				maxL[p.r] = max(maxL[p.r], p.l)
			} else {
				d[p.l]++
				d[p.r+1]--
			}
		}

		sumF := 1
		sumD := 0
		left := 0
		for i := 1; i <= n+1; i++ {
			for left < maxL[i-1] {
				sumF -= f[left]
				left++
			}
			sumD += d[i]
			if sumD > 0 {
				f[i] = 0
				continue
			}
			sumF %= mod
			f[i] = sumF
			sumF *= 2
		}
		ans = ans * f[n+1] % mod
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf1327F(os.Stdin, os.Stdout) }
