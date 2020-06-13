package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1366C(_r io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var t, n, m, v int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		a := make([][2]int, n+m-1)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				Fscan(in, &v)
				a[i+j][v]++
			}
		}
		ans := 0
		for i, j := 0, n+m-2; i < j; i++ {
			ans += min(a[i][0]+a[j][0], a[i][1]+a[j][1])
			j--
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1366C(os.Stdin, os.Stdout) }
