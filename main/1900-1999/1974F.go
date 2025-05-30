package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1974F(in io.Reader, out io.Writer) {
	var T, R, C, n, m, k int
	var op string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &R, &C, &n, &m)
		type pair struct{ r, c int }
		ud := make([]pair, n)
		for i := range ud {
			Fscan(in, &ud[i].r, &ud[i].c)
		}
		slices.SortFunc(ud, func(a, b pair) int { return a.r - b.r })
		lr := slices.Clone(ud)
		slices.SortFunc(lr, func(a, b pair) int { return a.c - b.c })

		ans := [2]int{}
		u, d, l, r := 1, R, 1, C
		for i := range m {
			Fscan(in, &op, &k)
			if op == "U" {
				u += k
				for len(ud) > 0 && ud[0].r < u {
					if v := ud[0].c; l <= v && v <= r {
						ans[i%2]++
					}
					ud = ud[1:]
				}
			} else if op == "D" {
				d -= k
				for len(ud) > 0 && ud[len(ud)-1].r > d {
					if v := ud[len(ud)-1].c; l <= v && v <= r {
						ans[i%2]++
					}
					ud = ud[:len(ud)-1]
				}
			} else if op == "L" {
				l += k
				for len(lr) > 0 && lr[0].c < l {
					if v := lr[0].r; u <= v && v <= d {
						ans[i%2]++
					}
					lr = lr[1:]
				}
			} else {
				r -= k
				for len(lr) > 0 && lr[len(lr)-1].c > r {
					if v := lr[len(lr)-1].r; u <= v && v <= d {
						ans[i%2]++
					}
					lr = lr[:len(lr)-1]
				}
			}
		}
		Fprintln(out, ans[0], ans[1])
	}
}

//func main() { cf1974F(bufio.NewReader(os.Stdin), os.Stdout) }
