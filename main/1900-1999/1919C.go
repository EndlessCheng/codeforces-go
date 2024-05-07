package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1919C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		for x, y := n, n; n > 0; n-- {
			Fscan(in, &v)
			if v > x && v > y {
				ans++
				if x < y {
					x = v
				} else {
					y = v
				}
			} else if v > y || v <= x && x < y {
				x = v
			} else {
				y = v
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1919C(os.Stdin, os.Stdout) }
