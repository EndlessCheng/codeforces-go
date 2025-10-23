package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf771D(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(in, &n, &s)
	var pv, pk, pa []int
	for i, b := range s {
		if b == 'V' {
			pv = append(pv, i)
		} else if b == 'K' {
			pk = append(pk, i)
		} else {
			pa = append(pa, i)
		}
	}
	dp := make([][][][2]int, len(pv)+1)
	for i := range dp {
		dp[i] = make([][][2]int, len(pk)+1)
		for j := range dp[i] {
			dp[i][j] = make([][2]int, len(pa)+1)
			for k := range dp[i][j] {
				dp[i][j][k] = [2]int{-1, -1}
			}
		}
	}
	var f func(int, int, int, int) int
	f = func(i, j, k, preV int) (res int) {
		if i+j+k == n {
			return
		}
		ptr := &dp[i][j][k][preV]
		if *ptr >= 0 {
			return *ptr
		}

		res = 1e9
		if i < len(pv) {
			p := pv[i]
			r := j - sort.SearchInts(pk[:j], p)
			r += k - sort.SearchInts(pa[:k], p)
			res = min(res, f(i+1, j, k, 1)+r)
		}
		if preV == 0 && j < len(pk) {
			p := pk[j]
			r := i - sort.SearchInts(pv[:i], p)
			r += k - sort.SearchInts(pa[:k], p)
			res = min(res, f(i, j+1, k, 0)+r)
		}
		if k < len(pa) {
			p := pa[k]
			r := i - sort.SearchInts(pv[:i], p)
			r += j - sort.SearchInts(pk[:j], p)
			res = min(res, f(i, j, k+1, 0)+r)
		}
		*ptr = res
		return
	}
	Fprint(out, f(0, 0, 0, 0))
}

//func main() { cf771D(os.Stdin, os.Stdout) }
