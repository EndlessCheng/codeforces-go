package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1615F(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	const mod = 1_000_000_007
	var T, n int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		for i := 1; i < n; i += 2 {
			s[i] ^= 1
			t[i] ^= 1
		}
		f := make([][]int, n+1)
		g := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n*2+1)
			g[i] = make([]int, n*2+1)
		}
		f[0][n] = 1
		for i := 1; i <= n; i++ {
			for j := n - i; j <= n+i; j++ {
				for x := 0; x <= 1; x++ {
					for y := 0; y <= 1; y++ {
						if x^1 != int(s[i-1]-'0') && y^1 != int(t[i-1]-'0') {
							nj := j + x - y
							if nj < 0 || nj > n*2 {
								continue
							}
							f[i][nj] = (f[i][nj] + f[i-1][j]) % mod
							g[i][nj] = (g[i][nj] + g[i-1][j] + abs(j-n)*f[i-1][j]) % mod
						}
					}
				}
			}
		}
		Fprintln(out, g[n][n])
	}
}

//func main() { cf1615F(bufio.NewReader(os.Stdin), os.Stdout) }
