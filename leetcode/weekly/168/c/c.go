package main

func maxFreq(s string, maxLetters, minSize, _ int) (ans int) {
	cntStr := map[string]int{}
	cntChar := [26]int{}
	kinds := 0
	for i, b := range s {
		// 1. 进入窗口
		if cntChar[b-'a'] == 0 {
			kinds++
		}
		cntChar[b-'a']++

		left := i - minSize + 1
		if left < 0 { // 窗口大小不足 minSize
			continue
		}

		// 2. 更新统计量
		if kinds <= maxLetters {
			cntStr[s[left:i+1]]++
		}

		// 3. 离开窗口，为下一个循环做准备
		out := s[left]
		cntChar[out-'a']--
		if cntChar[out-'a'] == 0 {
			kinds--
		}
	}

	for _, cnt := range cntStr {
		ans = max(ans, cnt)
	}
	return
}
