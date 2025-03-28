package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf56D(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := range f[0] {
		f[0][j] = j
	}
	for i, v := range s {
		f[i+1][0] = i + 1
		for j, w := range t {
			if v == w {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
			}
		}
	}
	Fprintln(out, f[n][m])

	// 注：也可以迭代，从右往左操作，比如可以这样写 https://codeforces.com/problemset/submission/56/7486578
	var output func(int, int)
	output = func(i, j int) {
		switch {
		case i < 0 && j < 0:
		case i < 0 || j >= 0 && f[i+1][j+1] == f[i+1][j]+1:
			output(i, j-1)
			Fprintf(out, "INSERT %d %c\n", j+1, t[j])
		case j < 0 || f[i+1][j+1] == f[i][j+1]+1:
			output(i-1, j)
			Fprintf(out, "DELETE %d\n", j+2)
		default:
			output(i-1, j-1)
			if s[i] != t[j] {
				Fprintf(out, "REPLACE %d %c\n", j+1, t[j])
			}
		}
	}
	output(n-1, m-1)
}

//func main() { cf56D(bufio.NewReader(os.Stdin), os.Stdout) }
