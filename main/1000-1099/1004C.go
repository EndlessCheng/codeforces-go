package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1004C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	var l, r [1e5 + 1]int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		if l[a[i]] == 0 {
			l[a[i]] = i
			k++
		}
		r[a[i]] = i
	}

	ans := int64(0)
	for i := 1; i <= n; i++ {
		if i == r[a[i]] {
			k--
		}
		if i == l[a[i]] {
			ans += int64(k)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1004C(os.Stdin, os.Stdout) }
