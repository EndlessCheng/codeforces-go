package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m, ans int
	Fscan(in, &n, &m)
	n++
	k := bits.Len(uint(n))
	l1 := 0
	r1 := bits.OnesCount(uint(m & (1<<k - 1)))
	for i := k - 1; i >= 0; i-- {
		r1 -= m >> i & 1
		if n>>i&1 == 0 {
			continue
		}
		if i > 0 {
			ans += 1 << (i - 1) % mod * r1
		}
		ans += 1 << i % mod * l1
		l1 += n & m >> i & 1
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
