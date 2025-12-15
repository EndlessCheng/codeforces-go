package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1983D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := [][]int{}
		a := make([]int, n)
		b := make([]int, n)
		f := func() (res int) {
			for i := range a {
				Fscan(in, &a[i])
				b[i] = i
			}

			c := slices.Clone(a)
			slices.Sort(c)
			s = append(s, c)

			slices.SortFunc(b, func(x, y int) int { return a[x] - a[y] })
			for i, v := range b {
				if v < 0 {
					continue
				}
				for b[i] >= 0 {
					nxt := b[i]
					b[i] = -1
					i = nxt
				}
				res ^= 1
			}
			return
		}
		if f() == f() && slices.Equal(s[0], s[1]) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1983D(bufio.NewReader(os.Stdin), os.Stdout) }
