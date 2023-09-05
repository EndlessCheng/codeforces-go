package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	var n, ans int
	var v uint
	Fscan(in, &n)
	cnt := [60]int{}
	for i := 0; i < n; i++ {
		for Fscan(in, &v); v > 0; v &= v - 1 {
			cnt[bits.TrailingZeros(v)]++
		}
	}
	for i, c := range cnt {
		ans = (ans + c*(n-c)%mod*(1<<i%mod)) % mod
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
