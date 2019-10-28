package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1251D(reader io.Reader, writer io.Writer) {
	search := func(n int64, f func(int64) bool) int64 {
		i, j := int64(0), n
		for i < j {
			h := (i + j) >> 1
			if f(h) {
				j = h
			} else {
				i = h + 1
			}
		}
		return i
	}
	type pair struct{ l, r int64 }
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var q, n int
	for Fscan(in, &q); q > 0; q-- {
		var money, baseCost int64
		Fscan(in, &n, &money)
		ps := make([]pair, n)
		for i := range ps {
			Fscan(in, &ps[i].l, &ps[i].r)
			baseCost += ps[i].l
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].l > ps[j].l })
		ans := search(money+1e9+1, func(x int64) bool {
			// 把 r >= x 的 l 找出来计算额外花费（n/2 个即可），然后判断中位数是否已达标，或者 cost <= money
			cnt := 0
			cost := baseCost
			for _, p := range ps {
				if p.r >= x {
					cnt++
					if p.l < x {
						cost += x - p.l
					}
					if 2*cnt-1 == n {
						if p.l > x {
							return !true
						}
						return !(cost <= money)
					}
				}
			}
			return !false
		})
		Fprintln(out, ans-1)
	}
}

func main() {
	Sol1251D(os.Stdin, os.Stdout)
}
