package main

func createTargetArray(nums []int, index []int) (ans []int) {
	for i, v := range nums {
		id := index[i]
		ans = append(ans, 0)
		copy(ans[id+1:], ans[id:])
		ans[id] = v
	}
	return
}
