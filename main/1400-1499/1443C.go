package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1443C(_r io.Reader, _w io.Writer) {
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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y int }, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		s := 0
		for i := range a {
			Fscan(in, &a[i].y)
			s += a[i].y
		}
		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
		ans := s
		for _, p := range a {
			s -= p.y
			ans = min(ans, max(p.x, s))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1443C(os.Stdin, os.Stdout) }
