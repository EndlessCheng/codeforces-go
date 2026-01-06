package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2146D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		a := make([]int, r-l+1)
		base := l
		lv, rv := l, r
		for l < r {
			w := bits.Len(uint(l ^ r))
			hb := l&^(1<<w-1) | 1<<(w-1)
			lsz, rsz := hb-l, r-hb+1
			if lsz < rsz {
				for i := hb + lsz - 1; i >= l; i-- {
					a[i-base] = lv
					lv++
				}
				l = hb + lsz
			} else {
				for i := hb - rsz; i <= r; i++ {
					a[i-base] = rv
					rv--
				}
				r = hb - rsz - 1
			}
		}
		if l == r {
			a[l-base] = lv
		}

		or := 0
		for i, v := range a {
			or += v | (base + i)
		}
		Fprintln(out, or)
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2146D2(bufio.NewReader(os.Stdin), os.Stdout) }
