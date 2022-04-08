package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1651C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a ...int64) int64 {
		ans := a[0]
		for _, val := range a[1:] {
			if val < ans {
				ans = val
			}
		}
		return ans
	}
	f := func(a []int64, v int64) int64 {
		res := int64(1e9)
		for _, w := range a {
			res = min(res, abs(w-v))
		}
		return res
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int64, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		Fprintln(out, min(
			abs(a[0]-b[0])+abs(a[n-1]-b[n-1]),
			abs(a[0]-b[n-1])+abs(a[n-1]-b[0]),
			abs(a[0]-b[0])+f(a, b[n-1])+f(b, a[n-1]),
			abs(a[0]-b[n-1])+f(a, b[0])+f(b, a[n-1]),
			abs(a[n-1]-b[0])+f(a, b[n-1])+f(b, a[0]),
			abs(a[n-1]-b[n-1])+f(a, b[0])+f(b, a[0]),
			f(a, b[0])+f(a, b[n-1])+f(b, a[0])+f(b, a[n-1]),
		))
	}
}

//func main() { CF1651C(os.Stdin, os.Stdout) }
