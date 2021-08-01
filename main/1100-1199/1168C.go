package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1168C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mx = 19

	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]uint, n+1)
	left := make([][mx]int, n+1) // init = 0
	curLeft := [mx]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		v := a[i]
		for j := 0; j < mx; j++ {
			if v>>j&1 > 0 {
				left[i][j] = i
			} else {
				for s := v; s > 0; s &= s - 1 {
					left[i][j] = max(left[i][j], left[curLeft[bits.TrailingZeros(s)]][j])
				}
			}
		}
		for ; v > 0; v &= v - 1 {
			curLeft[bits.TrailingZeros(v)] = i
		}
	}

o:
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		for s := a[l]; s > 0; s &= s - 1 {
			if left[r][bits.TrailingZeros(s)] >= l {
				Fprintln(out, "Shi")
				continue o
			}
		}
		Fprintln(out, "Fou")
	}
}

//func main() { CF1168C(os.Stdin, os.Stdout) }
