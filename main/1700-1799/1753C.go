package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1753C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	const mx int = 2e5
	inv := [mx + 1]int64{}
	inv[1] = 1
	for i := 2; i <= mx; i++ {
		inv[i] = int64(mod-mod/i) * inv[mod%i] % mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]byte, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		c0 := bytes.Count(a, []byte{0})
		p1 := bytes.Count(a[:c0], []byte{1})
		s := int64(0)
		for i := 1; i <= p1; i++ {
			s = (s + inv[i]*inv[i]) % mod
		}
		Fprintln(out, int64(n)*int64(n-1)/2%mod*s%mod)
	}
}

//func main() { CF1753C(os.Stdin, os.Stdout) }
