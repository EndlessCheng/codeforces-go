package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF484D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 0)
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := make([]int64, n)
	f[1] = int64(abs(a[1] - a[0]))
	for i := 2; i < n; i++ {
		if a[i-2] < a[i-1] == (a[i-1] < a[i]) {
			f[i] = f[i-1] + int64(abs(a[i]-a[i-1]))
		} else {
			f[i] = max(f[i-1], f[i-2]+int64(abs(a[i]-a[i-1])))
		}
	}
	Fprint(out, f[n-1])
}

//func main() { CF484D(os.Stdin, os.Stdout) }
