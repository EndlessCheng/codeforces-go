无脑做法：二维前缀和 + 二维差分

不了解二维前缀和同学可以先看看 304 题。

二维差分可以结合二维前缀和与一维差分的思想推导出来。当我们对一个左上角在 $(x1,y1)$，右下角在 $(x2,y2)$ 矩形区域全部增加 $x$ 时，相当于在二维差分矩阵上对 $(x1,y1)$ 增加 $x$，对 $(x1,y2+1)$ 和 $(x2+1,y1)$ 减少 $x$，由于这样两个地方都减少了 $x$，我们还需要在 $(x2+1,y2+1)$ 处增加 $x$，读者可以用二维前缀和对比体会这一做法。

回到本题。由于邮票可以互相重叠，我们遵从能放就放邮票的策略，遍历所有的空位，尝试以该空位为左上角放置邮票。如果这一矩形没有出界且区域内没有被占据的格子，那么就可以放置邮票，并按照二维差分的做法将区域内的所有元素值加一。

遍历结束后，我们需要从二维差分矩阵还原出二维计数矩阵，这可以通过对二维差分矩阵求二维前缀和求出。遍历计数矩阵，如果存在一个空格子的计数值为 $0$，就表明该空格子没有被邮票覆盖，返回 $\texttt{false}$，否则返回 $\texttt{true}$。

```go [sol1-Golang]
func possibleToStamp(grid [][]int, stampHeight, stampWidth int) bool {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	diff := make([][]int, m+1)
	diff[0] = make([]int, n+1)
	for i, row := range grid {
		sum[i+1] = make([]int, n+1)
		for j, v := range row { // grid 的二维前缀和
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
		diff[i+1] = make([]int, n+1)
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				x, y := i+stampHeight, j+stampWidth // 注意这是矩形右下角横纵坐标都 +1 后的位置
				if x <= m && y <= n && sum[x][y]-sum[x][j]-sum[i][y]+sum[i][j] == 0 {
					diff[i][j]++
					diff[i][y]--
					diff[x][j]--
					diff[x][y]++ // 更新二维差分
				}
			}
		}
	}

	// 还原二维差分矩阵对应的计数矩阵，这里用滚动数组实现
	cnt := make([]int, n+1)
	pre := make([]int, n+1)
	for i, row := range grid {
		for j, v := range row {
			cnt[j+1] = cnt[j] + pre[j+1] - pre[j] + diff[i][j]
			if cnt[j+1] == 0 && v == 0 {
				return false
			}
		}
		cnt, pre = pre, cnt
	}
	return true
}
```

```C++ [sol1-C++]
class Solution {
public:
    bool possibleToStamp(vector<vector<int>> &grid, int stampHeight, int stampWidth) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> sum(m + 1, vector<int>(n + 1)), diff(m + 1, vector<int>(n + 1));
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) { // grid 的二维前缀和
                sum[i + 1][j + 1] = sum[i + 1][j] + sum[i][j + 1] - sum[i][j] + grid[i][j];
            }
        }

        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                int x = i + stampHeight, y = j + stampWidth; // 注意这是矩形右下角横纵坐标都 +1 后的位置
                if (x <= m && y <= n && sum[x][y] - sum[x][j] - sum[i][y] + sum[i][j] == 0) {
                    ++diff[i][j];
                    --diff[i][y];
                    --diff[x][j];
                    ++diff[x][y]; // 更新二维差分
                }
            }
        }

        // 还原二维差分矩阵对应的计数矩阵，这里用滚动数组实现
        vector<int> cnt(n + 1), pre(n + 1);
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                cnt[j + 1] = cnt[j] + pre[j + 1] - pre[j] + diff[i][j];
                if (cnt[j + 1] == 0 && grid[i][j] == 0) {
                    return false;
                }
            }
            swap(cnt, pre);
        }
        return true;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def possibleToStamp(self, grid: List[List[int]], stampHeight: int, stampWidth: int) -> bool:
        m, n = len(grid), len(grid[0])
        sum = [[0] * (n + 1) for _ in range(m + 1)]
        diff = [[0] * (n + 1) for _ in range(m + 1)]
        for i, row in enumerate(grid):
            for j, v in enumerate(row):  # grid 的二维前缀和
                sum[i + 1][j + 1] = sum[i + 1][j] + sum[i][j + 1] - sum[i][j] + v

        for i, row in enumerate(grid):
            for j, v in enumerate(row):
                if v == 0:
                    x, y = i + stampHeight, j + stampWidth  # 注意这是矩形右下角横纵坐标都 +1 后的位置
                    if x <= m and y <= n and sum[x][y] - sum[x][j] - sum[i][y] + sum[i][j] == 0:
                        diff[i][j] += 1
                        diff[i][y] -= 1
                        diff[x][j] -= 1
                        diff[x][y] += 1  # 更新二维差分

        # 还原二维差分矩阵对应的计数矩阵，这里用滚动数组实现
        cnt, pre = [0] * (n + 1), [0] * (n + 1)
        for i, row in enumerate(grid):
            for j, v in enumerate(row):
                cnt[j + 1] = cnt[j] + pre[j + 1] - pre[j] + diff[i][j]
                if cnt[j + 1] == 0 and v == 0:
                    return False
            cnt, pre = pre, cnt
        return True
```
