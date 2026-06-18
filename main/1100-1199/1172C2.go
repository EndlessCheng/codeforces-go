package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1172C2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	w := make([]int, n)
	s := [2]int{}
	for i := range w {
		Fscan(in, &w[i])
		s[a[i]] += w[i]
	}

	f := make([][]int, m+2)
	g := make([][]int, m+2)
	for i := range f {
		f[i] = make([]int, m+2)
		g[i] = make([]int, m+2)
	}
	for j := 0; j <= m; j++ {
		f[j][m-j] = 1
		g[j][m-j] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			k := i - j
			if k > s[0] {
				continue
			}
			s1 := s[1] + j
			s0 := s[0] - k
			sum := s1 + s0
			inv := pow(sum, mod-2)
			f[j][k] = (2*f[j+1][k] + (s1-1)*f[j+1][k] + s0*f[j][k+1]) % mod
			f[j][k] = f[j][k] * inv % mod
			g[j][k] = ((s0-1)*g[j][k+1] + s1*g[j+1][k]) % mod
			g[j][k] = g[j][k] * inv % mod
		}
	}

	for i := 0; i < n; i++ {
		if a[i] > 0 {
			Fprintln(out, f[0][0]*w[i]%mod)
		} else {
			Fprintln(out, g[0][0]*w[i]%mod)
		}
	}
}

//func main() { cf1172C2(bufio.NewReader(os.Stdin), os.Stdout) }
