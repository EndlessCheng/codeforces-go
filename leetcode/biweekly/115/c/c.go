package main

// https://space.bilibili.com/206214
func ok(s, t string) (diff bool) {
	if len(s) != len(t) {
		return
	}
	for i := range s {
		if s[i] != t[i] {
			if diff { // 汉明距离大于 1
				return false
			}
			diff = true
		}
	}
	return
}

func getWordsInLongestSubsequence1(words []string, groups []int) []string {
	n := len(words)
	f := make([]int, n)
	from := make([]int, n)
	maxI := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 提前比较 f[j] 与 f[i] 的大小，如果 f[j] <= f[i]，就不用执行更耗时的 check 了
			if f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j]) {
				f[i] = f[j]
				from[i] = j
			}
		}
		f[i]++ // 加一写在这里
		if f[i] > f[maxI] {
			maxI = i
		}
	}

	i := maxI
	ans := make([]string, f[i])
	for k := range ans {
		ans[k] = words[i]
		i = from[i]
	}
	return ans
}

// 线性做法
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	type pair struct{ maxF, j int }
	fMap := map[int]pair{}
	from := make([]int, n)
	maxF, maxI := 0, 0
	for i := n - 1; i >= 0; i-- {
		w, g := words[i], groups[i]

		// 计算 w 的哈希值
		hash := 0
		for _, ch := range w {
			hash = hash<<5 | int(ch&31)
		}

		// 计算方法一中的 f[i]
		f := 0
		for j := range w {
			h := hash | 31<<(j*5) // 用记号笔把 w[k] 涂黑（置为 11111）
			t := fMap[h]
			if t.maxF > f && g != groups[t.j] {
				f = t.maxF
				from[i] = t.j
			}
		}

		f++
		if f > maxF {
			maxF, maxI = f, i
		}

		// 用 f 更新 fMap[h]
		for j := range w {
			h := hash | 31<<(j*5)
			if f > fMap[h].maxF {
				fMap[h] = pair{f, i}
			}
		}
	}

	ans := make([]string, maxF)
	i := maxI
	for k := range ans {
		ans[k] = words[i]
		i = from[i]
	}
	return ans
}
