package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1272D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := 1
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			suf[i] = suf[i+1] + 1
		} else {
			suf[i] = 1
		}
		ans = max(ans, suf[i])
	}
	for i, pre := 1, 1; i < n; i++ {
		if i+1 < n && a[i-1] < a[i+1] {
			ans = max(ans, pre+suf[i+1])
		}
		if a[i] > a[i-1] {
			pre++
			ans = max(ans, pre)
		} else {
			pre = 1
		}
	}
	Fprint(out, ans)
}

//func main() { CF1272D(os.Stdin, os.Stdout) }
