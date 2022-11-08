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
	var n, v, s, a, b int
	Fscan(in, &n, &s)
	for i := 2; i <= n; i++ {
		Fscan(in, &v)
		if i&1 == 0 {
			a, b = b, max(s, a+v)
		} else {
			s += v
			a, b = b, max(b, a+v)
		}
	}
	Fprint(out, b)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
