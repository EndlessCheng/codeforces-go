package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1392G(in io.Reader, out io.Writer) {
	var n, m, k int
	var s, t string
	Fscan(in, &n, &m, &k, &s, &t)
	u := 1<<k - 1
	tot, ans := k, -1
	a := make([]int, k)
	b := make([]int, k)
	f := make([]int, 1<<k)
	g := make([]int, 1<<k)
	for i := range g {
		g[i] = n + 1
	}
	for i := range a {
		a[i] = i
		tot -= int(s[i]&1) + int(t[i]&1)
	}

	for i := range n + 1 {
		x, y := 0, 0
		for j := range k {
			x ^= int(s[a[j]]-'0') << j
			y ^= int(t[a[j]]-'0') << j
		}
		f[y] = i
		g[x] = min(g[x], i)

		if i == n {
			break
		}

		Fscan(in, &x, &y)
		for j := range k {
			b[a[j]] = j
		}
		for j := 1; j <= k; j++ {
			v := j
			if j == x {
				v = y
			} else if j == y {
				v = x
			}
			a[b[j-1]] = v - 1
		}
	}

	for i := range k {
		for j := range u + 1 {
			if j>>i&1 != 0 {
				continue
			}
			f[j] = max(f[j], f[j^(1<<i)])
			g[j] = min(g[j], g[j^(1<<i)])
		}
	}

	x, y := 0, 0
	for i := range u + 1 {
		ones := bits.OnesCount(uint(i))
		if f[i]-g[i] >= m && ones > ans {
			ans = ones
			x, y = g[i]+1, f[i]
		}
	}

	Fprintln(out, ans*2+tot)
	Fprintln(out, x, y)
}

//func main() { cf1392G(bufio.NewReader(os.Stdin), os.Stdout) }
