package main

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func earliestAndLatestDP(n, firstPlayer, secondPlayer int) []int {
	type pair struct{ min, max int } // 最早回合数和最晚回合数
	memo := make([][][]pair, n+1)
	for i := range memo {
		memo[i] = make([][]pair, n)
		for j := range memo[i] {
			memo[i][j] = make([]pair, n)
		}
	}
	var dfs func(int, int, int) pair
	dfs = func(n, first, second int) pair {
		if first+second == n-1 { // first 和 second 发生比拼
			return pair{1, 1}
		}
		if first >= n-1-first || first > n-1-second { // 为简化后续枚举过程，在枚举前处理一下两名选手的位置
			first, second = n-1-second, n-1-first
		}
		dv := &memo[n][first][second]
		if dv.min > 0 {
			return *dv
		}
		res := pair{math.MaxInt, 0}
		mid := (n + 1) / 2 // 下一轮人数
		var r pair
		for i := range first + 1 { // 枚举第一名选手左侧保留多少个人
			for j := range min(second, n-1-second) - first { // 枚举第一名选手和第二名选手中间保留多少个人
				if second < mid { // 两人同侧（处理位置后都位于中间位置左侧）
					r = dfs(mid, i, i+j+1)
				} else { // 两人异侧
					r = dfs(mid, i, i+j+1+(second*2-n+1)/2)
				}
				res.min = min(res.min, r.min)
				res.max = max(res.max, r.max)
			}
		}
		// 加上当前回合数
		res.min++
		res.max++
		*dv = res
		return res
	}
	ans := dfs(n, firstPlayer-1, secondPlayer-1) // 编号改成从 0 开始
	return []int{ans.min, ans.max}
}

func earliestAndLatest(n, first, second int) []int {
	// AB 一开始就相遇
	if first+second == n+1 {
		return []int{1, 1}
	}

	// 保证 A 左边人数比 B 右边人数少
	// 注：题目已保证 first < second
	if first+second > n+1 {
		first, second = n+1-second, n+1-first
	}

	// 计算最早回合数
	calcEarliest := func(n int) int {
		earliest := 1 // 初始回合

		// AB 太靠左了（靠左的定义见情况 xx）
		if first+second <= (n+1)/2 {
			for first+second <= (n+1)/2 {
				earliest++
				n = (n + 1) / 2
			}

			// 情况 xx，AB 不相邻
			// 总是可以调整为情况 xx 或者情况 xx，只需一回合就能相遇
			if second-first > 1 {
				return earliest + 1
			}
		}

		// 情况 xx，AB 相邻（由于 AB 不相遇，所以 B 不可能在中轴线右侧）
		if second-first == 1 {
			// 先过一回合
			earliest++
			n = (n + 1) / 2
			// 在相邻的情况下，当且仅当 n 是偶数的时候相遇（推导过程见图）
			for n%2 > 0 {
				earliest++
				n = (n + 1) / 2
			}
			return earliest
		}

		// 下面讨论 AB 不相邻且 first+second > ceil(n/2) 的情况

		// 情况 xx，B 在中轴线或中轴线左侧     todo
		if second <= (n+1)/2 {
			// 下回合就能相遇（构造方式见图）
			return earliest + 1
		}

		// 情况 xx，B 在中轴线右侧，且两人隔了一个人
		if second-first == 2 { 
			// 由于 B 右侧不能凭空赢一个，所以下回合 AB 必定相邻，变成情况 xx
			earliest++
			n = (n + 1) / 2
			for n%2 > 0 {
				earliest++
				n = (n + 1) / 2
			}
			return earliest
		}

		// A 左侧有奇数个人，且 B 与 A' 相邻
		if first%2 == 0 && first+second == n {
			// 一回合后，总是可以转化为情况 xx   todo
			earliest++
		}

		// 下回合就能相遇
		return earliest + 1
	}

	// 计算最晚回合数
	latest := min(bits.Len(uint(n-1)), n+1-second)

	return []int{calcEarliest(n), latest}
}

// 对拍
func main() {
	for n := 2; n <= 60; n++ {
		for f := 1; f < n; f++ {
			for s := f + 1; s <= n; s++ {
				myAns := earliestAndLatest(n, f, s)
				correctAns := earliestAndLatestDP(n, f, s)
				if !slices.Equal(myAns, correctAns) {
					fmt.Println(n, f, s)
				}
			}
		}
		fmt.Println("done", n)
	}
}
