请看 [视频讲解](https://www.bilibili.com/video/BV1Fc411R7xA/)。

根据题意，如果第 $j$ 列的元素值都是 $0$（忽略第 $j$ 行的元素），说明没有队伍可以击败它，返回 $j$。

```py [sol-Python3]
class Solution:
    def findChampion(self, grid: List[List[int]]) -> int:
        for j, col in enumerate(zip(*grid)):
            if 1 not in col[:j] + col[j + 1:]:  # 没有队伍可以击败 j
                return j
```

```java [sol-Java]
class Solution {
    public int findChampion(int[][] grid) {
        int n = grid.length;
        for (int j = 0; ; j++) {
            boolean ok = true;
            for (int i = 0; i < n; i++) {
                if (i != j && grid[i][j] != 0) { // 有队伍可以击败 j
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return j;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findChampion(vector<vector<int>> &grid) {
        int n = grid.size();
        for (int j = 0;; j++) {
            bool ok = true;
            for (int i = 0; i < n; i++) {
                if (i != j && grid[i][j]) { // 有队伍可以击败 j
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return j;
            }
        }
    }
};
```

```go [sol-Go]
func findChampion(grid [][]int) int {
next:
	for j := range grid[0] {
		for i, row := range grid {
			if i != j && row[j] > 0 { // 有队伍可以击败 j
				continue next
			}
		}
		return j
	}
	panic(-1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 可以把切片改成手动枚举。
