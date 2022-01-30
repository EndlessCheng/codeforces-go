package main

// github.com/EndlessCheng/codeforces-go
func subStrHash(s string, prime, mod, k, hashValue int) (ans string) {
	hash, pow := 0, 1
	i, n := len(s)-1, len(s)
	for ; i >= n-k; i-- {
		hash = (hash*prime + int(s[i]&31)) % mod // 计算 s[n-k:] 的哈希值
		pow = pow * prime % mod
	}
	if hash == hashValue {
		ans = s[n-k:]
	}
	for ; i >= 0; i-- { // 倒着向前滑动窗口
		hash = (hash*prime + int(s[i]&31) + mod - pow*int(s[i+k]&31)%mod) % mod // 计算新哈希值
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return
}
