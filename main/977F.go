package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF977F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	dp := map[int]int{}
	var n, end int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if dp[a[i]] = dp[a[i]-1] + 1; dp[a[i]] > dp[end] {
			end = a[i]
		}
	}
	st := end - dp[end] + 1
	ans := []interface{}{}
	for i, v := range a {
		if v == st {
			st++
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	Fprint(out, ans...)
}

//func main() { CF977F(os.Stdin, os.Stdout) }
