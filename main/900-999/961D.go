package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF961D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	if n < 5 {
		Fprint(out, "YES")
		return
	}
	type pair struct{ x, y int64 }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	line := func(a, b, c pair) bool { return (b.x-a.x)*(c.y-a.y) == (b.y-a.y)*(c.x-a.x) }
	f := func(p0, p1 pair) bool {
		other := []pair{}
		for _, p := range a {
			if !line(p0, p1, p) {
				if len(other) < 2 {
					other = append(other, p)
				} else if !line(other[0], other[1], p) {
					return false
				}
			}
		}
		return true
	}
	if f(a[0], a[1]) || f(a[0], a[2]) || f(a[1], a[2]) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF961D(os.Stdin, os.Stdout) }
