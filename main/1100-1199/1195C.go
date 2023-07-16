package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1195C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	var n, w, f0, f1, f2 int64
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for _, v := range a {
		Fscan(in, &w)
		f0, f1, f2 = max(max(f0, f1), f2), max(f0, f2)+int64(v), max(f0, f1)+w
	}
	Fprint(out, max(max(f0, f1), f2))
}

//func main() { CF1195C(os.Stdin, os.Stdout) }
