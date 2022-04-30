package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1622C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var T, n int
	var k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		tot := int64(0)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
			tot += a[i]
		}
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		if tot <= k {
			Fprintln(out, 0)
			continue
		}
		ans := tot - k
		for i, s := n-1, a[0]-ans; i > 0; i-- {
			s += a[i]
			var x int64
			y := int64(n - i + 1)
			if s >= 0 {
				x = s / y
			} else {
				x = (s - y + 1) / y
			}
			x = min(x, a[0])
			ans = min(ans, a[0]-x+y-1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1622C(os.Stdin, os.Stdout) }
