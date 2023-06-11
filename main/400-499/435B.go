package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF435B(in io.Reader, out io.Writer) {
	var s []byte
	var k int
	Fscan(in, &s, &k)
	n := len(s)
	for i := 0; i < n && k > 0; i++ {
		mxI := i
		for j := i + 1; j < n && j <= i+k; j++ {
			if s[j] > s[mxI] {
				mxI = j
			}
		}
		for ; mxI > i; mxI-- {
			s[mxI], s[mxI-1] = s[mxI-1], s[mxI]
			k--
		}
	}
	Fprintf(out, "%s", s)
}

//func main() { CF435B(os.Stdin, os.Stdout) }
