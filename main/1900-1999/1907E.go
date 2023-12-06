package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1907E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := 1
		for _, b := range s {
			v := int(b & 15)
			ans *= (v + 1) * (v + 2) / 2
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1907E(os.Stdin, os.Stdout) }
