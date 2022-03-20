package main

// github.com/EndlessCheng/codeforces-go
func maximumBobPoints(numArrows int, aliceArrows []int) (ans []int) {
	for i, maxScore := 0, -1; i < 1<<len(aliceArrows); i++ { // 二进制枚举
		score, arrow, bobArrows := 0, 0, [12]int{}
		for j, v := range aliceArrows {
			if i>>j&1 == 1 {
				score += j
				arrow += v + 1
				bobArrows[j] = v + 1 // Bob 多射一支箭
			}
		}
		if arrow > numArrows {
			continue
		}
		if score > maxScore {
			maxScore = score
			bobArrows[0] += numArrows - arrow // 随便找个位置补满至 numArrows
			ans = bobArrows[:]
		}
	}
	return
}
