下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

遍历每一列，求出数字的字符串形式的长度的最大值。

```py [sol1-Python3]
class Solution:
    def findColumnWidth(self, grid: List[List[int]]) -> List[int]:
        return [max(len(str(x)) for x in col) for col in zip(*grid)]
```

```go [sol1-Go]
func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			ans[j] = max(ans[j], len(strconv.Itoa(row[j])))
		}
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
```

也可以手动计算长度。


```py [sol2-Python3]
class Solution:
    def findColumnWidth(self, grid: List[List[int]]) -> List[int]:
        ans = [0] * len(grid[0])
        for j, col in enumerate(zip(*grid)):
            for x in col:
                x_len = int(x <= 0)
                x = abs(x)
                while x:
                    x_len += 1
                    x //= 10
                ans[j] = max(ans[j], x_len)
        return ans
```

```go [sol2-Go]
func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			xLen := 0
			if row[j] <= 0 {
				xLen = 1
			}
			for x := row[j]; x != 0; x /= 10 {
				xLen++
			}
			ans[j] = max(ans[j], xLen)
		}
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(1)$。手动算长度则为 $O(1)$。Python 忽略 zip* 的空间。
