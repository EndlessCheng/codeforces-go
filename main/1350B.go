package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1350B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mx int = 1e5
	ds := [mx + 1][]int{}
	for i := 1; i <= mx; i++ {
		for j := 2 * i; j <= mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		dp := make([]int, n)
		ans := 0
		for i := range a {
			Fscan(in, &a[i])
			dp[i] = 1
			for _, d := range ds[i+1] {
				if a[d-1] < a[i] {
					dp[i] = max(dp[i], dp[d-1]+1)
				}
			}
			ans = max(ans, dp[i])
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1350B(os.Stdin, os.Stdout) }
