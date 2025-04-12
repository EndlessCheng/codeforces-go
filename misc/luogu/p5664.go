package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p5664(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	s := make([]int, n)
	ans := 1
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			s[i] += a[i][j]
		}
		s[i] %= mod
		ans = ans * (s[i] + 1) % mod
	}
	ans--
	f := make([]int, n*2+3)
	for j := 0; j < m; j++ {
		clear(f)
		f[n+1] = 1
		for i, r := range a {
			v := r[j]
			w := s[i] - v
			old := 0
			for d := n - i - 1; d <= n+i+1; d++ {
				t := f[d+1]
				f[d+1] = (f[d+1] + old*v + f[d+2]*w) % mod
				old = t
			}
		}
		for _, v := range f[n+2:] {
			ans -= v
		}
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { p5664(bufio.NewReader(os.Stdin), os.Stdout) }
