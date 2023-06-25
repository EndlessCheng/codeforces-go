package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF987C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e9
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	
	var n int
	Fscan(in, &n)
	a := make([]struct{ x, y int }, n)
	for i := range a { Fscan(in, &a[i].x) }
	for i := range a { Fscan(in, &a[i].y) }
	ans := inf
	for j := 1; j < n-1; j++ {
		l := inf
		for _, p := range a[:j] {
			if p.x < a[j].x {
				l = min(l, p.y)
			}
		}
		r := inf
		for _, p := range a[j+1:] {
			if p.x > a[j].x {
				r = min(r, p.y)
			}
		}
		ans = min(ans, l+a[j].y+r)
	}
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF987C(os.Stdin, os.Stdout) }
