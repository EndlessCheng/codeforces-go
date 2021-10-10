package main

// 位运算写法

// github.com/EndlessCheng/codeforces-go
func twoOutOfThree(nums1, nums2, nums3 []int) (ans []int) {
	mask := map[int]int8{}
	for i, a := range [][]int{nums1, nums2, nums3} {
		for _, v := range a {
			mask[v] |= 1 << i
		}
	}
	for v, m := range mask {
		if m > 2 && m != 4 { // 二进制包含至少两个 1，即 3 5 6 7。也可以用 bits.OnesCount8 来求
			ans = append(ans, v)
		}
	}
	return
}
