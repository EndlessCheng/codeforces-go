package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2113C(in io.Reader, out io.Writer) {
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([][]byte, n)
		sum := make([][]int, n+1)
		for i := range sum {
			sum[i] = make([]int, m+1)
		}
		for i := range a {
			Fscan(in, &a[i])
			for j, v := range a[i] {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
				if v == 'g' {
					sum[i+1][j+1]++
				}
			}
		}
		query := func(r1, c1, r2, c2 int) int {
			return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
		}

		mn := int(1e9)
		for i, row := range a {
			for j, b := range row {
				if b == '.' {
					mn = min(mn, query(max(i-k+1, 0), max(j-k+1, 0), min(i+k, n), min(j+k, m)))
				}
			}
		}
		if mn < 1e9 {
			Fprintln(out, sum[n][m]-mn)
		} else {
			Fprintln(out, 0)
		}
	}
}

//func main() { cf2113C(bufio.NewReader(os.Stdin), os.Stdout) }
