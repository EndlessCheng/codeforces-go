package main

// https://space.bilibili.com/206214
func maxProduct(nums []int, k, limit int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total < abs(k) { // |k| 太大
		return -1
	}

	ans := -1
	type args struct {
		i, s, m    int
		odd, empty bool
	}
	vis := map[args]bool{}
	var dfs func(int, int, int, bool, bool)
	dfs = func(i, s, m int, odd, empty bool) {
		if ans == limit {
			return
		}

		if i == len(nums) {
			if !empty && s == k && m <= limit {
				ans = max(ans, m)
			}
			return
		}

		t := args{i, s, m, odd, empty}
		if vis[t] {
			return
		}
		vis[t] = true

		// 不选 x
		dfs(i+1, s, m, odd, empty)

		// 选 x
		x := nums[i]
		if odd {
			s -= x
		} else {
			s += x
		}
		dfs(i+1, s, min(m*x, limit+1), !odd, false)
	}
	dfs(0, 0, 1, false, true)
	return ans
}

func maxProduct2(nums []int, k, limit int) int {
	// 如果数组和小于 |k|，则返回 -1
	total := 0
	for _, x := range nums {
		total += x
	}
	if total < abs(k) {
		return -1
	}

	// s -> {m}
	oddS := map[int]map[int]struct{}{}
	evenS := map[int]map[int]struct{}{}
	add := func(m map[int]map[int]struct{}, key, val int) {
		if _, ok := m[key]; !ok {
			m[key] = map[int]struct{}{}
		}
		m[key][val] = struct{}{}
	}

	for _, x := range nums {
		// 长为偶数的子序列的计算结果 newEvenS
		newEvenS := map[int]map[int]struct{}{}
		for s, set := range oddS {
			newEvenS[s-x] = map[int]struct{}{}
			for m := range set {
				if m*x <= limit {
					newEvenS[s-x][m*x] = struct{}{}
				}
			}
		}

		// 长为奇数的子序列的计算结果 oddS
		for s, set := range evenS {
			if _, ok := oddS[s+x]; !ok {
				oddS[s+x] = map[int]struct{}{}
			}
			for m := range set {
				if m*x <= limit {
					oddS[s+x][m*x] = struct{}{}
				}
			}
			if x == 0 {
				add(oddS, s, 0)
			}
		}

		// 更新 evenS
		for s, set := range newEvenS {
			if eSet, ok := evenS[s]; ok {
				for m := range set {
					eSet[m] = struct{}{}
				}
			} else {
				evenS[s] = set
			}
			if x == 0 {
				add(evenS, s, 0)
			}
		}

		// 子序列只有一个数的情况
		if x <= limit {
			add(oddS, x, x)
		}

		if set, ok := oddS[k]; ok {
			if _, ok := set[limit]; ok {
				return limit // 提前返回
			}
		}
		if set, ok := evenS[k]; ok {
			if _, ok := set[limit]; ok {
				return limit // 提前返回
			}
		}
	}

	calcMax := func(m map[int]struct{}) int {
		maxVal := -1
		if m != nil {
			for v := range m {
				maxVal = max(maxVal, v)
			}
		}
		return maxVal
	}
	return max(calcMax(oddS[k]), calcMax(evenS[k]))
}

func abs(x int) int { if x < 0 { return -x }; return x }
