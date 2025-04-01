package main

import "strconv"

// https://space.bilibili.com/206214
func countSymmetricIntegers(low, high int) int {
	lowS := strconv.Itoa(low)
	highS := strconv.Itoa(high)
	n := len(highS)
	diffLH := n - len(lowS)

	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, n*18+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int, bool, bool) int
	dfs = func(i, start, diff int, limitLow, limitHigh bool) (res int) {
		if i == n {
			if diff != 0 {
				return 0
			}
			return 1
		}
		if start != -1 && !limitLow && !limitHigh {
			p := &memo[i][start][diff+n*9]
			if *p != -1 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		// 如果前面没有填数字，且剩余数位个数是奇数，那么当前数位不能填数字
		if start < 0 && (n-i)%2 > 0 {
			if lo > 0 {
				return 0 // 必须填数字但 lo > 0，不合法
			}
			return dfs(i+1, start, diff, true, false)
		}

		isLeft := start < 0 || i < (start+n)/2
		for d := lo; d <= hi; d++ {
			newStart := start
			if start < 0 && d > 0 {
				newStart = i // 记录第一个填数字的位置
			}
			newDiff := diff
			if isLeft {
				newDiff += d
			} else {
				newDiff -= d
			}
			res += dfs(i+1, newStart, newDiff, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return dfs(0, -1, 0, true, true)
}

func countSymmetricIntegers1(low, high int) (ans int) {
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		n := len(s)
		if n%2 > 0 {
			continue
		}
		sum := 0
		for _, c := range s[:n/2] {
			sum += int(c)
		}
		for _, c := range s[n/2:] {
			sum -= int(c)
		}
		if sum == 0 {
			ans++
		}
	}
	return
}
