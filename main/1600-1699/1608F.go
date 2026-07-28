package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1608F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, K int
	Fscan(in, &n, &K)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
	}
	f := [2][][]int{}
	for i := range f {
		f[i] = make([][]int, n+2)
		for j := range f[i] {
			f[i][j] = make([]int, n+2)
		}
	}

	cur := 0
	ans := 0
	l, r := 0, 0
	f[0][0][0] = 1

	add := func(x int) int {
		if x < 0 {
			return x + mod
		}
		return x
	}

	for i := 1; i <= n; i++ {
		r = min(i, b[i]+K)
		for j := 0; j <= n; j++ {
			for k := 0; k <= n; k++ {
				f[cur^1][j][k] = 0
			}
		}
		for j := l + 1; j <= r; j++ {
			for k := 0; k <= i-j; k++ {
				f[cur^1][j][k] = f[cur^1][j-1][k+1] * (k + 1) % mod
				if j-1 <= b[i-1]+K {
					f[cur^1][j][k] = add(f[cur^1][j][k] + f[cur][j-1][k] - mod)
				}
			}
		}
		l = max(l, b[i]-K)
		r = min(r, min(b[i-1]+K, i-1))
		for j := l; j <= r; j++ {
			for k := 0; k <= i-j; k++ {
				if k < i-j {
					f[cur^1][j][k] = (f[cur^1][j][k] + f[cur][j][k]*(j+k)) % mod
				}
				if k > 0 {
					f[cur^1][j][k] = add(f[cur^1][j][k] + f[cur][j][k-1] - mod)
				}
			}
		}
		cur ^= 1
	}

	r = min(n, b[n]+K)
	for i := l; i <= r; i++ {
		s := 1
		for j := 0; j <= n-i; j++ {
			ans = (ans + f[cur][i][j]*s) % mod
			s = s * (n - i - j) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { cf1608F(bufio.NewReader(os.Stdin), os.Stdout) }
