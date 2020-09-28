package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1416C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	cnt, x := int64(0), 0
	group := [][]int{a}
	for i := 29; i >= 0; i-- {
		g := group
		group = nil
		s := [2]int64{} // s[0] 为逆序对，s[1] 为正序对
		for _, a := range g {
			b := [2][]int{}
			c := [2]int64{}
			for _, v := range a {
				w := v >> i & 1
				b[w] = append(b[w], v)
				s[w] += c[w^1]
				c[w]++
			}
			for _, a := range b {
				if len(a) > 0 {
					group = append(group, a)
				}
			}
		}
		cnt += min(s[0], s[1])
		if s[0] > s[1] {
			x |= 1 << i
		}
	}
	Fprintln(out, cnt, x)
}

//func main() { CF1416C(os.Stdin, os.Stdout) }
