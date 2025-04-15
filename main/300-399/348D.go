package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf348D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m int
	var s []byte
	Fscan(in, &n, &m)
	f := make([][]int, n+1)
	g := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
		g[i] = make([]int, m+1)
	}
	f[0][2] = 1
	g[2][0] = 1
	for i := range n {
		Fscan(in, &s)
		for j, b := range s {
			if b == '.' {
				f[i+1][j+1] = (f[i+1][j] + f[i][j+1]) % mod
				g[i+1][j+1] = (g[i+1][j] + g[i][j+1]) % mod
			}
		}
	}
	Fprint(out, ((f[n-1][m]*g[n][m-1]-f[n][m-1]*g[n-1][m])%mod+mod)%mod)
}

//func main() { cf348D(bufio.NewReader(os.Stdin), os.Stdout) }
