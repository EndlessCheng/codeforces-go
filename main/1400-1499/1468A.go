package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1468A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ v, i int }
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &v)
		dp, dp2 := []int{}, []int{v}
		stk := []pair{{1e9, 0}}
		for n--; n > 0; n-- {
			Fscan(in, &v)
			p := sort.SearchInts(dp, v+1)
			if p < len(dp) {
				dp[p] = v
				dp2[p+1] = v
			} else {
				dp = append(dp, min(v, dp2[len(dp2)-1]))
				dp2 = append(dp2, v)
			}
			for stk[len(stk)-1].v <= v {
				p := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				dp[p.i] = min(dp[p.i], p.v)
			}
			if p+1 < len(dp) {
				stk = append(stk, pair{v, p + 1})
			}
		}
		Fprintln(out, len(dp2))
	}
}

//func main() { CF1468A(os.Stdin, os.Stdout) }
