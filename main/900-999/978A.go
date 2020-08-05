package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF978A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := []int{}
	vis := [1001]bool{}
	for i := n - 1; i >= 0; i-- {
		if !vis[a[i]] {
			vis[a[i]] = true
			ans = append(ans, a[i])
		}
	}
	Fprintln(out, len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, ans[i], " ")
	}
}

//func main() { CF978A(os.Stdin, os.Stdout) }
