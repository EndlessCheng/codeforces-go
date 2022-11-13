package main

// https://space.bilibili.com/206214
// 点评：这题和【70. 爬楼梯】有什么区别呢？不就是把那道题的 1 和 2 替换成了 zero 和 one 嘛！
// 如果这题你没有做出来，建议把【70. 爬楼梯】认真做一遍。
func countGoodStrings(low, high, zero, one int) (ans int) {
	const mod int = 1e9 + 7
	f := make([]int, high+1) // f[i] 表示构造长为 i 的字符串的方案数
	f[0] = 1 // 构造空串的方案数为 1
	for i := 1; i <= high; i++ {
		if i >= one  { f[i] = (f[i] + f[i-one]) % mod }
		if i >= zero { f[i] = (f[i] + f[i-zero]) % mod }
		if i >= low  { ans = (ans + f[i]) % mod }
	}
	return
}
