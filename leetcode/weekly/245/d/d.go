package main

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func earliestAndLatestDP(n, firstPlayer, secondPlayer int) []int {
	type pair struct{ earliest, latest int }
	memo := make([][][]pair, n+1)
	for i := range memo {
		memo[i] = make([][]pair, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]pair, n+1)
		}
	}

	var dfs func(int, int, int) pair
	dfs = func(n, first, second int) (res pair) {
		// AB 相遇
		if first+second == n+1 {
			return pair{1, 1}
		}

		// 保证 A 左边人数比 B 右边人数少
		// 注：题目已保证 first < second
		if first+second > n+1 {
			first, second = n+1-second, n+1-first
		}

		p := &memo[n][first][second]
		if p.earliest > 0 {
			return *p
		}
		defer func() { *p = res }()

		m := (n + 1) / 2 // 下一回合人数
		// AB 之间保留 [minMid, maxMid) 个人
		minMid, maxMid := 0, second-first
		if second > m { // B 在中轴线右侧
			minMid, maxMid = second-n/2-1, m-first
		}

		res.earliest = math.MaxInt
		for left := range first { // 枚举 A 左侧保留 left 个人
			for mid := minMid; mid < maxMid; mid++ { // 枚举 AB 之间保留 mid 个人
				// 无需枚举 B 右侧保留多少个人，因为剩下的 m-2-left-mid 个人都在 B 右侧
				r := dfs(m, left+1, left+mid+2)
				res.earliest = min(res.earliest, r.earliest)
				res.latest = max(res.latest, r.latest)
			}
		}

		// 加上当前回合数
		res.earliest++
		res.latest++
		return res
	}

	ans := dfs(n, firstPlayer, secondPlayer)
	return []int{ans.earliest, ans.latest}
}

func earliestAndLatest2(n, first, second int) []int {
	// AB 一开始就相遇
	if first+second == n+1 {
		return []int{1, 1}
	}

	// 保证 A 左边人数比 B 右边人数少
	// 注：题目已保证 first < second
	if first+second > n+1 {
		first, second = n+1-second, n+1-first
	}

	calcEarliestRounds := func(n int) int {
		res := 1 // 初始回合

		// 情况 5：AB 太靠左了
		if first+second <= (n+1)/2 {
			for first+second <= (n+1)/2 {
				res++
				n = (n + 1) / 2
			}

			// 情况 5a：AB 不相邻
			// 在上面循环的最后一回合，总是可以把局面调整为某些情况，使 AB 下回合就能相遇
			if second-first > 1 {
				return res + 1
			}

			// 情况 5b：AB 相邻
			// 上面循环结束后，转化为情况 1
		}

		// 情况 1：AB 相邻（由于 AB 不相遇，B 不可能在中轴线右侧。注意我们已保证 A 左边人数比 B 右边人数少）
		if second-first == 1 {
			// 先过一回合
			res++
			n = (n + 1) / 2
			// 在 AB 相邻的情况下，当且仅当 n 是偶数的时候相遇（推导过程见图）
			for n%2 > 0 {
				res++
				n = (n + 1) / 2
			}
			return res
		}

		// 情况 2：B 在中轴线或中轴线左侧
		if second <= (n+1)/2 {
			// 可以让 AB 左右人数一样多（构造方式见图），下回合就能相遇
			return res + 1
		}

		// 情况 3：AB 之间恰有一个人
		if second-first == 2 {
			// 下回合 AB 必定相邻，变成情况 1
			res++
			n = (n + 1) / 2
			for n%2 > 0 {
				res++
				n = (n + 1) / 2
			}
			return res
		}

		// 情况 4c：A 左侧有奇数个人，且 B 与 A' 相邻
		if first%2 == 0 && first+second == n {
			// 一回合后，转化为情况 4a
			res++
		}

		// 情况 4a：A 左侧有偶数个人
		// 情况 4b：A 左侧有奇数个人，且 B 与 A' 不相邻
		// 下回合就能相遇
		return res + 1
	}

	// 计算最早回合数
	earliest := calcEarliestRounds(n)

	// 计算最晚回合数
	latest := min(bits.Len(uint(n-1)), n+1-second)

	return []int{earliest, latest}
}

func earliestAndLatest(n, first, second int) []int {
	if first+second == n+1 {
		return []int{1, 1}
	}

	if first+second > n+1 {
		first, second = n+1-second, n+1-first
	}

	calcEarliestRounds := func(n int) int {
		res := 1

		if first+second <= (n+1)/2 {
			// 计算满足 first+second > ceil(n / 2^(k+1)) 的最小 k，推导过程见题解
			k := bits.Len(uint((n-1)/(first+second-1))) - 1
			res += k
			n = (n-1)>>k + 1 // n = ceil(n / 2^k)

			if second-first > 1 {
				return res + 1
			}
		}

		// 情况 1 和情况 3 合并，情况 2 合并到最后的 return
		if second-first == 1 || second > (n+1)/2 && second-first == 2 {
			// 先把 n 变成 ceil(n/2)，然后计算需要多少次 ceil(n/2) 的操作才能把 n 变成偶数，推导过程见题解
			// 这里把 (n+1)/2 和 n-1 合并，得到 (n+1)/2-1 = (n-1)/2
			return res + 1 + bits.TrailingZeros(uint((n-1)/2))
		}

		if second > (n+1)/2 && first%2 == 0 && first+second == n {
			res++
		}

		return res + 1
	}

	latestRounds := min(bits.Len(uint(n-1)), n+1-second)

	return []int{calcEarliestRounds(n), latestRounds}
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
