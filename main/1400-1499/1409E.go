package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1409E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, l, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l)
		cnt := map[int]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
		}

		n = len(cnt)
		a := make([]int, 0, n)
		for v := range cnt {
			a = append(a, v)
		}
		sort.Ints(a)
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + cnt[v]
		}
		cs := make([]int, n)
		for i, v := range a {
			cs[i] = sum[i+1] - sum[sort.SearchInts(a[:i], v-l)]
		}
		mxC := make([]int, n)
		mxC[0] = cs[0]
		for i := 1; i < n; i++ {
			mxC[i] = max(mxC[i-1], cs[i])
		}
		ans := 0
		for i, c := range cs {
			if j := sort.SearchInts(a[:i], a[i]-l); j > 0 {
				c += mxC[j-1]
			}
			ans = max(ans, c)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1409E(os.Stdin, os.Stdout) }
