package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2009G2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] += n - 2 - i
		}

		mn := make([]int, n-k+1)
		cnt := make([]int, n*2-1)
		cc := make([]int, n+1)
		maxC := 0
		for r, v := range a {
			cc[cnt[v]]--
			cnt[v]++
			cc[cnt[v]]++
			maxC = max(maxC, cnt[v])

			l := r + 1 - k
			if l < 0 {
				continue
			}
			mn[l] = k - maxC

			v = a[l]
			cc[cnt[v]]--
			cnt[v]--
			cc[cnt[v]]++
			if cc[maxC] == 0 {
				maxC--
			}
		}

		type pair struct{ r, i int }
		qs := make([][]pair, n-k+1)
		for i := range q {
			var l, r int
			Fscan(in, &l, &r)
			qs[l-1] = append(qs[l-1], pair{r - k, i})
		}

		ans := make([]int, q)
		type data struct{ r, v, s int }
		st := []data{{len(mn), -1, 0}}
		for i := len(mn) - 1; i >= 0; i-- {
			v := mn[i]
			r := i
			for st[len(st)-1].v >= v {
				r = st[len(st)-1].r
				st = st[:len(st)-1]
			}
			st = append(st, data{r, v, st[len(st)-1].s + (r-i+1)*v})
			for _, p := range qs[i] {
				j := sort.Search(len(st), func(i int) bool { return st[i].r < p.r }) - 1
				ans[p.i] = st[len(st)-1].s - st[j-1].s - (st[j].r-p.r)*st[j].v
			}
		}
		for _, v := range ans {
			Fprintln(out, v)
		}
	}
}

//func main() { cf2009G2(bufio.NewReader(os.Stdin), os.Stdout) }
