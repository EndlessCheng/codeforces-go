package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	. "sort"
)

// github.com/EndlessCheng/codeforces-go
func CF377C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	a := make(IntSlice, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Sort(Reverse(a))
	Fscan(in, &n)
	op := make([]struct{ team, action string }, n)
	for i := n - 1; i >= 0; i-- {
		Fscan(in, &op[i].action, &op[i].team)
	}

	m := 1 << n
	dp := make([]int, m)
	for s := 1; s < m; s++ {
		op := op[bits.OnesCount(uint(s))-1]
		if op.team == "1" {
			dp[s] = -1e9
			for t, lb := s, 0; t > 0; t ^= lb {
				lb = t & -t
				v := dp[s^lb]
				if op.action == "p" {
					v += a[bits.TrailingZeros(uint(lb))]
				}
				dp[s] = max(dp[s], v)
			}
		} else {
			dp[s] = 1e9
			for t, lb := s, 0; t > 0; t ^= lb {
				lb = t & -t
				v := dp[s^lb]
				if op.action == "p" {
					v -= a[bits.TrailingZeros(uint(lb))]
				}
				dp[s] = min(dp[s], v)
			}
		}
	}
	Fprint(out, dp[m-1])
}

//func main() { CF377C(os.Stdin, os.Stdout) }
