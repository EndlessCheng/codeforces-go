package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	if n == 0 {
		Fprint(out, "0")
		return
	}
	ans := ""
	for ; n != 0; n = -(n >> 1) {
		ans = string(byte('0'+n&1)) + ans
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
