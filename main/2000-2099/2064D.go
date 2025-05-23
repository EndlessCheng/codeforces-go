package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://github.com/EndlessCheng
func cf2064D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		pos := [31][]int{}
		for i := range a {
			Fscan(in, &a[i])
			m := bits.Len(uint(a[i]))
			pos[m] = append(pos[m], i)
		}
		s := make([]int, n+1)
		for i := n - 1; i >= 0; i-- {
			s[i] = s[i+1] ^ a[i]
		}

		for range q {
			Fscan(in, &x)
			oriX := x
			cur, preM, maxP := n, 31, -1
			for {
				m := bits.Len(uint(x))
				for _, ps := range pos[m+1 : preM] {
					if j := sort.SearchInts(ps, cur); j > 0 {
						maxP = max(maxP, ps[j-1])
					}
				}
				p := -1
				if j := sort.SearchInts(pos[m], cur); j > 0 {
					p = pos[m][j-1]
				}
				if p <= maxP {
					Fprint(out, n-1-maxP, " ")
					break
				}
				x = oriX ^ s[p+1]
				if x < a[p] {
					Fprint(out, n-1-p, " ")
					break
				}
				x ^= a[p]
				preM = m + 1
				cur = p
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2064D(bufio.NewReader(os.Stdin), os.Stdout) }
