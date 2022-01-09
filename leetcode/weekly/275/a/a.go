package main

// 高效写法

// github.com/EndlessCheng/codeforces-go
func checkValid(matrix [][]int) bool {
	cnt := [101]int{}
	for i, row := range matrix {
		for _, v := range row {
			if cnt[v] != i {
				return false
			}
			cnt[v]++
		}
	}
	for j := range matrix[0] {
		for _, row := range matrix {
			v := row[j]
			if cnt[v] != len(matrix)+j {
				return false
			}
			cnt[v]++
		}
	}
	return true
}
