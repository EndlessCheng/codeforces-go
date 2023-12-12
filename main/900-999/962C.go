package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func isSQ(x int) bool {
	rt := int(math.Sqrt(float64(x)))
	return rt*rt == x
}

func cf962C(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	mx := -1
	for i := uint(1); i < 1<<len(s); i++ {
		if s[bits.TrailingZeros(i)] == '0' {
			continue
		}
		x := 0
		for j := i; j > 0; j &= j - 1 {
			x = x*10 + int(s[bits.TrailingZeros(j)]&15)
		}
		if isSQ(x) {
			mx = max(mx, bits.OnesCount(i))
		}
	}
	if mx < 0 {
		Fprint(out, -1)
	} else {
		Fprint(out, len(s)-mx)
	}
}

//func main() { cf962C(os.Stdin, os.Stdout) }
