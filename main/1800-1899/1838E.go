package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1838E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var T, n, m int
	var k int64
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		ans := pow(k, m)
		for i, c := 0, int64(1); i < n; i++ {
			Fscan(in, &s)
			ans -= c * pow(k-1, m-i) % mod
			c = c * int64(m-i) % mod * pow(int64(i+1), mod-2) % mod
		}
		Fprintln(out, (ans%mod+mod)%mod)
	}
}

//func main() { CF1838E(os.Stdin, os.Stdout) }
