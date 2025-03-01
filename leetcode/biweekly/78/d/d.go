package main

import "math"

func largestVariance(s string) (ans int) {
	var f0, f1 [26][26]int
	for i := range f1 {
		for j := range f1[i] {
			f1[i][j] = math.MinInt
		}
	}

	for _, ch := range s {
		ch -= 'a'
		// 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
		for i := range 26 {
			if i == int(ch) {
				continue
			}
			// 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
			f0[ch][i] = max(f0[ch][i], 0) + 1
			f1[ch][i] += 1
			// 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
			f0[i][ch] = max(f0[i][ch], 0) - 1
			f1[i][ch] = f0[i][ch]
			ans = max(ans, f1[ch][i], f1[i][ch])
		}
	}
	return
}

func largestVariance1(s string) (ans int) {
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			f0, f1 := 0, math.MinInt
			for _, ch := range s {
				if ch == a {
					f0 = max(f0, 0) + 1
					f1++
				} else if ch == b {
					f1, f0 = max(f0, 0)-1, max(f0, 0)-1
				} // else { f0 = max(f0, 0) } 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
				ans = max(ans, f1)
			}
		}
	}
	return
}
