package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1237D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n, 3*n)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(append(a, a...), a...)

	q := make([]int, 3*n)
	l, r, j := 0, 0, 0
	for i := range a[:n] {
		for ; j < 3*n && (l == r || 2*a[j] >= a[q[l]]); j++ {
			for ; l < r && a[q[r-1]] <= a[j]; r-- {
			}
			q[r] = j
			r++
		}
		ans := j - i
		if ans > 2*n {
			ans = -1
		}
		Fprint(out, ans, " ")
		if l < r && q[l] == i {
			l++
		}
	}
}

func main() { CF1237D(os.Stdin, os.Stdout) }
