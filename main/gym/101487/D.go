package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func runD(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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

	for T := r(); T > 0; T-- {
		n := r()
		a := make([]int, n)
		l := make([]int, n)
		for i := range a {
			a[i] = r()
			if i > 0 && a[i] > a[i-1] {
				l[i] = l[i-1] + 1
			} else {
				l[i] = 1
			}
		}
		r := make([]int, n)
		r[n-1] = 1
		for i := n - 2; i >= 0; i-- {
			if a[i] < a[i+1] {
				r[i] = r[i+1] + 1
			} else {
				r[i] = 1
			}
		}

		ans := 0
		dp := []int{}
		for i, v := range a {
			ans = max(ans, sort.SearchInts(dp, v)+r[i])
			if p := l[i] - 1; p < len(dp) {
				dp[p] = min(dp[p], v)
			} else {
				dp = append(dp, v)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { runD(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
