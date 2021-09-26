package main

// github.com/EndlessCheng/codeforces-go
func scoreOfStudents(s string, answers []int) (ans int) {
	n := len(s)
	// 计算正确答案
	correct := 0
	for i := 0; i < n; {
		mul := int(s[i] & 15)
		for i += 2; i < n && s[i-1] == '*'; i += 2 {
			mul *= int(s[i] & 15)
		}
		correct += mul
	}

	// 区间 DP
	// dp[l][r] 表示 s[l..r] 内的表达式在不同计算顺序下的不超过 1000 的所有可能值
	dp := make([][]map[int]bool, n)
	for i := range dp {
		dp[i] = make([]map[int]bool, n)
	}
	var f func(int, int) map[int]bool
	f = func(l, r int) map[int]bool {
		if l == r {
			return map[int]bool{int(s[l] & 15): true}
		}
		if dp[l][r] != nil {
			return dp[l][r]
		}
		res := map[int]bool{}
		for i := l + 1; i < r; i += 2 { // 枚举分界点
			for v := range f(l, i-1) { // 计算左边表达式的可能值
				for w := range f(i+1, r) { // 计算右边表达式的可能值
					x := v + w
					if s[i] == '*' {
						x = v * w
					}
					if x <= 1000 { // 超过 1000 的结果不计入
						res[x] = true
					}
				}
			}
		}
		dp[l][r] = res
		return res
	}
	res := f(0, n-1)
	for _, v := range answers {
		if v == correct {
			ans += 5
		} else if res[v] {
			ans += 2
		}
	}
	return
}
