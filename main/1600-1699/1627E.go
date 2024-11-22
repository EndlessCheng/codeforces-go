package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1627E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k, a, b, c, d, h int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		x := make([]int, n)
		for i := range x {
			Fscan(in, &x[i])
		}
		type tuple struct{ x, y, h int }
		to := make([]map[int][]tuple, n)
		f := make([]map[int]int, n)
		for i := range to {
			to[i] = map[int][]tuple{}
			f[i] = map[int]int{}
		}
		for range k {
			Fscan(in, &a, &b, &c, &d, &h)
			to[c-1][d-1] = append(to[c-1][d-1], tuple{a - 1, b - 1, h})
		}

		for y, ps := range to[n-1] {
			for _, p := range ps {
				res := (m-1-y)*x[n-1] - p.h
				if _, ok := f[p.x][p.y]; !ok {
					f[p.x][p.y] = res
				} else {
					f[p.x][p.y] = min(f[p.x][p.y], res)
				}
			}
		}
		for i := n - 2; i >= 0; i-- {
			type pair struct{ y, res int }
			fi := make([]pair, 2, len(f[i])+2)
			fi[0] = pair{-1, 1e18}
			fi[1] = pair{m, 1e18}
			for y, res := range f[i] {
				fi = append(fi, pair{y, res})
			}
			slices.SortFunc(fi, func(a, b pair) int { return a.y - b.y })
			for j := len(fi) - 3; j > 0; j-- {
				fi[j].res = min(fi[j].res, fi[j+1].res+(fi[j+1].y-fi[j].y)*x[i])
			}
			if i == 0 {
				ans := fi[1].res + fi[1].y*x[i]
				if ans >= 1e17 { // 易错点
					Fprintln(out, "NO ESCAPE")
				} else {
					Fprintln(out, ans)
				}
				break
			}

			mp := to[i]
			ti := make([]int, 0, len(mp))
			for y := range mp {
				ti = append(ti, y)
			}
			slices.Sort(ti)

			mn := int(1e18)
			j := 1
			for _, y := range ti {
				for fi[j].y <= y {
					mn = min(mn+(fi[j].y-fi[j-1].y)*x[i], fi[j].res)
					j++
				}
				for _, p := range mp[y] {
					res := min(mn+(y-fi[j-1].y)*x[i], fi[j].res+(fi[j].y-y)*x[i]) - p.h
					if _, ok := f[p.x][p.y]; !ok {
						f[p.x][p.y] = res
					} else {
						f[p.x][p.y] = min(f[p.x][p.y], res)
					}
				}
			}
		}
	}
}

//func main() { cf1627E(bufio.NewReader(os.Stdin), os.Stdout) }
