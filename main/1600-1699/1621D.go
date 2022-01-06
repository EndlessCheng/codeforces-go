package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1621D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := int64(0)
		a := make([][]int64, n*2)
		for i := range a {
			a[i] = make([]int64, n*2)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				if i >= n && j >= n {
					ans += a[i][j]
				}
			}
		}
		min := int64(1e9)
		for _, v := range []int64{a[0][n], a[0][n*2-1], a[n-1][n], a[n-1][n*2-1], a[n][0], a[n][n-1], a[n*2-1][0], a[n*2-1][n-1]} {
			if v < min {
				min = v
			}
		}
		Fprintln(out, ans+min)
	}
}

//func main() { CF1621D(os.Stdin, os.Stdout) }
