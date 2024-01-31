package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1800F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	var s string
	cnt := map[uint32]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		var m, all uint32
		for _, c := range s {
			b := uint32(1) << (c - 'a')
			m ^= b
			all |= b
		}
		none := 1<<26 - 1 ^ all
		cm := 1<<26 - 1 ^ m
		for t := none; t > 0; t &= t - 1 {
			p := uint32(bits.TrailingZeros32(t))
			ans += cnt[(cm^1<<p)<<5|p]
			cnt[m<<5|p]++
		}
	}
	Fprint(out, ans)
}

//func main() { cf1800F(os.Stdin, os.Stdout) }
