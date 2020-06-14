package main

func runningSum(a []int) (ans []int) {
	sum := make([]int, len(a)+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	return sum[1:]
}
