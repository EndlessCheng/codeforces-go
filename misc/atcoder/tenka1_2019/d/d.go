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
	const mod = 998244353
	var n, v, s, dup int
	var f, g [90001]int
	f[0] = 3
	g[0] = 3
	pow3 := 1
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		s += v
		for j := s; j >= 0; j-- {
			f[j] = (f[j]*2 + f[max(j-v, 0)]) % mod
			if j >= v {
				g[j] = (g[j] + g[j-v]) % mod
			}
		}
		pow3 = pow3 * 3 % mod
	}
	if s%2 == 0 {
		dup = g[s/2]
	}
	ans := pow3 - (f[(s+1)/2] - dup)
	Fprint(out, (ans%mod+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
