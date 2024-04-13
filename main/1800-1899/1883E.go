package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1883E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, pre int
	for Fscan(in, &T); T > 0; T-- {
		ans, p2 := 0, 0
		for Fscan(in, &n, &pre); n > 1; n-- {
			Fscan(in, &v)
			if v <= pre {
				p2 += bits.Len(uint((pre - 1) / v))
			} else {
				p2 = max(p2-bits.Len(uint(v/pre))+1, 0)
			}
			ans += p2
			pre = v
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1883E(os.Stdin, os.Stdout) }
