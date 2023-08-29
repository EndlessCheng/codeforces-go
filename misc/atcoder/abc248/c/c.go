package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m, k int
	Fscan(in, &n, &m, &k)
	suf := make([]int, k+2)
	suf[0] = 1
	for ; n > 0; n-- {
		for j := k; j >= 0; j-- {
			f := suf[max(j-m, 0)] - suf[j]
			suf[j] = (suf[j+1] + f) % mod
		}
	}
	Fprint(out, (suf[0]+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
