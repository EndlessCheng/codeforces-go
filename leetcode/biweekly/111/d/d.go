package main

import "strconv"

// https://space.bilibili.com/206214
func numberOfBeautifulIntegers(low, high, k int) int {
	calc := func(high int) int {
		s := strconv.Itoa(high)
		n := len(s)
		memo := make([][][]int, n)
		for i := range memo {
			memo[i] = make([][]int, k+1)
			for j := range memo[i] {
				memo[i][j] = make([]int, n*2+1)
				for k := range memo[i][j] {
					memo[i][j][k] = -1 // -1 表示没有计算过
				}
			}
		}
		var dfs func(int, int, int, bool, bool) int
		dfs = func(i, val, diff int, isLimit, isNum bool) (res int) {
			if i == n {
				if isNum && val == 0 && diff == n {
					return 1 // 找到了一个合法数字
				}
				return 0
			}
			if !isLimit && isNum {
				p := &memo[i][val][diff]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }() // 记忆化搜索
			}
			if !isNum { // 可以跳过当前数位
				res += dfs(i+1, val, diff, false, false)
			}
			up := 9
			if isLimit {
				up = int(s[i] - '0') // 如果前面填的数字都和 high 的一样，那么这一位至多填数字 s[i]（否则就超过 high 啦）
			}
			d := 0
			if !isNum {
				d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
			}
			for ; d <= up; d++ { // 枚举要填入的数字 d
				res += dfs(i+1, (val*10+d)%k, diff+d%2*2-1, isLimit && d == up, true)
			}
			return
		}
		return dfs(0, 0, n, true, false)
	}
	return calc(high) - calc(low-1)
}
