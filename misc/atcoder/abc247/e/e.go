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
	var n, x, ans, l, r, p, q, i0 int
	Fscan(in, &n, &r, &l)
	for i := 1; i <= n; i++ {
		Fscan(in, &x)
		if x == l {
			p = i
		}
		if x == r {
			q = i
		}
		if x < l || x > r {
			i0 = i
		}
		x := min(p, q) - i0
		if x > 0 {
			ans += x
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
