package main

// https://space.bilibili.com/206214
func kthDigit(K int64) int {
	k := int(K - 1) // k 改成从 0 开始，方便计算

	// 十进制长为 length 的正整数有 cnt = 9 * 10^(length-1) 个
	cnt, length := 9, 1
	for cnt*length <= k {
		k -= cnt * length // 这里减小了 k
		cnt *= 10
		length++
	}

	// k 在正整数 x 中
	x := cnt/9 + k/length
	if x/10%2 > 0 {
		// 改成递减顺序，例如 10 变成 19，11 变成 18 ……
		x += 9 - x%10*2
	}

	// 计算 x 从高到低第 k%length（从 0 开始）个数字
	for range length - k%length - 1 {
		x /= 10
	}
	return x % 10
}
