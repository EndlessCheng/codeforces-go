package main

import "sort"

/* 枚举 higher[0] + 双指针

将 $\textit{nums}$ 排序后，$\textit{lower}[0]$ 必然是 $\textit{nums}[0]$。我们可以在 $\textit{nums}$ 中枚举 $\textit{higher}[0]$ 的值，从而得到 $k=\dfrac{\textit{higher}[0]-\textit{lower}[0]}{2}$。

由于 $\textit{higher}[i]-\textit{lower}[i]=2k$ 是个定值，我们可以用双指针去遍历 $\textit{nums}$，计算其余的 $\textit{lower}$ 和 $\textit{higher}$ 的元素值。

细节见代码注释。

*/

// github.com/EndlessCheng/codeforces-go
func recoverArray(nums []int) []int {
	sort.Ints(nums)
	for i, n := 1, len(nums); ; i++ {
		d := nums[i] - nums[0]
		if d == 0 || d&1 > 0 { continue } // 必须保证 k 是正整数
		k := d / 2
		vis := make([]bool, n) // 用来标记出现在 higher 中的数
		vis[i] = true
		ans := []int{(nums[0] + nums[i]) / 2}
		for lo, hi := 0, i+1; hi < n; hi++ { // 双指针：lo 指向 lower，hi 指向 higher
			for lo++; lo < n && vis[lo]; lo++ {} // 跳过出现在 higher 中的数
			for ; hi < n && nums[hi]-nums[lo] < 2*k; hi++ {}
			if hi == n || nums[hi]-nums[lo] > 2*k { break }
			vis[hi] = true
			ans = append(ans, (nums[lo]+nums[hi])/2)
		}
		if len(ans) == n/2 { return ans }
	}
}
