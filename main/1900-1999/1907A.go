package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1907A(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		for i := byte('1'); i <= '8'; i++ {
			if i != s[1] {
				Fprintf(out, "%c%c\n", s[0], i)
			}
		}
		for i := byte('a'); i <= 'h'; i++ {
			if i != s[0] {
				Fprintf(out, "%c%c\n", i, s[1])
			}
		}
	}
}

//func main() { cf1907A(os.Stdin, os.Stdout) }
