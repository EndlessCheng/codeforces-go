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
	fMap := map[int]struct{ maxF, j, maxF2, j2 int }{}
	from := make([]int, n)
	maxF, maxI := 0, 0
	for i := n - 1; i >= 0; i-- {
		w, g := words[i], groups[i]

		// 计算 w 的哈希值
		hash := 0
		for _, ch := range w {
			hash = hash<<5 | int(ch&31)
		}

		f := 0 // 方法一中的 f[i]
		for j := range w {
			h := hash | 31<<(j*5) // 用记号笔把 w[k] 涂黑（置为 11111）
			t := fMap[h]
			if g != groups[t.j] { // 可以从最大值转移过来
				if t.maxF > f {
					f = t.maxF
					from[i] = t.j
				}
			} else { // 只能从次大值转移过来
				if t.maxF2 > f {
					f = t.maxF2
					from[i] = t.j2
				}
			}
		}

		f++
		if f > maxF {
			maxF, maxI = f, i
		}

		// 用 f 更新 fMap[h] 的最大次大
		// 注意要保证最大次大的 group 值不同
		for j := range w {
			h := hash | 31<<(j*5)
			t := fMap[h]
			if f > t.maxF { // 最大值需要更新
				if g != groups[t.j] {
					t.maxF2 = t.maxF // 最大值变成次大值
					t.j2 = t.j
				}
				t.maxF = f
				t.j = i
			} else if f > t.maxF2 && g != groups[t.j] { // 次大值需要更新
				t.maxF2 = f
				t.j2 = i
			}
			fMap[h] = t
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
