package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1540B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	const mod2 = (mod + 1) / 2
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, v, w int
	Fscan(in, &n)
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = 1e9
		}
	}
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		dis[v][w] = 1
		dis[w][v] = 1
	}
	for k := range dis {
		for i := range dis {
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	f := make([][]int64, n+1)
	for i := range f {
		f[i] = make([]int64, n+1)
		f[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = (f[i-1][j] + f[i][j-1]) * mod2 % mod
		}
	}
	ans := int64(0)
	for _, d := range dis {
		for v, r := range dis {
			for w, dvw := range r[:v] {
				commonDis := (d[v] + d[w] - dvw) / 2
				ans += f[d[v]-commonDis][d[w]-commonDis]
			}
		}
	}
	Fprint(out, ans%mod*pow(int64(n), mod-2)%mod)
}

//func main() { CF1540B(os.Stdin, os.Stdout) }
