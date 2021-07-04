package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1296E2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s []byte
	Fscan(in, &n, &s)
	dp := []int{}
	ans := make([]interface{}, n)
	for i, b := range s {
		v := int(-b)
		p := sort.SearchInts(dp, v)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
		ans[i] = p + 1
	}
	Fprintln(out, len(dp))
	Fprintln(out, ans...)
}

//func main() { CF1296E2(os.Stdin, os.Stdout) }
