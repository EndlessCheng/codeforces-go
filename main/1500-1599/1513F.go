package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1513F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, d int
	Fscan(in, &n)
	type pair struct{ x, y, t int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x)
	}
	ans := int64(0)
	for i := range a {
		Fscan(in, &a[i].y)
		if a[i].x > a[i].y {
			a[i] = pair{a[i].y, a[i].x, 1}
		}
		ans += int64(a[i].y - a[i].x)
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x < b.x || a.x == b.x && a.y > b.y })
	mx := [2]int{}
	for _, p := range a {
		mx[p.t] = max(mx[p.t], p.y)
		d = max(d, min(p.y, mx[p.t^1])-p.x)
	}
	Fprint(out, ans-int64(d*2))
}

//func main() { CF1513F(os.Stdin, os.Stdout) }
