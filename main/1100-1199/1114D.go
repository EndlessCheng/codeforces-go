package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1114D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, k int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	for _, w := range a[1:] {
		if a[k] != w {
			k++
			a[k] = w
		}
	}

	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for d := 1; d <= k; d++ {
		for l := 0; l+d <= k; l++ {
			r := l + d
			if a[l] == a[r] {
				f[l][r] = f[l+1][r-1] + 1
			} else {
				f[l][r] = min(f[l+1][r], f[l][r-1]) + 1
			}
		}
	}
	Fprint(out, f[0][k])
}

//func main() { CF1114D(os.Stdin, os.Stdout) }
