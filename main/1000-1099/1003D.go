package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1003D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	coins := [31]int{}
	var n, q int
	var v uint
	for Fscan(in, &n, &q); n > 0; n-- {
		Fscan(in, &v)
		coins[bits.TrailingZeros(v)]++
	}
o:
	for ; q > 0; q-- {
		s, c := 0, coins
		for Fscan(in, &v); v > 0; v &= v - 1 {
			p := bits.TrailingZeros(v)
			for cost := 1; ; cost <<= 1 {
				if cost <= c[p] {
					c[p] -= cost
					s += cost
					break
				}
				if p == 0 {
					Fprintln(out, -1)
					continue o
				}
				cost -= c[p]
				s += c[p]
				c[p] = 0
				p--
			}
		}
		Fprintln(out, s)
	}
}

//func main() { CF1003D(os.Stdin, os.Stdout) }
