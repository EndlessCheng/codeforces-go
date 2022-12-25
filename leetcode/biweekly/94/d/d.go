package main

import "strings"

// https://space.bilibili.com/206214
const mod int = 1e9 + 7

func countAnagrams(s string) int {
	ans, mul := 1, 1
	for _, s := range strings.Split(s, " ") {
		cnt := [26]int{}
		for i, c := range s {
			cnt[c-'a']++
			mul = mul * cnt[c-'a'] % mod
			ans = ans * (i + 1) % mod
		}
	}
	return ans * pow(mul, mod-2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
