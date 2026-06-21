package main

// https://space.bilibili.com/206214
func countValidSubarrays1(nums []int, x int) (ans int) {
	// 枚举子数组的左右端点
	for i := range nums {
		sum := 0
		for _, v := range nums[i:] {
			sum += v
			// 计算 sum 的最低位
			if sum%10 != x {
				continue
			}
			// 计算 sum 的最高位
			s := sum
			for s > 9 {
				s /= 10
			}
			if s == x {
				ans++
			}
		}
	}
	return
}

func countValidSubarrays(nums []int, x int) (ans int) {
	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	// 枚举子数组和的十进制长度
	for low, high := x, x+1; low <= sum[n]; low, high = low*10, high*10 {
		// 计算子数组和在 [low, high-1] 中，且子数组和模 10 为 x 的子数组个数
		cnt := [10]int{}
		left1, left2 := 0, 0
		for _, s := range sum {
			// 随着 s 的增大，<= s-high 的前缀和离开窗口，<= s-low 的前缀和进入窗口
			for sum[left1] <= s-high {
				cnt[sum[left1]%10]--
				left1++
			}
			for sum[left2] <= s-low {
				cnt[sum[left2]%10]++
				left2++
			}
			ans += cnt[(s-x+10)%10]
		}
	}

	return
}
