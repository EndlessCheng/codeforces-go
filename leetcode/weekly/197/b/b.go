package main

func numSub(s string) (ans int) {
	const mod = 1_000_000_007
	last0 := -1
	for i, ch := range s {
		if ch == '0' {
			last0 = i // 记录上个 0 的位置
		} else {
			ans += i - last0 // 右端点为 i 的全 1 子串个数
		}
	}
	return ans % mod
}
