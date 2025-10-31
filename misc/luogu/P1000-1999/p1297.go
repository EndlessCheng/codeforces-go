package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1297(in io.Reader, out io.Writer) {
	var n, a, b, c, v int
	Fscan(in, &n, &a, &b, &c, &v)
	a0 := v%c + 1
	ans := 0.
	for i := 1; i < n; i++ {
		nxt := (v*a + b) % 100000001
		ans += 1 / float64(max(v%c, nxt%c)+1)
		v = nxt
	}
	ans += 1 / float64(max(v%c+1, a0))
	Fprintf(out, "%.3f", ans)
}

//func main() { p1297(os.Stdin, os.Stdout) }
