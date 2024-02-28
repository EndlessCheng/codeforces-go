package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1705C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, c, q, r, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c, &q, &s)
		op := make([]struct{ sz, l int }, c)
		op[0].sz = len(s)
		for i := range op {
			if i > 0 {
				op[i].sz = op[i-1].sz + r - op[i-1].l
			}
			Fscan(in, &op[i].l, &r)
			op[i].l--
		}
		for ; q > 0; q-- {
			Fscan(in, &k)
			k--
			for i := c - 1; i >= 0; i-- {
				if k >= op[i].sz {
					k -= op[i].sz - op[i].l
				}
			}
			Fprintln(out, s[k:k+1])
		}
	}
}

//func main() { cf1705C(os.Stdin, os.Stdout) }
