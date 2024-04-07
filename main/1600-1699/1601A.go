package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1601A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		cnt := [30]int{}
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			for _s := uint(v); _s > 0; _s &= _s - 1 {
				cnt[bits.TrailingZeros(_s)]++
			}
		}
		g := 0
		for _, c := range cnt {
			g = gcd(g, c)
		}
		for i := 1; i <= n; i++ {
			if g%i == 0 {
				Fprint(out, i, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1601A(os.Stdin, os.Stdout) }
