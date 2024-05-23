package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func p4310(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v uint
	f := [30]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		mx := 0
		for i := v; i > 0; i &= i - 1 {
			mx = max(mx, f[bits.TrailingZeros(i)])
		}
		for ; v > 0; v &= v - 1 {
			f[bits.TrailingZeros(v)] = mx + 1
		}
	}
	Fprintln(out, slices.Max(f[:]))
}

//func main() { p4310(os.Stdin, os.Stdout) }
