package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF360B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := make([]int, n) // dp[i] 表示不修改 a[i] 时，[1,i] 最多有多少个数可以不用修改
	ans := sort.Search(2e9, func(x int) bool {
		for i, v := range a {
			dp[i] = 1
			for j, w := range a[:i] {
				if int64(abs(v-w)) <= int64(i-j)*int64(x) {
					dp[i] = max(dp[i], dp[j]+1)
				}
			}
			if dp[i] >= n-k {
				return true
			}
		}
		return false
	})
	Fprint(out, ans)
}

//func main() { CF360B(os.Stdin, os.Stdout) }
