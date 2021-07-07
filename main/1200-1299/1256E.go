package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1256E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	dp := make([]pair, n+1)
	dp[0].v = -a[0].v
	dp[1].v = 2e9
	dp[2].v = 2e9
	from := make([]int, n+1)
	for i, mi, pre := 3, dp[0], 0; i <= n; i++ {
		if dp[i-3].v < mi.v {
			mi, pre = dp[i-3], i-3
		}
		dp[i] = pair{mi.v + a[i-1].v, mi.i + 1}
		if i < n {
			dp[i].v -= a[i].v
		}
		from[i] = pre
	}
	Fprintln(out, dp[n].v, dp[n].i)
	ans := make([]interface{}, n)
	for i, id := n, 1; i > 0; i = from[i] {
		for j := from[i]; j < i; j++ {
			ans[a[j].i] = id
		}
		id++
	}
	Fprintln(out, ans...)
}

//func main() { CF1256E(os.Stdin, os.Stdout) }
