package main

// https://space.bilibili.com/206214
func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	for k := n - 2; k > 0; k-- {
		great[k] = append([]int(nil), great[k+1]...)
		for x := nums[k+1] - 1; x > 0; x-- {
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
