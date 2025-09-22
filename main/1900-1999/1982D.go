package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1982D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n, m, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}

		d := 0
		sum := make([][]int, n+1)
		for i := range sum {
			sum[i] = make([]int, m+1)
		}
		for i, row := range a {
			Fscan(in, &s)
			for j, v := range row {
				c := int(s[j]%2*2) - 1
				d += v * c
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + c
			}
		}
		query := func(r1, c1, r2, c2 int) int {
			return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
		}

		g := 0
		for i := k; i <= n; i++ {
			for j := k; j <= m; j++ {
				g = gcd(g, query(i-k, j-k, i, j))
			}
		}
		if d == 0 || g != 0 && d%g == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1982D(bufio.NewReader(os.Stdin), os.Stdout) }
