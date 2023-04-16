下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

遍历每一行，统计一整行的元素和的最大值及其下标。

由于是从上到下遍历行的，所以找到的最大值的下标一定是相同最大值中最小的。

```py [sol1-Python3]
class Solution:
    def rowAndMaximumOnes(self, mat: List[List[int]]) -> List[int]:
        max_sum, idx = -1, 0
        for i, row in enumerate(mat):
            s = sum(row)
            if s > max_sum:
                max_sum, idx = s, i
        return [idx, max_sum]
```

```go [sol1-Go]
func rowAndMaximumOnes(mat [][]int) []int {
	maxSum, idx := -1, 0
	for i, row := range mat {
		sum := 0
		for _, x := range row {
			sum += x
		}
		if sum > maxSum {
			maxSum, idx = sum, i
		}
	}
	return []int{idx, maxSum}
}
```

### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
