package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1519D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sum := int64(0)
	b := make([]int64, n)
	for i := range b {
		Fscan(in, &b[i])
		sum += a[i] * b[i]
	}

	ans := sum
	for i := 0; i < n; i++ {
		cur := sum
		for l, r := i-1, i+1; l >= 0 && r < n; l-- {
			cur -= (a[l] - a[r]) * (b[l] - b[r])
			ans = max(ans, cur)
			r++
		}
		cur = sum
		for l, r := i-1, i; l >= 0 && r < n; l-- {
			cur -= (a[l] - a[r]) * (b[l] - b[r])
			ans = max(ans, cur)
			r++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1519D(os.Stdin, os.Stdout) }
