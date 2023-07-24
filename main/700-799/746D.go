package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF746D(in io.Reader, out io.Writer) {
	var n, a, b, k int
	Fscan(in, &n, &k, &a, &b)
	cg, cb := "G", "B"
	if a < b {
		a, b, cg, cb = b, a, cb, cg
	}
	b++
	if (a-1)/b >= k {
		Fprint(out, "NO")
	} else {
		base := cb + strings.Repeat(cg, a/b)
		Fprint(out, strings.Repeat(base, b-a%b)[1:], strings.Repeat(base+cg, a%b))
	}
}

//func main() { CF746D(os.Stdin, os.Stdout) }
