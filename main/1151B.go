package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1151B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, di, dj int
	Fscan(in, &n, &m)
	mat := make([][]uint16, n)
	sum := uint16(0)
	for i := range mat {
		mat[i] = make([]uint16, m)
		for j := range mat[i] {
			Fscan(in, &mat[i][j])
			if mat[i][j] != mat[i][0] {
				di, dj = i, j
			}
		}
		sum ^= mat[i][0]
	}
	ans := make([]interface{}, n)
	for i := range ans {
		ans[i] = 1
	}
	if sum == 0 {
		if dj == 0 {
			Fprint(out, "NIE")
			return
		}
		ans[di] = dj + 1
	}
	Fprintln(out, "TAK")
	Fprint(out, ans...)
}

//func main() {
//	CF1151B(os.Stdin, os.Stdout)
//}
