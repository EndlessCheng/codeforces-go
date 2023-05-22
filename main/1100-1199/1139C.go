package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1139C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	pow := func(x int64, n int) (res int64) {
		x %= mod
		res = 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var n, k, v, w, c int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &c)
		if c == 0 {
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
	}
	vis := make([]bool, n)
	var f func(int) int
	f = func(v int) int {
		vis[v] = true
		sz := 1
		for _, w := range g[v] {
			if !vis[w] {
				sz += f(w)
			}
		}
		return sz
	}
	ans := pow(int64(n), k)
	for i, b := range vis {
		if !b {
			ans -= pow(int64(f(i)), k)
		}
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF1139C(os.Stdin, os.Stdout) }
