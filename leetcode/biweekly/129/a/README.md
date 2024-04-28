遍历矩阵中的每个 $2\times 2$ 子矩形。

对于每个子矩形，统计 B 和 W 的个数，如果其中一个字母的出现次数 $\ge 3$，则返回 $\texttt{true}$。

如果四个子矩形都不满足要求，返回 $\texttt{false}$。

```py [sol-Python3]
class Solution:
    def canMakeSquare(self, grid: List[List[str]]) -> bool:
        def check(i: int, j: int) -> bool:
            cnt = Counter()
            cnt[grid[i][j]] += 1
            cnt[grid[i][j + 1]] += 1
            cnt[grid[i + 1][j]] += 1
            cnt[grid[i + 1][j + 1]] += 1
            return cnt['B'] >= 3 or cnt['W'] >= 3
        return check(0, 0) or check(0, 1) or check(1, 0) or check(1, 1)
```

```java [sol-Java]
class Solution {
    public boolean canMakeSquare(char[][] grid) {
        return check(grid, 0, 0) || check(grid, 0, 1) || check(grid, 1, 0) || check(grid, 1, 1);
    }

    private boolean check(char[][] grid, int i, int j) {
        int[] cnt = new int[2];
        cnt[grid[i][j] & 1]++;
        cnt[grid[i][j + 1] & 1]++;
        cnt[grid[i + 1][j] & 1]++;
        cnt[grid[i + 1][j + 1] & 1]++;
        return cnt[0] >= 3 || cnt[1] >= 3;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canMakeSquare(vector<vector<char>>& grid) {
        auto check = [&](int i, int j) {
            int cnt[2]{};
            cnt[grid[i][j] & 1]++;
            cnt[grid[i][j + 1] & 1]++;
            cnt[grid[i + 1][j] & 1]++;
            cnt[grid[i + 1][j + 1] & 1]++;
            return cnt[0] >= 3 || cnt[1] >= 3;
        };
        return check(0, 0) || check(0, 1) || check(1, 0) || check(1, 1);
    }
};
```

```go [sol-Go]
func canMakeSquare(grid [][]byte) bool {
	check := func(i, j int) bool {
		cnt := [2]int{}
		cnt[grid[i][j]&1]++
		cnt[grid[i][j+1]&1]++
		cnt[grid[i+1][j]&1]++
		cnt[grid[i+1][j+1]&1]++
		return cnt[0] >= 3 || cnt[1] >= 3
	}
	return check(0, 0) || check(0, 1) || check(1, 0) || check(1, 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
