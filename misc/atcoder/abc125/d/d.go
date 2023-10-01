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
	var n, v, neg, sum int
	mn := int(1e9)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v < 0 {
			neg ^= 1
			v = -v
		}
		sum += v
		mn = min(mn, v)
	}
	Fprint(out, sum-mn*neg*2)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
