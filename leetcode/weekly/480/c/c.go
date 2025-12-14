package main

// https://space.bilibili.com/206214
func minMoves(balance []int) int64 {
	negIdx := -1
	total := 0
	for i, x := range balance {
		total += x
		if x < 0 {
			negIdx = i
		}
	}
	if total < 0 { // 总和必须非负
		return -1
	}
	if negIdx < 0 { // 没有负数，无需操作
		return 0
	}

	n := len(balance)
	need := -balance[negIdx]
	ans := 0
	for dis := 1; ; dis++ { // 把与 negIdx 相距 dis 的数移到 negIdx
		s := balance[(negIdx-dis+n)%n] + balance[(negIdx+dis)%n]
		if s >= need {
			ans += need * dis // need 个 1 移动 dis 次
			return int64(ans)
		}
		ans += s * dis // s 个 1 移动 dis 次
		need -= s
	}
}
