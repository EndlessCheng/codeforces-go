package main

// https://space.bilibili.com/206214
func continuousSubarrays(nums []int) (ans int64) {
	var minQ, maxQ []int
	left := 0
	for right, x := range nums {
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)
		
		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		for nums[maxQ[0]]-nums[minQ[0]] > 2 {
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}
		ans += int64(right - left + 1)
	}
	return
}
