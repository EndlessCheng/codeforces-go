package main

func createTargetArray(nums []int, index []int) (ans []int) {
	for i, v := range nums {
		id := index[i]
		tmp := make([]int, len(ans))
		copy(tmp, ans)
		ans = append(ans[:id], v)
		ans = append(ans, tmp[id:]...)
	}
	return
}
