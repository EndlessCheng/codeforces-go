package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf295D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m int
	Fscan(in, &n, &m)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		s, s2 := 0, 1
		for j := 2; j <= m; j++ {
			s = (s + f[i-1][j]) % mod
			s2 = (s2 + s) % mod
			f[i][j] = s2
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 2; j <= m; j++ {
			ans += (m - j + 1) * (f[i][j] - f[i-1][j]) % mod * f[n-i+1][j] % mod
		}
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf295D(bufio.NewReader(os.Stdin), os.Stdout) }
