package main

// https://space.bilibili.com/206214
var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var nsq2mask = [31]int{} // nsq2mask[i] 为 i 对应的质数集合（用二进制表示）

func init() {
	for i := 2; i <= 30; i++ {
		for j, p := range primes {
			if i%p == 0 {
				if i%(p*p) == 0 { // 有平方因子
					nsq2mask[i] = -1
					break
				}
				nsq2mask[i] |= 1 << j // 把 j 加到集合中
			}
		}
	}
}

func squareFreeSubsets(a []int) int {
	const mod int = 1e9 + 7
	cnt, pow2 := [31]int{}, 1
	for _, v := range a {
		if v == 1 {
			pow2 = pow2 * 2 % mod
		} else {
			cnt[v]++
		}
	}

	const m = 1 << len(primes)
	f := [m]int{1} // f[j] 表示恰好组成集合 j 的方案数，其中空集的方案数为 1
	for nsq, mask := range nsq2mask {
		if mask > 0 && cnt[nsq] > 0 {
			other := (m - 1) ^ mask // mask 的补集
			for j := other; ; { // 枚举 other 的子集 j
				f[j|mask] = (f[j|mask] + f[j]*cnt[nsq]) % mod // 不选 mask + 选 mask
				j = (j - 1) & other
				if j == other {
					break
				}
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans += v
	}
	return (ans%mod*pow2 - 1 + mod) % mod // -1 去掉空集，+mod 保证非负
}
