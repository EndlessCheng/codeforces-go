package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1918C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, a, b, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &r)
		// 最高位不用管
		for i := bits.Len(uint(a^b)) - 2; i >= 0; i-- {
			bit := 1 << i
			if bit <= r && (a^b)&bit > 0 && abs((a^bit)-(b^bit)) < abs(a-b) {
				a ^= bit
				b ^= bit
				r -= bit
			}
		}
		Fprintln(out, abs(a-b))
	}
}

//func main() { cf1918C(os.Stdin, os.Stdout) }
