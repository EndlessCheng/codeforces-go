package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1254B2(_r io.Reader, out io.Writer) {
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	n := r()
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + int64(r())
	}
	x := sum[n]
	if x < 2 {
		Fprint(out, -1)
		return
	}
	a := []int64{}
	for i := int64(2); i*i <= x; i++ {
		k := 0
		for ; x%i == 0; x /= i {
			k++
		}
		if k > 0 {
			a = append(a, i)
		}
	}
	if x > 1 {
		a = append(a, x)
	}
	ans := int64(1e18)
	for _, p := range a {
		s := int64(0)
		for _, v := range sum {
			s += min(v%p, p-v%p)
		}
		ans = min(ans, s)
	}
	Fprint(out, ans)
}

//func main() { CF1254B2(os.Stdin, os.Stdout) }
