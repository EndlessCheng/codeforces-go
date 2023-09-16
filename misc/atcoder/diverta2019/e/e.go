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
	var n, v, xor int
	f := [1 << 20]struct{ s0, sx, pre0 int }{}
	for i := range f {
		f[i].s0 = 1
	}
	cnt0 := 1
	pow2 := (mod + 1) / 2
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		xor ^= v
		if xor == 0 {
			cnt0++
			pow2 = pow2 * 2 % mod
		} else {
			t := &f[xor]
			t.s0 = (t.s0 + t.sx*(cnt0-t.pre0)) % mod
			t.sx = (t.sx + t.s0) % mod
			t.pre0 = cnt0
		}
	}
	if xor > 0 {
		Fprint(out, f[xor].s0)
	} else {
		ans := pow2 // pow(2, cnt0-2)
		for _, t := range f {
			ans += t.sx
		}
		Fprint(out, ans%mod)
	}
}

func main() { run(os.Stdin, os.Stdout) }
