package main

// github.com/EndlessCheng/codeforces-go
func subStrHash(s string, power, mod, k, hashValue int) (ans string) {
	n := len(s)
	p := pow(power, k-1, mod)
	hash := 0
	for i := n - 1; i >= 0; i-- { // 倒着滑窗
		// 1. 左端点进入窗口
		hash = (hash*power + int(s[i]&31)) % mod
		right := i + k - 1 // 窗口右端点
		if right >= n {    // 窗口大小不足 k，尚未形成第一个窗口
			continue
		}
		// 2. 更新答案
		if hash == hashValue {
			ans = s[i : right+1]
		}
		// 3. 右端点离开窗口，为下一个循环做准备
		hash = (hash - int(s[right]&31)*p%mod + mod) % mod // +mod 保证结果非负
	}
	return
}

func pow(x, n, mod int) int {
	res := 1 % mod
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func subStrHash2(s string, power, mod, k, hashValue int) (ans string) {
	hash, p := 0, 1
	i, n := len(s)-1, len(s)
	for ; i >= n-k; i-- {
		hash = (hash*power + int(s[i]&31)) % mod // 计算 s[n-k:] 的哈希值
		p = p * power % mod
	}
	if hash == hashValue {
		ans = s[n-k:]
	}
	for ; i >= 0; i-- { // 倒着向前滑动窗口
		hash = (hash*power + int(s[i]&31) + mod - p*int(s[i+k]&31)%mod) % mod // 计算新哈希值
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return
}
