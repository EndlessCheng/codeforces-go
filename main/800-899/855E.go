package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF855E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, b int
	var l, r int64
	var s string
	dp := [11][60][]int64{}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = make([]int64, 1<<i)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, bool, bool) int64
	f = func(p, mask int, limitUp, hasD bool) (res int64) {
		if p < 0 {
			if mask == 0 && hasD {
				return 1
			}
			return
		}
		if !limitUp && hasD {
			dv := &dp[b][p][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		up := byte(b - 1)
		if limitUp {
			up = s[len(s)-1-p] & 15
		}
		for ch := byte(0); ch <= up; ch++ {
			tmp := mask
			if hasD || ch > 0 {
				tmp ^= 1 << ch
			}
			res += f(p-1, tmp, limitUp && ch == up, hasD || ch > 0)
		}
		return
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &b, &l, &r)
		s = strconv.FormatInt(r, b)
		ans := f(len(s)-1, 0, true, false)
		s = strconv.FormatInt(l-1, b)
		Fprintln(out, ans-f(len(s)-1, 0, true, false))
	}
}

//func main() { CF855E(os.Stdin, os.Stdout) }
