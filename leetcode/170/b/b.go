package main

func xorQueries(arr []int, queries [][]int) (ans []int) {
	sum := make([]int, len(arr)+1)
	for i, v := range arr {
		sum[i+1] = sum[i] ^ v
	}
	for _, q := range queries {
		l, r := q[0], q[1]
		ans = append(ans, sum[r+1] ^ sum[l])
	}
	return
}
