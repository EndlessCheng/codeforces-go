package main

// 同 1004. 最大连续1的个数 III https://leetcode-cn.com/problems/max-consecutive-ones-iii/

// github.com/EndlessCheng/codeforces-go
func maxConsecutiveAnswers(answerKey string, k int) int {
	calc := func(target byte) (ans int) {
		var left, lsum, rsum int
		for right := range answerKey {
			if answerKey[right] != target {
				rsum++
			}
			for lsum < rsum-k {
				if answerKey[left] != target {
					lsum++
				}
				left++
			}
			ans = max(ans, right-left+1)
		}
		return
	}
	return max(calc('T'), calc('F'))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
