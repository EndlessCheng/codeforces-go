package _00_299

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF283B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 2; i <= n; i++ {
		Fscan(in, &a[i])
	}
	dp := make([][2]int64, n+1)
	var f func(int, int8) int64
	f = func(p int, d int8) (res int64) {
		if p > 1 {
			dv := &dp[p][d]
			if *dv != 0 {
				return *dv
			}
			*dv = -1
			defer func() { *dv = res }()
		}
		res = int64(a[p])
		if d == 0 {
			p -= a[p]
		} else {
			p += a[p]
		}
		if p == 1 {
			return -1
		}
		if p <= 0 || p > n {
			return
		}
		v := f(p, d^1)
		if v == -1 {
			return -1
		}
		return res + v
	}
	for a[1] = 1; a[1] < n; a[1]++ {
		Fprintln(out, f(1, 1))
	}
}

//func main() { CF283B(os.Stdin, os.Stdout) }
