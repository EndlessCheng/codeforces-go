package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	ans := math.MaxInt64
	for a := 1; a <= n; a++ {
		b := (m-1)/a + 1
		if b < a {
			break
		}
		if b <= n {
			ans = min(ans, a*b)
		}
	}
	if ans == math.MaxInt64 {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
