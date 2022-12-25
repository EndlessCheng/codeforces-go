package main

// https://space.bilibili.com/206214
func countPartitions(nums []int, k int) int {
	const mod int = 1e9 + 7
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum < k*2 {
		return 0
	}
	ans := 1
	f := make([]int, k)
	f[0] = 1
	for _, x := range nums {
		ans = ans * 2 % mod
		for j := k - 1; j >= x; j-- {
			f[j] = (f[j] + f[j-x]) % mod
		}
	}
	for _, x := range f {
		ans -= x * 2
	}
	return (ans%mod + mod) % mod // 保证答案非负
}
