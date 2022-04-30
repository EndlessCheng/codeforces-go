package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1671D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		a := make([]int, n)
		ans := int64(0)
		mi, mx := int(1e9), 0
		for i := range a {
			Fscan(in, &a[i])
			mi = min(mi, a[i])
			mx = max(mx, a[i])
			if i > 0 {
				ans += int64(abs(a[i] - a[i-1]))
			}
		}
		ans += int64(min(min(a[0], a[n-1])-1, (mi-1)*2))
		if mx < x {
			ans += int64(min(x-max(a[0], a[n-1]), (x-mx)*2))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1671D(os.Stdin, os.Stdout) }
