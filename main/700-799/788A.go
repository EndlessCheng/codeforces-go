package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF788A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, pre, v int64
	Fscan(in, &n, &pre)
	a := make([]int64, n-1)
	for i := range a {
		Fscan(in, &v)
		a[i] = abs(v - pre)
		pre = v
	}
	ans := int64(0)
	f := func(a []int64) {
		s := int64(0)
		for i, v := range a {
			if i&1 > 0 {
				v = -v
			}
			s = max(s+v, v)
			ans = max(ans, s)
		}
	}
	f(a)
	f(a[1:])
	Fprint(out, ans)
}

//func main() { CF788A(os.Stdin, os.Stdout) }
