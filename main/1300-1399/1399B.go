package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1399B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type pair struct{ x, y int }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var miA, miB int = 1e9, 1e9
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x)
			miA = min(miA, a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
			miB = min(miB, a[i].y)
		}
		ans := int64(0)
		for _, p := range a {
			ans += int64(max(p.x-miA, p.y-miB))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1399B(os.Stdin, os.Stdout) }
