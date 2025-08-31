package main

// https://space.bilibili.com/206214
func recoverOrder(order, friends []int) []int {
	n := len(order)
	isFriend := make([]bool, n+1)
	for _, x := range friends {
		isFriend[x] = true
	}

	ans := make([]int, 0, len(friends)) // 预分配空间
	for _, x := range order {
		if isFriend[x] {
			ans = append(ans, x)
		}
	}
	return ans
}
