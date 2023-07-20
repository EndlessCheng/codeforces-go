package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1772D(_r io.Reader, _w io.Writer) {
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

	var T, n, pre, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		l, r := 0, int(1e9)
		for ; n > 1; n-- {
			Fscan(in, &v)
			if pre > v {
				l = max(l, (pre+v+1)/2)
			} else if pre < v {
				r = min(r, (pre+v)/2)
			}
			pre = v
		}
		if l <= r {
			Fprintln(out, l)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1772D(os.Stdin, os.Stdout) }
