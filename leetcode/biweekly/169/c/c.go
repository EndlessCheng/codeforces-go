package main

// https://space.bilibili.com/206214
func longestSubarray1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	suf := make([]int, n)
	suf[n-1] = 1
	ans := 2
	for i := n - 2; i > 0; i-- {
		if nums[i] <= nums[i+1] {
			suf[i] = suf[i+1] + 1
			ans = max(ans, suf[i]+1) // 把 nums[i-1] 拼在 suf[i] 前面
		} else {
			suf[i] = 1
		}
	}

	pre := 1
	for i := 1; i < n-1; i++ {
		if nums[i-1] <= nums[i+1] {
			ans = max(ans, pre+1+suf[i+1]) // 替换 nums[i]
		}
		if nums[i-1] <= nums[i] {
			pre++
			ans = max(ans, pre+1) // 把 nums[i+1] 拼在 pre 后面
		} else {
			pre = 1
		}
	}
	return ans
}

func longestSubarray2(nums []int) int {
	n := len(nums)
	f := make([][2]int, n)
	f[0] = [2]int{1, 1}

	ans := 1 // 以 nums[0] 结尾的子数组长度
	for i := 1; i < n; i++ {
		if nums[i-1] <= nums[i] {
			f[i][0] = f[i-1][0] + 1
			f[i][1] = f[i-1][1] + 1
		} else {
			f[i][0] = 1
			// 不需要写 f[i][1] = 1，因为下面算出来的值至少是 2
		}

		if i >= 2 && nums[i-2] <= nums[i] {
			f[i][1] = max(f[i][1], f[i-2][0]+2)
		} else {
			f[i][1] = max(f[i][1], 2)
		}

		ans = max(ans, f[i-1][0]+1, f[i][1])
	}
	return ans
}

func longestSubarray3(nums []int) int {
	pre0, f0, f1 := 0, 1, 1

	ans := 1 // 以 nums[0] 结尾的子数组长度
	for i := 1; i < len(nums); i++ {
		tmp := f0
		if nums[i-1] <= nums[i] {
			f0++
			f1++
		} else {
			f0 = 1
			f1 = 0 // 清除旧数据
		}

		if i >= 2 && nums[i-2] <= nums[i] {
			f1 = max(f1, pre0+2)
		} else {
			f1 = max(f1, 2)
		}

		ans = max(ans, tmp+1, f1)
		pre0 = tmp
	}
	return ans
}

func longestSubarray(nums []int) int {
	n := len(nums)
	ans := min(n, 2)
	for i := 1; i < n; {
		if nums[i-1] > nums[i] {
			i++
			continue
		}

		// 枚举 i-1 和 i 作为非递减子数组的前两项
		start := i - 1
		// 往右移动，直到 nums[i] 不满足非递减
		for i++; i < n && nums[i-1] <= nums[i]; i++ {
		}

		// 现在 [start, i-1] 是非递减子数组
		// 要想让子数组更长，要么改左边的 nums[start-1]，要么改右边的 nums[i] 或者 nums[i-1]

		// 改 nums[start-1]
		ans = max(ans, i-max(start-1, 0)) // 非递减子数组 [max(start-1,0), i-1]
		// 继续往左延长的情况等同于上一段继续往右延长，无需重复计算

		if i == n {
			break
		}

		// 改 nums[i] 或者 nums[i-1]
		if i < n-1 && (nums[i-1] <= nums[i+1] || nums[i-2] <= nums[i] && nums[i] <= nums[i+1]) { // 可以和 nums[i+1] 连起来
			// 继续往右延长
			j := i + 2
			for ; j < n && nums[j-1] <= nums[j]; j++ {
			}
			ans = max(ans, j-start) // 非递减子数组 [start, j-1]
		} else { // 子数组右端点最远只能到 i
			ans = max(ans, i-start+1) // 非递减子数组 [start, i]
		}
	}
	return ans
}
