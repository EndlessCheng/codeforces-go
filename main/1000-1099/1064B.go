package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1064B(in io.Reader, out io.Writer) {
	var T, n uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, 1<<bits.OnesCount(n))
	}
}

//func main() { cf1064B(os.Stdin, os.Stdout) }
