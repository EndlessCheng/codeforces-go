package main

// https://space.bilibili.com/206214
func maximumWidth(planks []int) (ans int) {
	// 统计 planks 的元素出现次数
	cnt := map[int]int{}
	for _, x := range planks {
		cnt[x]++
	}

	// 枚举所有高度对 (x,y)
	cntPair := map[int]int{}
	for x, c := range cnt {
		cntPair[x] += c // 方便后面统计
		cntPair[x*2] += c / 2 // 高为 x 的木板内部配对
		for y, c2 := range cnt {
			if y > x { // 避免 x+y 和 y+x 重复统计
				cntPair[x+y] += min(c, c2)
			}
		}
	}

	// 枚举最终木板高度
	for _, c := range cntPair {
		ans = max(ans, c)
	}

	return
}
