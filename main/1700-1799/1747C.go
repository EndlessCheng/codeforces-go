package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1747C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, a1, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a1)
		mn := int(1e9)
		for ; n > 1; n-- {
			Fscan(in, &v)
			mn = min(mn, v)
		}
		if a1 > mn {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { cf1747C(os.Stdin, os.Stdout) }
