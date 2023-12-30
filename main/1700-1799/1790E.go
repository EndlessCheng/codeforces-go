package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1790E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, xor int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &xor)
		and := xor >> 1
		if xor&1 > 0 || xor&and > 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, xor|and, and)
		}
	}
}

//func main() { cf1790E(os.Stdin, os.Stdout) }
