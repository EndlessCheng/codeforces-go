package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2094H1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 100_001
	divisors := [mx][]int{}
	for i := 2; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n, q, v, l, r int
	pos := [mx][]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}
		for range q {
			Fscan(in, &v, &l, &r)
			l--
			ans := 0
			for {
				mnI := n
				for _, d := range divisors[v] {
					ps := pos[d]
					i := sort.SearchInts(ps, l)
					if i < len(ps) && ps[i] < r {
						mnI = min(mnI, ps[i])
					}
				}
				if mnI == n {
					ans += v * (r - l)
					break
				}
				ans += v * (mnI - l)
				l = mnI
				d := a[l]
				for v /= d; v%d == 0; v /= d {
				}
			}
			Fprintln(out, ans)
		}
		for _, v := range a {
			pos[v] = pos[v][:0]
		}
	}
}

// 离线做法，更快
func cf2094H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 100_001
	divisors := [mx][]int{}
	for i := 2; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n, q, v, l, r int
	pos := [mx][]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}
		type tuple struct{ v, r, i int }
		g := make([][]tuple, n)
		for i := range q {
			Fscan(in, &v, &l, &r)
			g[l-1] = append(g[l-1], tuple{v, r, i})
		}
		ans := make([]int, q)
		for l, qs := range g {
			for _, t := range qs {
				v := t.v
				mnI := n
				for _, d := range divisors[v] {
					ps := pos[d]
					for len(ps) > 0 && ps[0] < l {
						ps = ps[1:]
					}
					if len(ps) > 0 && ps[0] < t.r {
						mnI = min(mnI, ps[0])
					}
					pos[d] = ps
				}
				if mnI == n {
					ans[t.i] += v * (t.r - l)
					continue
				}
				ans[t.i] += v * (mnI - l)
				d := a[mnI]
				for v /= d; v%d == 0; v /= d {
				}
				ans[t.i] += v
				t.v = v
				mnI++
				if mnI < n {
					g[mnI] = append(g[mnI], t)
				}
			}
		}
		for _, v := range ans {
			Fprintln(out, v)
		}
		for _, v := range a {
			pos[v] = pos[v][:0]
		}
	}
}

//func main() { cf2094H(bufio.NewReader(os.Stdin), os.Stdout) }
