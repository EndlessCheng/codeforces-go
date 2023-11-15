package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1169B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n, &n)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}

	var f func(int, int) bool
	f = func(x, y int) bool {
		for _, p := range a {
			xx, yy := p.x, p.y
			if xx == x || xx == y || yy == x || yy == y {
				continue
			}
			if y > 0 {
				return false
			}
			return f(x, xx) || f(x, yy)
		}
		return true
	}
	if f(a[0].x, 0) || f(a[0].y, 0) {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

//func main() { CF1169B(os.Stdin, os.Stdout) }
