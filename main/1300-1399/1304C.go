package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1304C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, t, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &t)
		ok := true
		for ll, rr, pre := t, t, 0; n > 0; n-- {
			Fscan(in, &t, &l, &r)
			ll -= t - pre
			rr += t - pre
			if l > rr || r < ll {
				ok = false
			}
			ll = max(ll, l)
			rr = min(rr, r)
			pre = t
		}
		if ok {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1304C(os.Stdin, os.Stdout) }
