package _00_299

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF269B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s string
	Fscan(in, &n, &s)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i], &s)
	}
	dp := []int{}
	for _, v := range a {
		if i := sort.SearchInts(dp, v+1); i < len(dp) {
			dp[i] = v
		} else {
			dp = append(dp, v)
		}
	}
	Fprint(_w, n-len(dp))
}

//func main() { CF269B(os.Stdin, os.Stdout) }
