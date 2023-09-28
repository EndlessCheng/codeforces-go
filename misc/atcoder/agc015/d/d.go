package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var low, high uint
	Fscan(in, &low, &high)
	if low == high {
		Fprint(out, 1)
		return
	}
	ans := high - low + 1
	mask := uint(1)<<(bits.Len(high^low)-1) - 1
	high &= mask
	low &= mask
	nh := bits.Len(high)
	if bits.Len(low) <= nh {
		ans += mask - high
	} else {
		ans += mask - low + 1<<nh - high
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
