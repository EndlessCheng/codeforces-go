package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m int
	Fscan(in, &n, &m)
	f := make([][11][11][11]int, n+1)
	f[0][m][m][m] = 1
	for i := 1; i <= n; i++ {
		for x := 0; x <= m; x++ {
			for y := x; y <= m; y++ {
				for z := y; z <= m; z++ {
					for j := 0; j < m; j++ {
						val := f[i-1][x][y][z]
						if j <= x {
							f[i][j][y][z] = (f[i][j][y][z] + val) % mod
						} else if j <= y {
							f[i][x][j][z] = (f[i][x][j][z] + val) % mod
						} else if j <= z {
							f[i][x][y][j] = (f[i][x][y][j] + val) % mod
						}
					}
				}
			}
		}
	}

	ans := 0
	for x := 0; x < m; x++ {
		for y := x + 1; y < m; y++ {
			for z := y + 1; z < m; z++ {
				ans += f[n][x][y][z]
			}
		}
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
