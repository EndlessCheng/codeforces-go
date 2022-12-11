package main

// https://space.bilibili.com/206214
func minimumTotalCost(nums1, nums2 []int) (ans int64) {
	var swapCnt, modeCnt, mode int
	cnt := make([]int, len(nums1)+1)
	for i, x := range nums1 {
		if x == nums2[i] {
			ans += int64(i)
			swapCnt++
			cnt[x]++
			if cnt[x] > modeCnt {
				modeCnt, mode = cnt[x], x
			}
		}
	}

	for i, x := range nums1 {
		if modeCnt*2 <= swapCnt {
			break
		}
		if x != nums2[i] && x != mode && nums2[i] != mode {
			ans += int64(i)
			swapCnt++
		}
	}
	if modeCnt*2 > swapCnt {
		return -1
	}
	return
}
