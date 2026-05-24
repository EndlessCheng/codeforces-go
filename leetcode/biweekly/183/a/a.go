package main

// https://space.bilibili.com/206214
func minimumSwaps1(nums []int) (ans int) {
	cnt := 0
	for _, x := range nums {
		if x != 0 {
			cnt++
		}
	}

	for _, x := range nums[:cnt] {
		if x == 0 {
			ans++
		}
	}

	return
}

func minimumSwaps(nums []int) (ans int) {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] != 0 {
			l++
		} else if nums[r] == 0 {
			r--
		} else {
			// 交换 nums[l] 和 nums[r]
			ans++
			l++
			r--
		}
	}
	return
}
