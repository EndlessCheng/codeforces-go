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
	var n int
	Fscan(in, &n)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}

	ans := 1
	for i := 0; i < n; i++ {
		fixB := i == 0 || b[i] > b[i-1]
		fixC := i == n-1 || c[i] > c[i+1]
		if fixB && b[i] > c[i] || fixC && b[i] < c[i] {
			Fprint(out, 0)
			return
		}
		if !fixB && !fixC {
			ans = ans * min(b[i], c[i]) % mod
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
