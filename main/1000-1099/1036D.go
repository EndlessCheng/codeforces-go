package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1036D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, ans int
	Fscan(in, &n)
	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		a[i] = a[i-1] + int64(v)
	}
	Fscan(in, &m)
	b := make([]int64, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &v)
		b[i] = b[i-1] + int64(v)
	}
	if a[n] != b[m] {
		Fprint(out, -1)
		return
	}
	for i, j := 1, 1; i <= n && j <= m; {
		if a[i] == b[j] {
			ans++
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1036D(os.Stdin, os.Stdout) }
