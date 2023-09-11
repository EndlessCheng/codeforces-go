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
	const mod = 1_000_000_007
	var n, m, v, sx, sy int
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sx += (i*2 - n + 1) * v
	}
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		sy += (i*2 - m + 1) * v
	}
	Fprint(out, sx%mod*(sy%mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
