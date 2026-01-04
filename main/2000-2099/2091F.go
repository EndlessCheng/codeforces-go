package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2091F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, m, d int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &d)
		f := make([]int, m+1)
		g := make([]int, m+1)
		for i := range n {
			Fscan(in, &s)
			for j, b := range s {
				if b == '#' {
					f[j+1] = f[j]
				} else if i == 0 {
					f[j+1] = f[j] + 1
				} else {
					f[j+1] = (f[j] + g[min(j+d, m)] - g[max(j-d+1, 0)]) % mod
				}
			}
			for j, b := range s {
				if b == '#' {
					g[j+1] = g[j]
				} else {
					g[j+1] = (g[j] + f[min(j+d+1, m)] - f[max(j-d, 0)]) % mod
				}
			}
		}
		Fprintln(out, (g[m]+mod)%mod)
	}
}

//func main() { cf2091F(bufio.NewReader(os.Stdin), os.Stdout) }
