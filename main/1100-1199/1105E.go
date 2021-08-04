package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1105E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, op int
	var s string
	Fscan(in, &n, &m)
	ids := [][]int{}
	id := make(map[string]int, m)
	for ; n > 0; n-- {
		if Fscan(in, &op); op == 1 {
			ids = append(ids, nil)
		} else {
			Fscan(in, &s)
			if _, has := id[s]; !has {
				id[s] = len(id)
			}
			ids[len(ids)-1] = append(ids[len(ids)-1], id[s])
		}
	}

	g := make([]int64, m)
	for i := range g {
		g[i] = 1<<m - 1
	}
	for _, ids := range ids {
		for _, i := range ids {
			for _, j := range ids {
				g[i] &^= 1 << j
			}
		}
	}

	dp := map[int64]int{0: 0}
	var f func(int64) int
	f = func(s int64) int {
		if v, has := dp[s]; has {
			return v
		}
		dp[s] = max(f(s&(s-1)), 1+f(s&g[bits.TrailingZeros64(uint64(s))]))
		return dp[s]
	}
	Fprint(out, f(1<<m-1))
}

//func main() { CF1105E(os.Stdin, os.Stdout) }
