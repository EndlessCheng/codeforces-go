package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF543D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, v int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		v--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dp := make([]int64, n)
	ex := make([]int64, n)
	var f func(v, fa int)
	f = func(v, fa int) {
		z := false
		dp[v] = 1
		ex[v] = 1
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
				dw := dp[w] + 1
				dp[v] = dp[v] * dw % mod
				if z || dw != mod {
					ex[v] = ex[v] * dw % mod
				} else {
					z = true
				}
			}
		}
	}
	f(0, -1)

	ans := make([]int64, n)
	var reroot func(v, fa int, dpFa int64)
	reroot = func(v, fa int, dpFa int64) {
		ans[v] = dp[v] * (dpFa + 1) % mod
		for _, w := range g[v] {
			if w != fa {
				df := int64(0)
				if dp[w]+1 == mod {
					df = ex[v] * (dpFa + 1) % mod
				} else {
					df = ans[v] * pow(dp[w]+1, mod-2) % mod
				}
				reroot(w, v, df)
			}
		}
	}
	reroot(0, -1, 0)
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF543D(os.Stdin, os.Stdout) }
