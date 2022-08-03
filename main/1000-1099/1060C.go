package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1060C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, x, ans int
	Fscan(in, &n, &m)
	f := func(n int) []int {
		mins := make([]int, n)
		for i := range mins {
			mins[i] = 1e9
		}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			for j, s := i, 0; j >= 0; j-- {
				s += a[j]
				mins[i-j] = min(mins[i-j], s)
			}
		}
		return mins
	}
	minA, minB := f(n), f(m)
	Fscan(in, &x)

	for i, v := range minA {
		for j, w := range minB {
			if (i+1)*(j+1) > ans && int64(v)*int64(w) <= int64(x) {
				ans = (i + 1) * (j + 1)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1060C(os.Stdin, os.Stdout) }
