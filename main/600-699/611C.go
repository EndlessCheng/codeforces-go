package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf611C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, r1, c1, r2, c2 int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sum := make([][]int, n+1)
	sumR := make([][]int, n+1)
	for i := range sum {
		sum[i] = make([]int, m+1)
		sumR[i] = make([]int, m+1)
	}
	sumC := make([][]int, m)
	for i := range sumC {
		sumC[i] = make([]int, n+1)
	}
	for i, row := range a {
		for j, b := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
			sumR[i][j+1] = sumR[i][j]
			sumC[j][i+1] = sumC[j][i]
			if b == '.' {
				if j < m-1 && row[j+1] == '.' {
					sum[i+1][j+1]++
					sumR[i][j+1]++
				}
				if i < n-1 && a[i+1][j] == '.' {
					sum[i+1][j+1]++
					sumC[j][i+1]++
				}
			}
		}
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &r1, &c1, &r2, &c2)
		r1--
		c1--
		r2--
		c2--
		Fprintln(out, sum[r2][c2]-sum[r2][c1]-sum[r1][c2]+sum[r1][c1]+sumR[r2][c2]-sumR[r2][c1]+sumC[c2][r2]-sumC[c2][r1])
	}
}

//func main() { cf611C(bufio.NewReader(os.Stdin), os.Stdout) }
