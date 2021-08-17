package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1203F2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type pair struct{ req, delta int }
	var n, m, req, delta, ans int
	var a, b []pair
	for Fscan(in, &n, &m); n > 0; n-- {
		Fscan(in, &req, &delta)
		if delta >= 0 {
			a = append(a, pair{req, delta})
		} else {
			b = append(b, pair{req, delta})
		}
	}

	sort.Slice(a, func(i, j int) bool { return a[i].req < a[j].req })
	for _, p := range a {
		if m < p.req {
			break
		}
		m += p.delta
		ans++
	}

	sort.Slice(b, func(i, j int) bool { return b[i].req+b[i].delta < b[j].req+b[j].delta })
	dp := make([]int, m+1)
	for _, p := range b {
		w := -p.delta
		for j := m; j >= max(w, p.req); j-- {
			dp[j] = max(dp[j], dp[j-w]+1)
		}
	}
	Fprint(out, ans+dp[m])
}

//func main() { CF1203F2(os.Stdin, os.Stdout) }
