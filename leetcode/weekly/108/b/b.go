package main

// https://space.bilibili.com/206214
func numSubarraysWithSum(nums []int, goal int) (ans int) {
	sum1, sum2 := 0, 0
	left1, left2 := 0, 0
	for i, x := range nums {
		sum1 += x
		for left1 <= i && sum1 >= goal { // 避免 goal = 0 的情况下标越界
			sum1 -= nums[left1]
			left1++
		}
		ans += left1 // 先加上 >= goal 的子数组个数

		sum2 += x
		for sum2 > goal {
			sum2 -= nums[left2]
			left2++
		}
		ans -= left2 // 再减去 > goal 的，剩下的就是 = goal 的
	}
	return
}
