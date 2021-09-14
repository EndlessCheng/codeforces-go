package main

// github.com/EndlessCheng/codeforces-go
func findSolution(f func(int, int) int, z int) (ans [][]int) {
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			if f(i, j) == z {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
