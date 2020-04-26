package main

func findDiagonalOrder(mat [][]int) (ans []int) {
	order := [1e5][]int{}
	for i := len(mat) - 1; i >= 0; i-- {
		for j, v := range mat[i] {
			order[i+j] = append(order[i+j], v)
		}
	}
	for _, a := range order {
		for _, v := range a {
			ans = append(ans, v)
		}
	}
	return
}
