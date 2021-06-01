package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1285E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type seg struct{ l, r int64 }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]seg, n)
		es := make([]int64, 0, n*2)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
			a[i].l += 1e9
			a[i].r += 1e9
			es = append(es, a[i].l<<1, a[i].r<<1|1)
		}
		sort.Slice(es, func(i, j int) bool { return es[i] < es[j] })

		var ans, dup, mxDup int
		nonDupSeg := []seg{}
		l := int64(-1)
		for _, e := range es {
			if e&1 > 0 {
				dup--
				if dup == 1 {
					l = e >> 1
				} else if dup == 0 {
					l = -1
					ans++
				}
			} else {
				if dup == 1 && l >= 0 {
					nonDupSeg = append(nonDupSeg, seg{l, e >> 1})
				}
				dup++
				mxDup = max(mxDup, dup)
			}
		}
		if mxDup == 1 { // 互不相交的特殊情况，删除一条会使答案变小！
			Fprintln(out, n-1)
			continue
		}

		mx := 0
		for _, p := range a {
			i := sort.Search(len(nonDupSeg), func(i int) bool { return nonDupSeg[i].l >= p.l })
			j := sort.Search(len(nonDupSeg), func(i int) bool { return nonDupSeg[i].l >= p.r })
			mx = max(mx, j-i)
		}
		Fprintln(out, ans+mx)
	}
}

//func main() { CF1285E(os.Stdin, os.Stdout) }
