package main

// https://space.bilibili.com/206214
func getSum(s []int) (ans int64) {
	// 将 s 改造为 t，这样就不需要分 len(s) 的奇偶来讨论了，因为新数组 t 的每个回文子数组都是奇回文子数组（都有回文中心）
	// s 和 t 的下标转换关系：
	// (si+1)*2 = ti
	// ti/2-1 = si
	// ti 为偶数（2,4,6,...）对应 s 中的奇回文子数组
	// ti 为奇数（3,5,7,...）对应 s 中的偶回文子数组
	t := append(make([]int, 0, len(s)*2+3), -2)
	for _, c := range s {
		t = append(t, -1, c)
	}
	t = append(t, -1, -3)

	// 定义一个奇回文子数组的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余子数组的长度
	// halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子数组的回文半径
	// 具体地，闭区间 [i-halfLen[i]+1, i+halfLen[i]-1] 是 t 上的一个回文子数组
	// 由于 t 中回文子数组的首尾元素一定是 -1，根据下标转换关系，
	// 可以得到其在 s 中对应的回文子数组的区间为 [(i-halfLen[i])/2, (i+halfLen[i])/2-2]，用这个结论去计算子数组和
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	// boxR 表示当前右边界下标最大的回文子数组的右边界下标+1（初始化成任意 <= 0 的数都可以）
	// boxM 为该最大回文子数组的中心位置，二者的关系为 boxR = boxM + halfLen[boxM]
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原数组的首尾元素
		hl := 1 // 注：如果题目比较的是抽象意义的值，单个值可能不满足要求，此时应初始化 hl = 0
		if i < boxR {
			// 记 i 关于 boxM 的对称位置 i'=boxM*2-i
			// 若以 i' 为中心的最长回文子数组范围超出了以 boxM 为中心的回文子数组的范围（即 i+halfLen[i'] >= boxR）
			// 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
			// 否则 halfLen[i] 与 halfLen[i'] 相等
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		// 暴力扩展
		// 算法的复杂度取决于这部分执行的次数
		// 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
		// 因此算法的复杂度 = O(len(t)) = O(len(s))
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl
	}

	sum := make([]int64, len(s)+1)
	for i, x := range s {
		sum[i+1] = sum[i] + int64(x)
	}

	for i := 2; i < len(halfLen); i++ {
		hl := halfLen[i]
		// 见上面注释
		ans = max(ans, sum[(i+hl)/2-1]-sum[(i-hl)/2])
	}
	return
}
