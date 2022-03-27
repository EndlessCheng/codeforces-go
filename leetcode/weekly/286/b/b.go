package main

// github.com/EndlessCheng/codeforces-go
func minDeletion(a []int) (ans int) {
	odd := false // 栈大小的奇偶性
	for i, n := 0, len(a); i < n; {
		start := i
		// 注意这里的 i 和外层循环的 i 是同一个变量，因此时间复杂度为 O(n)
		for i < n && a[i] == a[start] { i++ }
		l := i - start // 连续相同元素个数
		if !odd { // 只能放一个元素
			ans += l - 1
			odd = true
		} else if l == 1 {
			odd = false
		} else { // 至多放两个元素
			ans += l - 2
		}
	}
	if odd { // 栈大小必须为偶数
		ans++
	}
	return
}
