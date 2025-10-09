package main

// https://space.bilibili.com/206214
func findSmallestInteger1(nums []int, m int) (mex int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[(x%m+m)%m]++
	}
	for cnt[mex%m] > 0 {
		cnt[mex%m]--
		mex++
	}
	return
}

func findSmallestInteger(nums []int, m int) int {
	cnt := make([]int, m)
	for _, x := range nums {
		cnt[(x%m+m)%m]++
	}

	i := 0
	for j := 1; j < m; j++ {
		if cnt[j] < cnt[i] {
			i = j
		}
	}

	return m*cnt[i] + i
}
