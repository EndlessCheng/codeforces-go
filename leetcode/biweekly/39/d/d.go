package main

// todo O(m*2^m) https://leetcode-cn.com/problems/distribute-repeating-integers/solution/om2m-c-56ms-100-by-hqztrue-qmfl/

// github.com/EndlessCheng/codeforces-go
func canDistribute(a []int, quantity []int) bool {
	m := 1 << len(quantity)
	sum := make([]int, m)
	for i, v := range quantity {
		bit := 1 << i
		for mask := 0; mask < bit; mask++ {
			sum[bit|mask] = sum[mask] + v
		}
	}

	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	n := len(cnt)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m)
		dp[i][0] = true
	}
	i := 0
	for _, c := range cnt {
		row := dp[i]
		for s, b := range row {
			if b {
				dp[i+1][s] = true
				continue
			}
			for sub := s; sub > 0; sub = (sub - 1) & s {
				if row[s^sub] && sum[sub] <= c {
					dp[i+1][s] = true
					break
				}
			}
		}
		i++
	}
	return dp[n][m-1]
}
