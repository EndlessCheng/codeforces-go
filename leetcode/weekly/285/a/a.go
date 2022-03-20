package main

// github.com/EndlessCheng/codeforces-go
func countHillValley(a []int) (ans int) {
	for i, n := 0, len(a); i < n; {
		start, v := i, a[i]
		for i < n && a[i] == v { i++ } // 注意这里的 i 和外层循环的 i 是同一个变量，因此时间复杂度为 O(n)
		if start > 0 && i < n && a[start-1] < v == (a[i] < v) {
			ans++
		}
	}
	return
}
