请看 [视频讲解](https://www.bilibili.com/video/BV1EL411C7YU/) 第二题。

按题意模拟。

代码实现时，可以把每个值的坐标记录到一个数组中，方便判断。

```py [sol1-Python3]
class Solution:
    def checkValidGrid(self, grid: List[List[int]]) -> bool:
        pos = [0] * (len(grid) ** 2)
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                pos[x] = (i, j)  # 记录坐标
        if pos[0] != (0, 0):  # 必须从左上角出发
            return False
        for (i, j), (x, y) in pairwise(pos):
            dx, dy = abs(x - i), abs(y - j)  # 移动距离
            if (dx != 2 or dy != 1) and (dx != 1 or dy != 2):  # 不合法
                return False
        return True
```

```go [sol1-Go]
func checkValidGrid(grid [][]int) bool {
	type pair struct{ i, j int }
	pos := make([]pair, len(grid)*len(grid))
	for i, row := range grid {
		for j, x := range row {
			pos[x] = pair{i, j} // 记录坐标
		}
	}
	if pos[0] != (pair{}) { // 必须从左上角出发
		return false
	}
	for i := 1; i < len(pos); i++ {
		dx := abs(pos[i].i - pos[i-1].i)
		dy := abs(pos[i].j - pos[i-1].j) // 移动距离
		if (dx != 2 || dy != 1) && (dx != 1 || dy != 2) { // 不合法
			return false
		}
	}
	return true
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{grid}$ 的长度。
- 空间复杂度：$O(n^2)$。
