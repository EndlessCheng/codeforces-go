package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 本题需要用 int 存储计算，否则会 TLE
// 更进一步的优化见 https://codeforces.com/contest/1218/submission/118704484

// github.com/EndlessCheng/codeforces-go
func CF1218D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1e9 + 7
	const inv2 = (mod + 1) / 2
	const mx = 1 << 17
	fwt := func(a []int) {
		for l, k := 2, 1; l <= mx; l, k = l<<1, k<<1 {
			for i := 0; i < mx; i += l {
				for j := 0; j < k; j++ {
					a[i+j], a[i+j+k] = (a[i+j]+a[i+j+k])%mod, (a[i+j]-a[i+j+k])%mod
				}
			}
		}
	}
	ifwt := func(a []int) {
		for l, k := 2, 1; l <= mx; l, k = l<<1, k<<1 {
			for i := 0; i < mx; i += l {
				for j := 0; j < k; j++ {
					a[i+j], a[i+j+k] = int(int64(a[i+j]+a[i+j+k])*inv2%mod), int(int64(a[i+j]-a[i+j+k])*inv2%mod)
				}
			}
		}
	}

	var n, m, v, w, wt, xor int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
		xor ^= wt
	}

	// 仙人掌找环
	cnt := [][]int{}
	s := []nb{{1, 0}}
	vis := make([]int8, n+1)
	var f func(int, int)
	f = func(v, fa int) {
		vis[v] = 1
		for _, e := range g[v] {
			if w := e.to; vis[w] == 0 {
				s = append(s, e)
				f(w, v)
			} else if w != fa && vis[w] == 1 {
				c := make([]int, mx)
				for i := len(s) - 1; s[i].to != w; i-- {
					c[s[i].wt]++
				}
				c[e.wt]++
				cnt = append(cnt, c)
			}
		}
		vis[v] = 2
		s = s[:len(s)-1]
	}
	f(1, 0)

	has := make([]int, mx)
	for i, c := range cnt[0] {
		if c != 0 {
			has[i] = 1
		}
	}
	for _, c := range cnt {
		fwt(c)
	}
	for _, c := range cnt[1:] {
		fwt(has)
		for j, v := range c {
			cnt[0][j] = int(int64(cnt[0][j]) * int64(v) % mod)
			has[j] = int(int64(has[j]) * int64(v) % mod)
		}
		ifwt(has)
		for i, v := range has {
			if v != 0 {
				has[i] = 1
			}
		}
	}
	ifwt(cnt[0])

	for i := 0; ; i++ {
		if has[xor^i] != 0 {
			Fprint(out, i, (cnt[0][xor^i]+mod)%mod)
			break
		}
	}
}

//func main() { CF1218D(os.Stdin, os.Stdout) }
