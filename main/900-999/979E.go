package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf979E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m, x int
	Fscan(in, &n, &m)
	f := make([][2][2][2]int, n+1)
	f[0][0][0][0] = 1
	s := mod/2 + 1
	for i := range n {
		Fscan(in, &x)
		for j := range 2 {
			for k := range 2 {
				for l := range 2 {
					if x^1 != 0 {
						if l == 1 {
							f[i+1][j][k][1] = (f[i+1][j][k][1] + f[i][j][k][1]*s) % mod
							f[i+1][j^1][1][1] = (f[i+1][j^1][1][1] + f[i][j][k][1]*s) % mod
						} else {
							f[i+1][j^1][1][0] = (f[i+1][j^1][1][0] + f[i][j][k][0]*s*2) % mod
						}
					}
					if x != 0 {
						if k == 1 {
							f[i+1][j][1][l] = (f[i+1][j][1][l] + f[i][j][1][l]*s) % mod
							f[i+1][j^1][1][1] = (f[i+1][j^1][1][1] + f[i][j][1][l]*s) % mod
						} else {
							f[i+1][j^1][0][1] = (f[i+1][j^1][0][1] + f[i][j][0][l]*s*2) % mod
						}
					}
				}
			}
		}
		s = s * 2 % mod
	}
	Fprint(out, (f[n][m][0][0]+f[n][m][0][1]+f[n][m][1][0]+f[n][m][1][1])%mod)
}

//func main() { cf979E(os.Stdin, os.Stdout) }
