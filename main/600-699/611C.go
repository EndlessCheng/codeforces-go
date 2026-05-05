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
	sr := make([][]int, n+1)
	sc := make([][]int, n+1)
	for i := range sr {
		sr[i] = make([]int, m+1)
		sc[i] = make([]int, m+1)
	}
	for i, row := range a {
		for j, b := range row {
			sr[i+1][j+1] = sr[i+1][j] + sr[i][j+1] - sr[i][j]
			sc[i+1][j+1] = sc[i+1][j] + sc[i][j+1] - sc[i][j]
			if b == '.' {
				if j < m-1 && row[j+1] == '.' {
					sr[i+1][j+1]++
				}
				if i < n-1 && a[i+1][j] == '.' {
					sc[i+1][j+1]++
				}
			}
		}
	}
	query := func(s [][]int, r1, c1, r2, c2 int) int {
		return s[r2][c2] - s[r2][c1] - s[r1][c2] + s[r1][c1]
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &r1, &c1, &r2, &c2)
		Fprintln(out, query(sr, r1-1, c1-1, r2, c2-1)+query(sc, r1-1, c1-1, r2-1, c2))
	}
}

//func main() { cf611C(bufio.NewReader(os.Stdin), os.Stdout) }
