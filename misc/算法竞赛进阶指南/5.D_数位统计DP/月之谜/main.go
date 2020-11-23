package main

import (
	. "fmt"
	"io"
	"os"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
const lowerC, upperC byte = 0, 9

func calc(S int) (res int) {
	if S == 0 {
		return 0
	}
	s := []byte(strconv.Itoa(S))
	for i := range s {
		s[i] &= 15
	}
	n := len(s)
	dp := make([][90][90]int, n)
	for tar := 1; tar < 90; tar++ {
		for i := range dp {
			for j := 0; j < 90; j++ {
				for k := 0; k < 90; k++ {
					dp[i][j][k] = -1
				}
			}
		}
		var f func(p, v, ds int, isUpper bool) int
		f = func(p, v, ds int, isUpper bool) (res int) {
			if p == n {
				if v == 0 && ds == tar {
					return 1
				}
				return
			}
			if !isUpper {
				dv := &dp[p][v][ds]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			up := upperC
			if isUpper {
				up = s[p]
			}
			for digit := lowerC; digit <= up; digit++ {
				res += f(p+1, (v*10+int(digit))%tar, ds+int(digit), isUpper && digit == up)
			}
			return
		}
		res += f(0, 0, 0, true)
	}
	return
}

func run(in io.Reader, out io.Writer) {
	var l, r int
	Fscan(in, &l, &r)
	Fprint(out, calc(r)-calc(l-1))
}

func main() { run(os.Stdin, os.Stdout) }
