package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1811C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, pre, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		Fprint(out, pre, " ")
		for ; n > 2; n-- {
			Fscan(in, &v)
			Fprint(out, min(pre, v), " ")
			pre = v
		}
		Fprintln(out, pre)
	}
}

//func main() { CF1811C(os.Stdin, os.Stdout) }
