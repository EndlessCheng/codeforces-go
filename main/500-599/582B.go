package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF582B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, t, d int
	Fscan(in, &n, &t)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := []int{}
	for i := 0; i < n+10 && t > 0; i++ {
		t--
		cur := len(dp)
		for _, v := range a {
			if p := sort.SearchInts(dp, v+1); p < len(dp) {
				dp[p] = v
			} else {
				dp = append(dp, v)
			}
		}
		d = len(dp) - cur
		cur = len(dp)
	}
	if t == 0 {
		Fprint(out, len(dp))
		return
	}
	Fprintln(out, len(dp)+t*d)
}

//func main() { CF582B(os.Stdin, os.Stdout) }
