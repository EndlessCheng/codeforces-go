package main

// https://space.bilibili.com/206214
// 计算除了 x 以外的出现次数之和 sum，出现次数最大值 mx
func getSumAndMax(cnt []int, x byte) (sum, mx int) {
	for i, c := range cnt {
		if i != int(x-'a') {
			sum += c
			mx = max(mx, c)
		}
	}
	return
}

// 计算这一组在得到 k 个 xx 后的得分
func calcScore(sum, mx, k int) int {
	sum += k
	mx = max(mx, k)
	return min(sum/2, sum-mx)
}

func score(cards []string, x byte) (ans int) {
	var cnt1, cnt2 [10]int // 题目保证只有前 10 个小写字母
	for _, s := range cards {
		// 统计形如 x? 的每个 ? 的出现次数
		if s[0] == x {
			cnt1[s[1]-'a']++
		}
		// 统计形如 ?x 的每个 ? 的出现次数
		if s[1] == x {
			cnt2[s[0]-'a']++
		}
	}

	sum1, max1 := getSumAndMax(cnt1[:], x)
	sum2, max2 := getSumAndMax(cnt2[:], x)

	cntXX := cnt1[x-'a']
	// 枚举分配 k 个 xx 给第一组，其余的 xx 给第二组
	for k := range cntXX + 1 {
		score1 := calcScore(sum1, max1, k)
		score2 := calcScore(sum2, max2, cntXX-k)
		ans = max(ans, score1+score2)
	}
	return
}
