package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1221D(reader io.Reader, writer io.Writer) {
	const inf int64 = 2e18
	min := func(a, b int64) int64 {
		if a <= b {
			return a
		}
		return b
	}
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

	var dp [2][3]int64
	var t int
	for Fscan(in, &t); t > 0; t-- {
		var n int
		Fscan(in, &n)
		h := make([]int, n)
		costs := make([]int, n)
		for i := range h {
			Fscan(in, &h[i], &costs[i])
		}

		row := 0
		dp[0][0] = 0
		dp[0][1] = int64(costs[0])
		dp[0][2] = int64(2 * costs[0])
		for i := 1; i < n; i++ {
			row ^= 1
			for d0 := 0; d0 < 3; d0++ {
				ans := inf
				for d1 := 0; d1 < 3; d1++ {
					if h[i]+d0 != h[i-1]+d1 {
						ans = min(ans, dp[row^1][d1]+int64(d0*costs[i]))
					}
				}
				dp[row][d0] = ans
			}
		}
		Fprintln(out, mins(dp[row][:]...))
	}
}

//func main() {
//	Sol1221D(os.Stdin, os.Stdout)
//}
