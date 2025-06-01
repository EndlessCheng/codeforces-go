package main

// https://space.bilibili.com/206214
func checkEqualPartitions(nums []int, target int64) bool {
	tar := int(target)
	u := 1<<len(nums) - 1
	for s := 1; s < u; s++ { // 枚举 u 的非空真子集 s
		mul1, mul2 := 1, 1
		for i, x := range nums {
			if s>>i&1 > 0 { // i 在集合 s 中
				mul1 = min(mul1*x, tar+1) // 与 tar+1 取 min，防止溢出
			} else { // i 在补集中
				mul2 = min(mul2*x, tar+1)
			}
		}
		if mul1 == tar && mul2 == tar {
			return true
		}
	}
	return false
}
