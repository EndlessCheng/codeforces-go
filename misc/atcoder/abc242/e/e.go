package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans, le := 0, -1
		for i, p26 := (n-1)/2, 1; i >= 0; i-- {
			ans += int(s[i]-'A') * p26
			p26 = p26 * 26 % mod
			if le < 0 && s[i] != s[n-1-i] {
				if s[i] > s[n-1-i] {
					le = 0
				} else {
					le = 1
				}
			}
		}
		if le < 0 {
			le = 1
		}
		Fprintln(out, (ans+le)%mod)
	}
}

func main() { run(os.Stdin, os.Stdout) }
