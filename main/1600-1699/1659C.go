package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1659C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b)
		x := make([]int, n)
		suf := 0
		for i := range x {
			Fscan(in, &x[i])
			suf += x[i]
		}
		ans := b * suf
		for i, v := range x {
			suf -= v
			ans = min(ans, a*v+b*(v+suf-v*(n-1-i)))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1659C(os.Stdin, os.Stdout) }
