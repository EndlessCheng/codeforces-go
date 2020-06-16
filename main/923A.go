package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF923A(in io.Reader, out io.Writer) {
	const mx = 1e6
	f := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if f[i] == 0 {
			for j := i; j <= mx; j += i {
				f[j] = i
			}
		}
	}
	n := 0
	Fscan(in, &n)
	ans := n
	for i := n - f[n] + 1; i <= n; i++ {
		if f[i] < i && i-f[i]+1 < ans {
			ans = i - f[i] + 1
		}
	}
	Fprint(out, ans)
}

//func main() { CF923A(os.Stdin, os.Stdout) }
