package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf513E2(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	f := make([][][4]int, n+1)
	for i := range f {
		f[i] = make([][4]int, m+1)
		for j := range f[i] {
			for k := range f[i][j] {
				f[i][j][k] = -1e18
			}
		}
		for j := range 4 {
			f[i][0][j] = 0
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			mul := 1
			if j != 1 && j != m {
				mul = 2
			}
			f[i][j][0] = max(f[i-1][j][0], f[i-1][j-1][3]) - mul*a[i]
			f[i][j][1] = max(f[i-1][j][1], f[i][j][0])
			f[i][j][2] = max(f[i-1][j][2], f[i-1][j-1][1]) + mul*a[i]
			f[i][j][3] = max(f[i-1][j][3], f[i][j][2])
			if mul == 2 {
				f[i][j][1] = max(f[i][j][1], f[i-1][j-1][1])
				f[i][j][3] = max(f[i][j][3], f[i-1][j-1][3])
			}
		}
	}

	Fprint(out, max(f[n][m][1], f[n][m][3]))
}

//func main() { cf513E2(bufio.NewReader(os.Stdin), os.Stdout) }
