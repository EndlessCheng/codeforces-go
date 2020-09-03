package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1033C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := make([]byte, n)
	var f func(int) byte
	f = func(p int) (res byte) {
		d := &dp[p]
		if *d > 0 {
			return *d
		}
		defer func() { *d = res }()
		for i := p % a[p]; i < n; i += a[p] {
			if a[i] > a[p] && f(i) == 'B' {
				return 'A'
			}
		}
		return 'B'
	}
	for i := 0; i < n; i++ {
		f(i)
	}
	Fprint(out, string(dp))
}

//func main() { CF1033C(os.Stdin, os.Stdout) }
