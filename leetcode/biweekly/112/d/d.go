package main

import "sort"

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	cc := map[int]int{}
	for _, c := range cnt {
		if c > 0 {
			cc[c]++
		}
	}

	type KV struct{ cnt, num int }
	kv := make([]KV, 0, len(cc))
	for k, v := range cc {
		kv = append(kv, KV{k, v})
	}
	sort.Slice(kv, func(i, j int) bool { return kv[i].cnt > kv[j].cnt })

	ans := 1
	for _, p := range kv {
		if p.num >= k {
			return ans * pow(p.cnt, k) % mod * comb(p.num, k) % mod
		}
		ans = ans * pow(p.cnt, p.num) % mod
		k -= p.num
	}
	return 0 // k 太大，无法选 k 个不一样的字符
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func comb(n, k int) int {
	res := n
	for i := 2; i <= k; i++ {
		res = res * (n - i + 1) / i
	}
	return res % mod
}
