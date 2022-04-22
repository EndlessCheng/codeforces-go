package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1667A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	if n == 2 {
		Fprint(out, 1)
		return
	}
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := int64(1e18)
	for i := 1; i < n-1; i++ {
		c := int64(2)
		for j, pre := i-2, a[i-1]; j >= 0; j-- {
			v := pre/a[j] + 1
			c += v
			pre = a[j] * v
		}
		for j, pre := i+2, a[i+1]; j < n; j++ {
			v := pre/a[j] + 1
			c += v
			pre = a[j] * v
		}
		if c < ans {
			ans = c
		}
	}
	Fprintln(out, ans)
}

//func main() { CF1667A(os.Stdin, os.Stdout) }
