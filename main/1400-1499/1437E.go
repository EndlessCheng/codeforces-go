package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1437E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	lis := func(a []int) int {
		dp := []int{}
		for _, v := range a {
			if p := sort.SearchInts(dp, v+1); p < len(dp) {
				dp[p] = v
			} else {
				dp = append(dp, v)
			}
		}
		return len(dp)
	}

	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n+2)
	a[0] = -1e8
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	a[n+1] = 2e9
	b := make([]int, k+2)
	for i := 1; i <= k; i++ {
		Fscan(in, &b[i])
	}
	b[k+1] = n + 1

	for i := 0; i <= k; i++ {
		l, r := b[i], b[i+1]
		if a[r]-a[l] < r-l {
			Fprint(out, -1)
			return
		}
		c := []int{}
		for j := l + 1; j < r; j++ {
			v := a[j]
			if v-a[l] >= j-l && a[r]-v >= r-j {
				c = append(c, v-j)
			}
		}
		ans += lis(c)
	}
	Fprint(out, n-k-ans)
}

//func main() { CF1437E(os.Stdin, os.Stdout) }
