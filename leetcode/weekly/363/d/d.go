package main

// https://space.bilibili.com/206214
func core(n int) int {
	for i := 2; i*i <= n; i++ {
		for n%(i*i) == 0 {
			n /= i * i
		}
	}
	return n
}

func maximumSum(nums []int) (ans int64) {
	sum := make([]int64, len(nums)+1)
	for i, x := range nums {
		c := core(i + 1)
		sum[c] += int64(x)
		ans = max(ans, sum[c])
	}
	return
}

func maximumSum2(nums []int) (ans int64) {
	n := len(nums)
	for i := 1; i <= n; i++ {
		sum := int64(0)
		for j := 1; i*j*j <= n; j++ {
			sum += int64(nums[i*j*j-1])
		}
		ans = max(ans, sum)
	}
	return
}
