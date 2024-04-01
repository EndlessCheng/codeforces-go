package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1196A(in io.Reader, out io.Writer) {
	var T, a, b, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		Fprintln(out, (a+b+c)/2)
	}
}

//func main() { cf1196A(os.Stdin, os.Stdout) }
