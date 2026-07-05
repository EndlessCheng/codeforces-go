package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2063F2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for n > 0 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
			n >>= 1
		}
		return res
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mx := n * 2
		f := make([]int, mx+1)
		f[0] = 1
		for i := 1; i <= mx; i++ {
			f[i] = f[i-1] * (i*4 - 2) % mod * pow(i+1, mod-2) % mod
		}

		x := make([]int, n+1)
		y := make([]int, n+1)
		a := make([]int, mx+1)
		s := make([]int, mx+1)
		fa := make([]int, mx+1)
		sz := make([]int, mx+1)
		del := make([]bool, mx+1)
		res := make([]int, n+1)
		it := 0
		for i := 1; i <= n; i++ {
			Fscan(in, &x[i], &y[i])
			a[x[i]] = 1
			a[y[i]] = 0
		}

		for i := 1; i <= mx; i++ {
			if a[i] == 1 {
				fa[i] = s[it]
				it++
				s[it] = i
			} else {
				it--
			}
		}

		var find func(int) int
		find = func(x int) int {
			if !del[x] {
				return x
			}
			fa[x] = find(fa[x])
			return fa[x]
		}

		mul := 1
		for i := n; i >= 1; i-- {
			res[i] = mul
			del[x[i]] = true
			k := find(x[i])
			mul = mul * pow(f[sz[x[i]]]*f[sz[k]]%mod, mod-2) % mod
			sz[k] += sz[x[i]] + 1
			mul = mul * f[sz[k]] % mod
		}
		res[0] = mul

		for _, v := range res {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2063F2(bufio.NewReader(os.Stdin), os.Stdout) }
