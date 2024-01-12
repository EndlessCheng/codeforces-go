package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf309C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, ans uint
	Fscan(in, &n, &m)
	var leftA, cntB [30]int
	for ; n > 0; n-- {
		for Fscan(in, &v); v > 0; v &= v - 1 {
			leftA[bits.TrailingZeros(v)]++
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &v)
		cntB[v]++
	}

outer:
	for i, c := range cntB {
	next:
		for ; c > 0; c-- {
			for j := i; j < 30; j++ {
				if leftA[j] > 0 {
					ans++
					leftA[j]--
					continue next
				}
				leftA[j]++
			}
			break outer
		}
	}
	Fprint(out, ans)
}

//func main() { cf309C(os.Stdin, os.Stdout) }
