package main

// https://space.bilibili.com/206214
func canAliceWin(nums []int) bool {
	s := 0
	for _, x := range nums {
		if x < 10 {
			s += x
		} else {
			s -= x
		}
	}
	return s != 0
}
