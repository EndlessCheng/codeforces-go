package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF808E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n, maxW, w, v int
	Fscan(in, &n, &maxW)
	a := [3]sort.IntSlice{}
	for ; n > 0; n-- {
		Fscan(in, &w, &v)
		a[w-1] = append(a[w-1], v)
	}
	for _, b := range a {
		sort.Sort(sort.Reverse(b))
	}

	// 枚举 3，然后把两个 1 看成一个 2，贪心取
	// 这种做法也可以用 DP 来实现，dp[i] 表示重量 i 能取到的最大价值，以及该最大价值用了多少 1 和 2
	type pair struct {
		x, y int
		res  int64
	}
	dp := make([]pair, maxW+1)
	for i := 1; i <= maxW; i++ {
		dp[i] = dp[i-1]
		if dp[i].x < len(a[0]) {
			dp[i].res += int64(a[0][dp[i].x])
			dp[i].x++
		}
		if i > 1 && dp[i-2].y < len(a[1]) && dp[i-2].res+int64(a[1][dp[i-2].y]) > dp[i].res {
			dp[i] = dp[i-2]
			dp[i].res += int64(a[1][dp[i-2].y])
			dp[i].y++
		}
	}

	ans := dp[maxW].res
	s3 := int64(0)
	for i, v := range a[2] {
		if (i+1)*3 > maxW {
			break
		}
		s3 += int64(v)
		ans = max(ans, s3+dp[maxW-(i+1)*3].res)
	}
	Fprint(out, ans)
}

//func main() { CF808E(os.Stdin, os.Stdout) }
