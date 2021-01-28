package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1399F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type pair struct{ l, r, contains int }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ps := make([]pair, n)
		a := []int{}
		for i := range ps {
			Fscan(in, &ps[i].l, &ps[i].r)
			a = append(a, ps[i].l, ps[i].r)
		}
		sort.Ints(a)
		k := 1
		kth := map[int]int{a[0]: k}
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				k++
				kth[a[i]] = k
			}
		}
		for i, p := range ps {
			ps[i].l = kth[p.l]
			ps[i].r = kth[p.r]
		}
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.r-a.l < b.r-b.l })

		k++
		ids := make([][]int, k)
		for i, p := range ps {
			ids[p.r] = append(ids[p.r], i)
		}
		for i, p := range ps {
			dp := make([]int, k)
			for j := p.l; j <= p.r; j++ {
				dp[j] = dp[j-1]
				for _, id := range ids[j] {
					if q := ps[id]; q.l >= p.l {
						dp[j] = max(dp[j], dp[q.l-1]+q.contains)
					}
				}
			}
			ps[i].contains = dp[p.r] + 1
		}

		rp := make([][]pair, k)
		for _, p := range ps {
			rp[p.r] = append(rp[p.r], p)
		}
		dp := make([]int, k)
		for i := 1; i < k; i++ {
			dp[i] = dp[i-1]
			for _, p := range rp[i] {
				dp[i] = max(dp[i], dp[p.l-1]+p.contains)
			}
		}
		Fprintln(out, dp[k-1])
	}
}

//func main() { CF1399F(os.Stdin, os.Stdout) }
