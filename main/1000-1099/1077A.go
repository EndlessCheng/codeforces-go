package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1077A(in io.Reader, out io.Writer) {
	var T, a, b, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &k)
		Fprintln(out, k/2*(a-b)+k%2*a)
	}
}

//func main() { cf1077A(os.Stdin, os.Stdout) }
