package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF883I(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	Fprint(out, sort.Search(a[n]-a[1], func(lim int) bool {
		dp := make([]int, n+1) // dp[i] 表示前 i 个数中能满足 lim 的最后一个数的位置
		for i, p := k, 0; i <= n; i++ {
			if a[i]-a[dp[i-k]+1] <= lim {
				p = i
			}
			dp[i] = p
		}
		return dp[n] == n
	}))
}

//func main() { CF883I(os.Stdin, os.Stdout) }
