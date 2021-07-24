package main

/* 单调栈的应用

对第 $i$ 个人来说，他能看到的人中，靠右的不能被靠左的挡住，所以这些人从左往右的高度必须是严格单调递增的。

因此可以倒序遍历 $\textit{heights}$，用单调栈来维护人的高度，将 $\textit{heights}[i]$ 压入单调栈的同时，统计栈内比他矮的人数，压栈结束时，若栈不为空，则说明第 $i$ 个人还可以再看到一个人。

代码实现时，可以预先在单调栈中压入一个身高无限大的人，从而简化判断逻辑。

时间复杂度：$O(n)$。因为每个人至多入栈出栈各一次。

*/

// github.com/EndlessCheng/codeforces-go
func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	type pair struct{ h, i int }
	stack := []pair{{2e9, n}}
	for i := n - 1; i >= 0; i-- {
		for {
			if top := stack[len(stack)-1]; top.h >= heights[i] {
				if top.i < n { // 还可以再看到一个人
					ans[i]++
				}
				break
			}
			stack = stack[:len(stack)-1]
			ans[i]++
		}
		stack = append(stack, pair{heights[i], i})
	}
	return ans
}
