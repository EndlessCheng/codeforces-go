package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1466E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var p2, c [60]int64
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
			for j := range c {
				c[j] += a[i] >> j & 1
			}
		}
		for i := range c {
			p2[i] = 1 << i % mod * n % mod
			c[i] = 1 << i % mod * c[i] % mod
		}
		ans := int64(0)
		for _, v := range a {
			var x, y int64
			for j, c := range c {
				if v>>j&1 > 0 {
					x += c
					y += p2[j]
				} else {
					y += c
				}
			}
			ans = (ans + x%mod*(y%mod)) % mod
		}
		Fprintln(out, ans)
	}
}

func main() { CF1466E(os.Stdin, os.Stdout) }
