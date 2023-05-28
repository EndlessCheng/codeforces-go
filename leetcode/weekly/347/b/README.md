## 视频讲解

见[【周赛 347】](https://www.bilibili.com/video/BV1fo4y1T7MQ/) 第二题，欢迎点赞投币！

## 算法一：模拟

枚举每个位置，往左上和右下遍历，用哈希表统计不同元素个数。

```py [sol-Python3]
class Solution:
    def differenceOfDistinctValues(self, grid: List[List[int]]) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                # topLeft
                s = set()
                x, y = i - 1, j - 1
                while x >= 0 and y >= 0:
                    s.add(grid[x][y])
                    x -= 1
                    y -= 1
                sz = len(s)

                # bottomRight
                s.clear()
                x, y = i + 1, j + 1
                while x < m and y < n:
                    s.add(grid[x][y])
                    x += 1
                    y += 1
                ans[i][j] = abs(sz - len(s))
        return ans
```

```go [sol-Go]
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			// topLeft
			set := map[int]struct{}{}
			for x, y := i-1, j-1; x >= 0 && y >= 0; {
				set[grid[x][y]] = struct{}{}
				x--
				y--
			}
			sz := len(set)
			
			// bottomRight
			set = map[int]struct{}{}
			for x, y := i+1, j+1; x < m && y < n; {
				set[grid[x][y]] = struct{}{}
				x++
				y++
			}
			ans[i][j] = abs(sz - len(set))
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\cdot \min(m,n))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(\min(m,n))$。返回值不计入。

## 算法二：前后缀分解

我们可以从第一行和第一列的每个位置出发，向右下遍历，一边遍历一边用哈希表统计不同元素个数，遍历到 $\textit{grid}[i][j]$ 时，哈希表的大小就是 $\textit{topLeft}[i+1][j+1]$。对于 $\textit{bottomRight}$ 也是同理，从最后一行和最后一列出发去遍历。

代码实现时，可以直接把 $\textit{topLeft}$ 和 $\textit{bottomRight}$ 保存到 $\textit{ans}$ 中。

```py [sol-Python3]
class Solution:
    def differenceOfDistinctValues(self, grid: List[List[int]]) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]
        for k in range(1, m + n):
            min_j = max(n - k, 0)
            max_j = min(m + n - 1 - k, n - 1)
            # topLeft
            s = set()
            for j in range(min_j, max_j):
                i = k + j - n
                s.add(grid[i][j])
                ans[i + 1][j + 1] = len(s)
            # bottomRight
            s.clear()
            for j in range(max_j, min_j, -1):
                i = k + j - n
                s.add(grid[i][j])
                ans[i - 1][j - 1] = abs(ans[i - 1][j - 1] - len(s))
        return ans
```

```go [sol-Go]
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for s := 1; s < m+n; s++ {
		minJ := max(0, n-s)
		maxJ := min(n-1, n-s+m-1)
		// topLeft
		set := map[int]struct{}{}
		for j := minJ; j < maxJ; j++ {
			i := s + j - n
			set[grid[i][j]] = struct{}{}
			ans[i+1][j+1] = len(set)
		}
		// bottomRight
		set = map[int]struct{}{}
		for j := maxJ; j > minJ; j-- {
			i := s + j - n
			set[grid[i][j]] = struct{}{}
			ans[i-1][j-1] = abs(ans[i-1][j-1] - len(set))
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个单元格至多访问两次。
- 空间复杂度：$\mathcal{O}(\min(m,n))$。返回值不计入。
