package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1038E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	f := make([][][5][5]int, n+3)
	for i := range f {
		f[i] = make([][5][5]int, n+3)
		for j := range f[i] {
			for x := range f[i][j] {
				for y := range f[i][j][x] {
					f[i][j][x][y] = -1e9
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		var x, y, z int
		Fscan(in, &x, &z, &y)
		f[i][i][x][y] = z
		f[i][i][y][x] = z
	}

	ans := 0
	for i := n; i > 0; i-- {
		for j := i; j <= n; j++ {
			for x := 1; x <= 4; x++ {
				for y := 1; y <= 4; y++ {
					for k := i; k <= j+1; k++ {
						f[i][j][x][y] = max(f[i][j][x][y], f[i][k][x][y], f[k+1][j][x][y])
						for p := 1; p <= 4; p++ {
							f[i][j][x][y] = max(f[i][j][x][y], f[i][k][p][y]+f[k+1][j][x][p], f[i][k][x][p]+f[k+1][j][p][y])
						}
					}
					ans = max(ans, f[i][j][x][y])
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1038E(bufio.NewReader(os.Stdin), os.Stdout) }
