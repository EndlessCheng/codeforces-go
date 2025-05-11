package main

// https://space.bilibili.com/206214
func minOperations(nums []int) (ans int) {
	st := nums[:0] // 原地
	for _, x := range nums {
		for len(st) > 0 && x < st[len(st)-1] {
			st = st[:len(st)-1]
			ans++
		}
		// 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
		if len(st) == 0 || x != st[len(st)-1] {
			st = append(st, x)
		}
	}
	if st[0] == 0 { // 0 不需要操作
		ans--
	}
	return ans + len(st)
}
