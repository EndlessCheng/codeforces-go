package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1420C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	b := func(b bool) int64 {
		if b {
			return 1
		}
		return 0
	}

	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int64, n+2)
		f := func(x int) int64 {
			if d := a[x] - a[x-1]; d > 0 {
				return d
			}
			return 0
		}
		ans := int64(0)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			ans += f(i)
		}
		Fprintln(out, ans)
		for ; q > 0; q-- {
			Fscan(in, &l, &r)
			ans -= f(l) + f(l+1) + b(l+1 < r)*f(r) + b(l < r)*f(r+1)
			a[l], a[r] = a[r], a[l]
			ans += f(l) + f(l+1) + b(l+1 < r)*f(r) + b(l < r)*f(r+1)
			Fprintln(out, ans)
		}
	}
}

//func main() { CF1420C2(os.Stdin, os.Stdout) }
