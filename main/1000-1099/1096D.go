package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func Sol1096D(reader io.Reader, writer io.Writer) {
	mins := func(vals ...int64) int64 {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val < ans {
				ans = val
			}
		}
		return ans
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n, &s)
	costs := make([]int, n)
	for i := range costs {
		Fscan(in, &costs[i])
	}

	dp := [2][4]int64{}
	for i := 1; i <= n; i++ {
		i1, i0 := i&1, (i-1)&1
		for j := range dp[i1] {
			dp[i1][j] = dp[i0][j]
		}
		j := strings.Index("hard", string(s[i-1]))
		if j == -1 {
			continue
		}
		dp[i1][j] += int64(costs[i-1])
		if j < 3 && dp[i0][j] < dp[i1][j+1] {
			dp[i1][j+1] = dp[i0][j]
		}
	}
	Fprint(out, mins(dp[n&1][:]...))
}

//func main() {
//	Sol1096D(os.Stdin, os.Stdout)
//}
