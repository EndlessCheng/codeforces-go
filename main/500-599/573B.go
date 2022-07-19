package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF573B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] = min(a[i], a[i-1]+1)
	}
	for i := n; i > 0; i-- {
		a[i] = min(a[i], a[i+1]+1)
		ans = max(ans, a[i])
	}
	Fprint(out, ans)
}

//func main() { CF573B(os.Stdin, os.Stdout) }
