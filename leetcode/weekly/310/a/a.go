package main

// https://space.bilibili.com/206214
func mostFrequentEven(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		if x%2 == 0 {
			cnt[x]++
		}
	}
	if len(cnt) == 0 {
		return -1
	}
	mx := int(-1e9)
	for _, v := range cnt {
		if v > mx {
			mx = v
		}
	}
	ans := int(1e9)
	for x, c := range cnt {
		if c == mx {
			ans = min(ans, x)
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
