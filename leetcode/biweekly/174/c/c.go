package main

// https://space.bilibili.com/206214
func alternatingXOR1(nums []int, target1, target2 int) (ans int) {
	const mod = 1_000_000_007
	targets := []int{0, target1, target1 ^ target2, target2}

	n := len(nums)
	f := [4]int{1}
	preSum := 0

	for _, x := range nums[:n-1] {
		preSum ^= x
		tmp3 := f[3]
		if preSum == target2 {
			f[3] = (f[3] + f[2]) % mod
		}
		if preSum == target1^target2 {
			f[2] = (f[2] + f[1]) % mod
		}
		if preSum == target1 {
			f[1] = (f[1] + f[0]) % mod
		}
		if preSum == 0 {
			f[0] = (f[0] + tmp3) % mod
		}
	}

	preSum ^= nums[n-1]
	for j, t := range targets {
		if preSum == t {
			ans += f[(j+3)%4]
		}
	}
	return ans % mod
}

func alternatingXOR(nums []int, target1, target2 int) int {
	const mod = 1_000_000_007
	f1 := map[int]int{}
	f2 := map[int]int{0: 1}
	preSum := 0
	for i, x := range nums {
		preSum ^= x
		last1 := f2[preSum^target1] // [0,i] 的最后一段的异或和是 target1 的方案数
		last2 := f1[preSum^target2] // [0,i] 的最后一段的异或和是 target2 的方案数
		if i == len(nums)-1 {
			return (last1 + last2) % mod
		}
		f1[preSum] = (f1[preSum] + last1) % mod
		f2[preSum] = (f2[preSum] + last2) % mod
	}
	panic("unreachable")
}
