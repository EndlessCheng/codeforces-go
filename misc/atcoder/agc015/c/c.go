package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, r1, c1, r2, c2 int
	Fscan(in, &n, &m, &q)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sumV := make([][]int, n+1)
	sumE := make([][]int, n+1)
	rowE := make([][]int, n+1)
	colE := make([][]int, n+1)
	for i := range sumV {
		sumV[i] = make([]int, m+1)
		sumE[i] = make([]int, m)
		rowE[i] = make([]int, m+1)
		colE[i] = make([]int, m)
	}
	for i, row := range a {
		for j, v := range row {
			sumV[i+1][j+1] = sumV[i+1][j] + sumV[i][j+1] - sumV[i][j] + int(v-'0')
			if i < n-1 && j < m-1 {
				sumE[i+1][j+1] = sumE[i+1][j] + sumE[i][j+1] - sumE[i][j]
				if v == '1' {
					sumE[i+1][j+1] += int(a[i][j+1]-'0') + int(a[i+1][j]-'0')
				}
			}
			rowE[i][j+1] = rowE[i][j]
			if j < m-1 && v == '1' {
				rowE[i][j+1] += int(row[j+1] - '0')
			}
			colE[i+1][j] = colE[i][j]
			if i < n-1 && v == '1' {
				colE[i+1][j] += int(a[i+1][j] - '0')
			}
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &r1, &c1, &r2, &c2)
		r1--; c1--
		v := sumV[r2][c2] - sumV[r2][c1] - sumV[r1][c2] + sumV[r1][c1]
		r2--; c2--
		e := sumE[r2][c2] - sumE[r2][c1] - sumE[r1][c2] + sumE[r1][c1] +
			 rowE[r2][c2] - rowE[r2][c1] +
			 colE[r2][c2] - colE[r1][c2]
		Fprintln(out, v-e)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
