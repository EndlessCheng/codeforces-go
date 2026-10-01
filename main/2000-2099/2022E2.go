package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2022E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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
	const mul = 1 << 30 % mod
	inv := pow(mul, mod-2)

	var T, n, m, k, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k, &q)
		ans := pow(mul, n+m-1)
		if k == 0 {
			Fprintln(out, ans)
		}

		fa := make([]int, n+m+1)
		for i := 1; i <= n+m; i++ {
			fa[i] = i
		}
		dis := make([]int, n+m+1)

		var find func(int) int
		find = func(x int) int {
			if x == fa[x] {
				return x
			}
			z := find(fa[x])
			dis[x] ^= dis[fa[x]]
			fa[x] = z
			return z
		}

		merge := func(u, v, w int) {
			x, y := find(u), find(v)
			if x == y {
				if dis[u]^dis[v]^w != 0 {
					ans = 0
				}
				return
			}
			fa[x] = y
			dis[x] = w ^ dis[u] ^ dis[v]
			ans = ans * inv % mod
		}

		for i := 1; i <= k+q; i++ {
			var u, v, w int
			Fscan(in, &u, &v, &w)
			merge(u, v+n, w)
			if i >= k {
				Fprintln(out, ans%mod)
			}
		}
	}
}

//func main() { cf2022E2(bufio.NewReader(os.Stdin), os.Stdout) }
