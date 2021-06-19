package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF835C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 101

	var n, q, c, x, y, s, t, x2, y2 int
	Fscan(in, &n, &q, &c)
	c++
	sum := make([][mx][mx]int, c)
	for ; n > 0; n-- {
		Fscan(in, &x, &y, &s)
		for t := range sum {
			sum[t][x][y] += (s + t) % c
		}
	}
	for t := range sum {
		for i := 1; i < mx; i++ {
			for j := 1; j < mx; j++ {
				sum[t][i][j] += sum[t][i-1][j] + sum[t][i][j-1] - sum[t][i-1][j-1]
			}
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &t, &x, &y, &x2, &y2)
		t %= c
		Fprintln(out, sum[t][x2][y2]-sum[t][x-1][y2]-sum[t][x2][y-1]+sum[t][x-1][y-1])
	}
}

//func main() { CF835C(os.Stdin, os.Stdout) }
