package main

// https://space.bilibili.com/206214
func subarraySum(nums []int) (ans int) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	for i, num := range nums {
		ans += s[i+1] - s[max(i-num, 0)]
	}
	return
}

func subarraySum3(nums []int) (ans int) {
	diff := make([]int, len(nums)+1)
	for i, num := range nums {
		diff[max(i-num, 0)]++
		diff[i+1]--
	}

	sd := 0
	for i, x := range nums {
		sd += diff[i]
		ans += x * sd
	}
	return
}

func subarraySum1(nums []int) (ans int) {
	for i, num := range nums {
		for _, x := range nums[max(i-num, 0) : i+1] {
			ans += x
		}
	}
	return
}
