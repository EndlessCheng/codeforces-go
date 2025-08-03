package main

// https://space.bilibili.com/206214
func maxBalancedShipments(weight []int) (ans int) {
	for i := 1; i < len(weight); i++ {
		if weight[i-1] > weight[i] {
			ans++
			i++ // 每个装运至少要有两个包裹
		}
	}
	return
}
