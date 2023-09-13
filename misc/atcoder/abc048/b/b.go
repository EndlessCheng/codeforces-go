package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var a, b, x int
	Fscan(in, &a, &b, &x)
	ans := b/x - a/x
	if a%x == 0 {
		ans++
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
