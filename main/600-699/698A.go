package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF698A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	const inf int = 1e9
	var n, v, f0, f1, f2 int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		g1, g2 := inf, inf
		if v&1 > 0 {
			g1 = min(f0, f2)
		}
		if v>>1 > 0 {
			g2 = min(f0, f1)
		}
		f0 = min(min(f0, f1), f2) + 1
		f1, f2 = g1, g2
	}
	Fprint(out, min(min(f0, f1), f2))
}

//func main() { CF698A(os.Stdin, os.Stdout) }
