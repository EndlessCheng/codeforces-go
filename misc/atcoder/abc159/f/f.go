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
	var n, s, v, ans int
	Fscan(in, &n, &s)
	f := make([]int, s+1)
	for ; n > 0; n-- {
		Fscan(in, &v)
		f[0]++
		for j := s; j >= v; j-- {
			f[j] = (f[j] + f[j-v]) % mod
		}
		ans += f[s]
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
