package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF486E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	l := make([]int, n)
	dp := []int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		p := sort.SearchInts(dp, v)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
		l[i] = p + 1
	}
	r := make([]int, n)
	dp = nil
	for i := n - 1; i >= 0; i-- {
		v := -a[i]
		p := sort.SearchInts(dp, v)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
		r[i] = p + 1
	}
	ans := make([]byte, n)
	cnt := make([]int, n+1)
	for i, l := range l {
		if l+r[i]-1 != len(dp) {
			ans[i] = '1'
		} else {
			ans[i] = '3'
			cnt[l]++
		}
	}
	for i, tp := range ans {
		if tp == '3' && cnt[l[i]] > 1 {
			ans[i] = '2'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF486E(os.Stdin, os.Stdout) }
