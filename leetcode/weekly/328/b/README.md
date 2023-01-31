回想下，一维我们是怎么做的。

> 把一段区间的元素 $+1$，可以记录「变化量」，把区间起点 $+1$，区间**终点右侧** $-1$，这样「变化量的前缀和」就是实际结果。注意 $-1$ 的地方是终点下标 $+1$ 的位置。

推广到二维，也就是把一个区域的元素都 $+1$，那就需要用一个二维数组来记录变化量，然后对这个二维数组求二维前缀和，就得到了实际结果，即本题的 $\textit{mat}$。

> 不了解二维前缀和的同学可以看看 [304. 二维区域和检索 - 矩阵不可变](https://leetcode.cn/problems/range-sum-query-2d-immutable/)。

怎么记录二维的变化量呢？

从二维前缀和的角度来看，对区域左上角 $+1$ 会对所有右下位置产生影响，那么在区域右上角的右边相邻处和左下角的下边相邻处 $-1$ 可以消除这个影响，但是两个 $-1$ 又会对区域右下角的右下所有位置产生影响，所以要在右下角的右下相邻处再 $+1$ 还原回来。

附：[视频讲解](https://www.bilibili.com/video/BV1QT41127kJ/)。

```py [sol1-Python3]
class Solution:
    def rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        # 二维差分模板
        diff = [[0] * (n + 2) for _ in range(n + 2)]
        for r1, c1, r2, c2 in queries:
            diff[r1 + 1][c1 + 1] += 1
            diff[r1 + 1][c2 + 2] -= 1
            diff[r2 + 2][c1 + 1] -= 1
            diff[r2 + 2][c2 + 2] += 1

        # 用二维前缀和复原（原地修改）
        for i in range(1, n + 1):
            for j in range(1, n + 1):
                diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1]
        # 保留中间 n*n 的部分，即为答案
        diff = diff[1:-1]
        for i, row in enumerate(diff):
            diff[i] = row[1:-1]
        return diff
```

```java [sol1-Java]
class Solution {
    public int[][] rangeAddQueries(int n, int[][] queries) {
        // 二维差分模板
        int[][] diff = new int[n + 2][n + 2], ans = new int[n][n];
        for (int[] q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            ++diff[r1 + 1][c1 + 1];
            --diff[r1 + 1][c2 + 2];
            --diff[r2 + 2][c1 + 1];
            ++diff[r2 + 2][c2 + 2];
        }
        // 用二维前缀和复原
        for (int i = 1; i <= n; ++i)
            for (int j = 1; j <= n; ++j)
                ans[i - 1][j - 1] = diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1];
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<vector<int>> rangeAddQueries(int n, vector<vector<int>> &queries) {
        // 二维差分模板
        vector<vector<int>> diff(n + 2, vector<int>(n + 2));
        for (auto &q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            ++diff[r1 + 1][c1 + 1];
            --diff[r1 + 1][c2 + 2];
            --diff[r2 + 2][c1 + 1];
            ++diff[r2 + 2][c2 + 2];
        }

        // 用二维前缀和复原（原地修改）
        for (int i = 1; i <= n; ++i)
            for (int j = 1; j <= n; ++j)
                diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1];
        // 保留中间 n*n 的部分，即为答案
        diff.pop_back(), diff.erase(diff.begin());
        for (auto &row : diff)
            row.pop_back(), row.erase(row.begin());
        return diff;
    }
};
```

```go [sol1-Go]
func rangeAddQueries(n int, queries [][]int) [][]int {
	// 二维差分模板
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	update := func(r1, c1, r2, c2, x int) {
		diff[r1+1][c1+1] += x
		diff[r1+1][c2+2] -= x
		diff[r2+2][c1+1] -= x
		diff[r2+2][c2+2] += x
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原（原地修改）
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
		}
	}
	// 保留中间 n*n 的部分，即为答案
	diff = diff[1 : n+1]
	for i, row := range diff {
		diff[i] = row[1 : n+1]
	}
	return diff
}
```

#### 复杂度分析

- 时间复杂度：$O(n^2+q)$，其中 $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n^2)$。

#### 相似题目

- [2132. 用邮票贴满网格图](https://leetcode.cn/problems/stamping-the-grid/)
