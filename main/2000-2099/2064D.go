package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf2064D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		pre := [31]int{}
		for i := range pre {
			pre[i] = -1
		}
		pres := make([][31]int, n+1)
		for i := range a {
			Fscan(in, &a[i])
			m := bits.Len(uint(a[i]))
			pres[i] = pre
			pre[m] = i
		}
		pres[n] = pre
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
				if m+1 < preM {
					maxP = max(maxP, slices.Max(pres[cur][m+1:preM]))
				}
				p := pres[cur][m]
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
