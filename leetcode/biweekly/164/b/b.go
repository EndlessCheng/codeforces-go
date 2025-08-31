package main

// https://space.bilibili.com/206214
// 计算这一组的得分（配对个数），以及剩余元素个数
func calc(cnt []int, x byte) (int, int) {
	sum, mx := 0, 0
	for i, c := range cnt {
		if i != int(x-'a') {
			sum += c
			mx = max(mx, c)
		}
	}
	pairs := min(sum/2, sum-mx)
	return pairs, sum - pairs*2
}

func score(cards []string, x byte) int {
	var cnt1, cnt2 [10]int
	for _, s := range cards {
		if s[0] == x {
			cnt1[s[1]-'a']++
		}
		if s[1] == x {
			cnt2[s[0]-'a']++
		}
	}

	pairs1, left1 := calc(cnt1[:], x)
	pairs2, left2 := calc(cnt2[:], x)
	ans := pairs1 + pairs2 // 不考虑 xx 时的得分

	cntXX := cnt1[x-'a']
	// 把 xx 和剩下的 x? 和 ?x 配对
	// 每有 1 个 xx，得分就能增加一，但这不能超过剩下的 x? 和 ?x 的个数 left1+left2
	if cntXX > 0 {
		mn := min(cntXX, left1+left2)
		ans += mn
		cntXX -= mn
	}

	// 如果还有 xx，就撤销之前的配对，比如 (ax,bx) 改成 (ax,xx) 和 (bx,xx)
	// 每有 2 个 xx，得分就能增加一，但这不能超过之前的配对个数 pairs1+pairs2
	// 由于这种方案平均每个 xx 只能增加 0.5 分，不如上面的，所以先考虑把 xx 和剩下的 x? 和 ?x 配对，再考虑撤销之前的配对
	if cntXX > 0 {
		ans += min(cntXX/2, pairs1+pairs2)
	}

	return ans
}
