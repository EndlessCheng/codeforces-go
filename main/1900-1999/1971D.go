package main

import (
	"bytes"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf1971D(in io.Reader, out io.Writer) {
	T, s := 0, []byte{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		s = slices.Compact(s)
		if bytes.Contains(s, []byte("01")) {
			Fprintln(out, len(s)-1)
		} else {
			Fprintln(out, len(s))
		}
	}
}

//func main() { cf1971D(bufio.NewReader(os.Stdin), os.Stdout) }
