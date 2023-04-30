package main

// https://space.bilibili.com/206214
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	type pair struct{ r, c int }
	pos := make([]pair, m*n+1)
	for i, row := range mat {
		for j, x := range row {
			pos[x] = pair{i, j}
		}
	}

	rowCnt := make([]int, m)
	colCnt := make([]int, n)
	for i, x := range arr {
		p := pos[x]
		rowCnt[p.r]++
		colCnt[p.c]++
		if rowCnt[p.r] == n || colCnt[p.c] == m {
			return i
		}
	}
	return -1
}
