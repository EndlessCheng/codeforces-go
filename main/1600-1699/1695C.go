package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1695C(_r io.Reader, _w io.Writer) {
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

	var T, n, m int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		mn := make([]int, m+1)
		mx := make([]int, m+1)
		for i := 1; i <= m; i++ {
			mn[i], mx[i] = 1e9, -1e9
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				Fscan(in, &s)
				x := 3 - len(s)*2
				mn[j+1] = min(mn[j+1], mn[j]) + x
				mx[j+1] = max(mx[j+1], mx[j]) + x
			}
			mn[0], mx[0] = 1e9, -1e9
		}
		if (n+m)%2 == 0 || mn[m] > 0 || mx[m] < 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1695C(os.Stdin, os.Stdout) }
