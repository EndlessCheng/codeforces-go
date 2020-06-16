package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF1036C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	calc := func(u int64) int {
		s := strconv.FormatInt(u, 10)
		n := len(s)
		dp := make([][4]int, n)
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(p, sum int, isUpper bool) int
		f = func(p, sum int, isUpper bool) (cnt int) {
			if sum > 3 {
				return
			}
			if p >= n {
				return 1
			}
			dv := &dp[p][sum]
			if !isUpper && *dv >= 0 {
				return *dv
			}
			defer func() {
				if !isUpper {
					*dv = cnt
				}
			}()
			up := byte('9')
			if isUpper {
				up = s[p]
			}
			for digit := byte('0'); digit <= up; digit++ {
				tmp := sum
				if digit > '0' {
					tmp++
				}
				cnt += f(p+1, tmp, isUpper && digit == up)
			}
			return
		}
		return f(0, 0, true)
	}

	var t, l, r int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &l, &r)
		Fprintln(out, calc(r)-calc(l-1))
	}
}

//func main() { CF1036C(os.Stdin, os.Stdout) }
