[视频讲解](https://www.bilibili.com/video/BV11d4y1i7Gs) 已出炉，**包括记忆化搜索和文末思考题的讲解**，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

套路：把路径和模 $k$ 的结果当成一个扩展维度。

具体地，定义 $f[i][j][v]$ 表示从左上走到 $(i,j)$，且路径和模 $k$ 的结果为 $v$ 时的路径数。

初始值 $f[0][0][\textit{grid}[0][0]\bmod k] = 1$，答案为 $f[m-1][n-1][0]$。

考虑从左边和上边转移过来，则有

$$
f[i][j][(v+\textit{grid}[i][j])\bmod k] += f[i][j-1][v] + f[i-1][j][v]
$$

代码实现时，为了避免判断是否越界，可以把下标都加一。此时可以设初始值 $f[0][1][0] = 1$（或者 $f[1][0][0] = 1$）简化一点点代码。

```py [sol1-Python3]
class Solution:
    def numberOfPaths(self, grid: List[List[int]], k: int) -> int:
        MOD = 10 ** 9 + 7
        m, n = len(grid), len(grid[0])
        f = [[[0] * k for _ in range(n + 1)] for _ in range(m + 1)]
        f[0][1][0] = 1
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                for v in range(k):
                    f[i + 1][j + 1][(v + x) % k] = (f[i + 1][j + 1][(v + x) % k] + f[i + 1][j][v] + f[i][j + 1][v]) % MOD
        return f[m][n][0]
```

```java [sol1-Java]
class Solution {
    public int numberOfPaths(int[][] grid, int k) {
        final var mod = (int) 1e9 + 7;
        int m = grid.length, n = grid[0].length;
        var f = new int[m + 1][n + 1][k];
        f[0][1][0] = 1;
        for (var i = 0; i < m; ++i)
            for (var j = 0; j < n; ++j)
                for (var v = 0; v < k; ++v)
                    f[i + 1][j + 1][(v + grid[i][j]) % k] = ((f[i + 1][j + 1][(v + grid[i][j]) % k] + f[i + 1][j][v]) % mod + f[i][j + 1][v]) % mod;
        return f[m][n][0];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int numberOfPaths(vector<vector<int>> &grid, int k) {
        const int mod = 1e9 + 7;
        int m = grid.size(), n = grid[0].size(), f[m + 1][n + 1][k];
        memset(f, 0, sizeof(f));
        f[0][1][0] = 1;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                for (int v = 0; v < k; ++v)
                    f[i + 1][j + 1][(v + grid[i][j]) % k] = ((f[i + 1][j + 1][(v + grid[i][j]) % k] + f[i + 1][j][v]) % mod + f[i][j + 1][v]) % mod;
        return f[m][n][0];
    }
};
```

```go [sol1-Go]
func numberOfPaths(grid [][]int, k int) int {
	const mod int = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k)
		}
	}
	f[0][1][0] = 1
	for i, row := range grid {
		for j, x := range row {
			for v := 0; v < k; v++ {
				f[i+1][j+1][(v+x)%k] = (f[i+1][j+1][(v+x)%k] + f[i+1][j][v] + f[i][j+1][v]) % mod
			}
		}
	}
	return f[m][n][0]
}
```

#### 复杂度分析

- 时间复杂度：$O(mnk)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(mnk)$。

#### 思考题

如果 $k=10^{18}$，而 $n=m=20$，要怎么做？

折半枚举。具体见视频讲解，题目见 [Codeforces 1006F. Xor-Paths](https://codeforces.com/problemset/problem/1006/F)。

力扣上的折半枚举题目：

- [805. 数组的均值分割](https://leetcode-cn.com/problems/split-array-with-same-average/)
- [2035. 将数组分成两个数组并最小化数组和的差](https://leetcode.cn/problems/partition-array-into-two-arrays-to-minimize-sum-difference/)
