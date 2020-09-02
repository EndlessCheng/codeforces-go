package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1372D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var sum, cur, ans int64
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		sum += a[i]
	}
	for i := 0; i < n; i += 2 {
		cur += a[i]
	}
	for _, v := range a {
		if cur > ans {
			ans = cur
		}
		cur = sum - cur + v
	}
	Fprint(out, ans)
}

//func main() { CF1372D(os.Stdin, os.Stdout) }
