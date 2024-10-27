package main

// https://space.bilibili.com/206214
func possibleStringCount(word string, k int) int {
	// 无法满足要求
	if len(word) < k {
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			if len(cnts) < k { // 保证空间复杂度为 O(k)
				cnts = append(cnts, cnt)
			}
			ans = ans * cnt % mod
			cnt = 0
		}
	}
	// 任何输入的字符串都至少为 k
	m := len(cnts)
	if m >= k {
		return ans
	}

	f := make([]int, k)
	f[0] = 1
	for i, c := range cnts {
		// 原地计算 f 的前缀和
		for j := 1; j < k; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		// 计算子数组和
		for j := k - 1; j > i; j-- {
			f[j] = f[j-1]
			if j > c {
				f[j] -= f[j-c-1]
			}
		}
		f[i] = 0
	}

	for _, v := range f[m:] {
		ans -= v
	}
	return (ans%mod + mod) % mod // 保证结果非负
}

func possibleStringCount2(word string, k int) int {
	// 无法满足要求
	if len(word) < k {
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			if len(cnts) < k { // 保证空间复杂度为 O(k)
				cnts = append(cnts, cnt)
			}
			ans = ans * cnt % mod
			cnt = 0
		}
	}
	// 任何输入的字符串都至少为 k
	m := len(cnts)
	if m >= k {
		return ans
	}

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	f[0][0] = 1
	s := make([]int, k+1)
	for i, c := range cnts {
		for j, v := range f[i] {
			s[j+1] = (s[j] + v) % mod
		}
		// j <= i 的 f[i][j] 都是 0
		for j := i + 1; j < k; j++ {
			f[i+1][j] = s[j] - s[max(j-c, 0)]
		}
	}

	for _, v := range f[m][m:] {
		ans -= v
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
