package main

func xorQueries(arr []int, queries [][]int) []int {
	sum := make([]int, len(arr)+1)
	for i, v := range arr {
		sum[i+1] = sum[i] ^ v
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		ans[i] = sum[r+1] ^ sum[l]
	}
	return ans
}
