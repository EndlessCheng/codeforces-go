package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF626F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n, n+1)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	a = append(a, 0) // 避免越界

	memo := [200][101][1001]int{}
	for i := range memo {
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int64
	dfs = func(i, groups, leftK int) (res int64) {
		if leftK < 0 || groups > i+1 { // groups > i+1 说明剩余数字不够组成最小值
			return
		}
		if i < 0 {
			if groups == 0 {
				return 1
			}
			return
		}
		p := &memo[i][groups][leftK]
		if *p != -1 {
			return int64(*p)
		}
		leftK -= (a[i+1] - a[i]) * groups
		res = dfs(i-1, groups+1, leftK) // a[i] 作为最大值
		res += dfs(i-1, groups, leftK) * int64(groups+1) // 不参与最大最小：从 groups 中选一个组   这里 +1 是只有一个数的组的方案数
		if groups > 0 {
			res += dfs(i-1, groups-1, leftK) * int64(groups) // a[i] 作为最小值：从 groups 中选一个组
		}
		res %= mod
		*p = int(res) // 记忆化
		return
	}
	Fprint(out, dfs(n-1, 0, k))
}

//func main() { CF626F(os.Stdin, os.Stdout) }
