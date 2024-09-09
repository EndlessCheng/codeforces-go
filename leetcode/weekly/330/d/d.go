package main

import "slices"

// https://space.bilibili.com/206214
func countQuadruplets(nums []int) (cnt4 int64) {
	cnt3 := make([]int, len(nums))
	for l := 2; l < len(nums); l++ {
		cnt2 := 0
		for j := 0; j < l; j++ {
			if nums[j] < nums[l] { // 3 < 4
				cnt4 += int64(cnt3[j])
				// 把 j 当作 i，把 l 当作 k，现在 nums[i] < nums[k]，即 1 < 2
				cnt2++
			} else { // 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
				cnt3[j] += cnt2
			}
		}
	}
	return
}

func countQuadruplets2(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	for k := n - 2; k > 0; k-- {
		great[k] = slices.Clone(great[k+1])
		for x := 1; x < nums[k+1]; x++ {
			great[k][x]++
		}
	}

	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			x := nums[k]
			if nums[j] > x {
				ans += int64((x - n + 1 + j + great[j][x]) * great[k][nums[j]])
			}
		}
	}
	return
}
