package main

func maxProduct(nums []int) (ans int) {
	mx := 0
	for _, x := range nums {
		ans = max(ans, (mx-1)*(x-1))
		mx = max(mx, x)
	}
	return
}
