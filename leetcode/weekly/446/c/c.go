package main

// https://space.bilibili.com/206214
func resultArray2(nums []int, k int) []int64 {
	ans := make([]int64, k)
	f := make([]int, k)
	for _, v := range nums {
		nf := make([]int, k)
		nf[v%k] = 1
		for y, c := range f {
			nf[y*v%k] += c
		}
		f = nf
		for x, c := range f {
			ans[x] += int64(c)
		}
	}
	return ans
}

func resultArray(nums []int, k int) []int64 {
	ans := make([]int64, k)
	f := make([][]int, len(nums)+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	for i, v := range nums {
		f[i+1][v%k] = 1
		for y, c := range f[i] {
			f[i+1][y*v%k] += c
		}
		for x, c := range f[i+1] {
			ans[x] += int64(c)
		}
	}
	return ans
}
