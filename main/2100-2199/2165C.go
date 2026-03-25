package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2165C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		a = a[max(n-30, 0):]

		for range q {
			Fscan(in, &c)
			b := slices.Clone(a)
			ans := 0
			for len(b) > 0 && b[len(b)-1] < c {
				v := b[len(b)-1]
				b = b[:len(b)-1]
				m := bits.Len(uint(c))
				hb := 1 << (m - 1)
				if bits.Len(uint(v)) == m {
					c ^= hb
					v ^= hb
					j := sort.SearchInts(b, v)
					b = slices.Insert(b, j, v)
				} else {
					if len(b) == 0 {
						ans += c - v
						break
					}
					ans += hb - v
					c ^= hb
				}
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { cf2165C(bufio.NewReader(os.Stdin), os.Stdout) }
