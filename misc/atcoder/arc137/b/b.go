package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, inc, maxInc, dec, maxDec int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		inc = max(inc, 0) + 1 - v*2
		maxInc = max(maxInc, inc)
		dec = max(dec, 0) + v*2 - 1
		maxDec = max(maxDec, dec)
	}
	Fprint(out, maxInc+maxDec+1)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
