package main

// github.com/EndlessCheng/codeforces-go
func maxValueOfCoins(piles [][]int, k int) int {
	f := make([]int, k+1)
	sumN := 0
	for _, pile := range piles {
		n := len(pile)
		for i := 1; i < n; i++ {
			pile[i] += pile[i-1] // pile 前缀和
		}
		sumN = min(sumN+n, k) // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
		for j := sumN; j > 0; j-- {
			for w, v := range pile[:min(n, j)] {
				f[j] = max(f[j], f[j-w-1]+v)
			}
		}
	}
	return f[k]
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
