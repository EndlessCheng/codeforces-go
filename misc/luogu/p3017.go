package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func p3017(in io.Reader, out io.Writer) {
	var n, m, r, c int
	Fscan(in, &n, &m, &r, &c)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	cs := make([]int, m)
	ans := sort.Search(n*m*4000, func(low int) bool {
		low++
		clear(cs)
		r := r
		for _, row := range a {
			s, c := 0, c
			for j, v := range row {
				cs[j] += v
				s += cs[j]
				if s >= low {
					s = 0
					c--
				}
			}
			if c <= 0 {
				r--
				if r == 0 {
					return false
				}
				clear(cs)
			}
		}
		return true
	})
	Fprint(out, ans)
}

//func main() { p3017(bufio.NewReader(os.Stdin), os.Stdout) }
