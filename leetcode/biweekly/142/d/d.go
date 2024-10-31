package main

// https://space.bilibili.com/206214
func possibleStringCount(word string, k int) int {
	if len(word) < k { // 无法满足要求
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			// 如果 cnt = 1，这组字符串必选，无需参与计算
			if cnt > 1 {
				if k > 0 { // 保证空间复杂度为 O(k)
					cnts = append(cnts, cnt-1)
				}
				ans = ans * cnt % mod
			}
			k-- // 注意这里把 k 减小了
			cnt = 0
		}
	}

	if k <= 0 {
		return ans
	}

	f := make([]int, k)
	f[0] = 1
	for _, c := range cnts {
		// 原地计算 f 的前缀和
		for j := 1; j < k; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		// 计算子数组和
		for j := k - 1; j > c; j-- {
			f[j] -= f[j-c-1]
		}
	}

	for _, x := range f {
		ans -= x
	}
	return (ans%mod + mod) % mod // 保证结果非负
}

func possibleStringCount2(word string, k int) int {
	if len(word) < k { // 无法满足要求
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			// 如果 cnt = 1，这组字符串必选，无需参与计算
			if cnt > 1 {
				if k > 0 {
					cnts = append(cnts, cnt-1)
				}
				ans = ans * cnt % mod
			}
			k-- // 注意这里把 k 减小了
			cnt = 0
		}
	}

	if k <= 0 {
		return ans
	}

	m := len(cnts)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	f[0][0] = 1
	s := make([]int, k+1)
	for i, c := range cnts {
		// 计算 f[i] 的前缀和数组 s
		for j, v := range f[i] {
			s[j+1] = (s[j] + v) % mod
		}
		// 计算子数组和
		for j := range f[i+1] {
			f[i+1][j] = s[j+1] - s[max(j-c, 0)]
		}
	}

	for _, v := range f[m] {
		ans -= v
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
