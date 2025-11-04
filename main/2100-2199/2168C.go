package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2168C(in io.Reader, out io.Writer) {
	pos := []int{3, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20}
	var tp string
	var T, n, v int
	Fscan(in, &tp, &T)
	if tp[0] == 'f' {
		for range T {
			Fscan(in, &v)
			a := []any{}
			s := 0
			for t := uint(v - 1); t > 0; t &= t - 1 {
				i := bits.TrailingZeros(t)
				a = append(a, pos[i])
				s ^= pos[i]
			}
			for ; s > 0; s &= s - 1 {
				a = append(a, s&-s)
			}
			Fprintln(out, len(a))
			Fprintln(out, a...)
		}
	} else {
		for range T {
			Fscan(in, &n)
			has := [21]bool{}
			s := 0
			for range n {
				Fscan(in, &v)
				has[v] = true
				s ^= v
			}
			has[s] = !has[s]
			ans := 0
			for i, p := range pos {
				if has[p] {
					ans |= 1 << i
				}
			}
			Fprintln(out, ans+1)
		}
	}
}

//func main() { cf2168C(bufio.NewReader(os.Stdin), os.Stdout) }
