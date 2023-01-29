package main

// https://space.bilibili.com/206214
func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	for k := n - 2; k >= 2; k-- {
		great[k] = append([]int(nil), great[k+1]...)
		for x := nums[k+1] - 1; x > 0; x-- {
			great[k][x]++ // x < nums[k+1]，对于 x，大于它的数的个数 +1
		}
	}

	less := make([]int, n+1)
	for j := 1; j < n-2; j++ {
		for x := nums[j-1] + 1; x <= n; x++ {
			less[x]++ // x > nums[j-1]，对于 x，小于它的数的个数 +1
		}
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				ans += int64(less[nums[k]] * great[k][nums[j]])
			}
		}
	}
	return
}
