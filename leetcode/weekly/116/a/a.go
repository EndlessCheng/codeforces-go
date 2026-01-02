package main

import "math/rand"

// https://space.bilibili.com/206214
func repeatedNTimes1(nums []int) int {
	seen := map[int]struct{}{}
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return x
		}
		seen[x] = struct{}{}
	}
	panic(-1)
}

func repeatedNTimes2(nums []int) (ans int) {
	hp := 0
	for _, x := range nums[1:] {
		if x == nums[0] {
			return x
		}
		if hp == 0 { // x 是初始擂主，生命值为 1
			ans, hp = x, 1
		} else if x == ans { // 比武，同门加血，否则扣血
			hp++
		} else {
			hp--
		}
	}
	return
}

func repeatedNTimes3(nums []int) int {
	for i := 1; ; i++ {
		x := nums[i]
		if x == nums[i-1] ||
			i > 1 && x == nums[i-2] ||
			i > 2 && x == nums[i-3] {
			return x
		}
	}
}

func repeatedNTimes(nums []int) int {
	n := len(nums)
	for {
		// 在 [0, n-1] 中随机生成两个不同下标
		i := rand.Intn(n)
		j := rand.Intn(n - 1)
		if j >= i {
			j++
		}
		if nums[i] == nums[j] {
			return nums[i]
		}
	}
}
