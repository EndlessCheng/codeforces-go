package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1101C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type tuple struct{ l, r, i int }
		a := make([]tuple, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
			a[i].i = i
		}
		slices.SortFunc(a, func(a, b tuple) int { return a.l - b.l })

		maxR, cnt, preI := a[0].r, 1, a[0].i
		for _, p := range a[1:] {
			if p.l > maxR {
				if cnt == 1 {
					for i := range n {
						if i == preI {
							Fprint(out, "1 ")
						} else {
							Fprint(out, "2 ")
						}
					}
					Fprintln(out)
					continue o
				}
				cnt = 0
			}
			preI = p.i
			cnt++
			maxR = max(maxR, p.r)
		}
		Fprintln(out, -1)
	}
}

//func main() { cf1101C(bufio.NewReader(os.Stdin), os.Stdout) }
