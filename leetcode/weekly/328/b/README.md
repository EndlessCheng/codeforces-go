## 前置知识

1. [【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/)
2. [【图解】一张图秒懂二维前缀和](https://leetcode.cn/problems/range-sum-query-2d-immutable/solution/tu-jie-yi-zhang-tu-miao-dong-er-wei-qian-84qp/)

## 思路

利用**二维差分**处理 $\textit{queries}$，然后计算二维差分矩阵的**二维前缀和**，即为答案。

附：[视频讲解](https://www.bilibili.com/video/BV1QT41127kJ/)。

```py [sol-Python3]
class Solution:
    def rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        # 二维差分
        diff = [[0] * (n + 2) for _ in range(n + 2)]
        for r1, c1, r2, c2 in queries:
            diff[r1 + 1][c1 + 1] += 1
            diff[r1 + 1][c2 + 2] -= 1
            diff[r2 + 2][c1 + 1] -= 1
            diff[r2 + 2][c2 + 2] += 1

        # 计算 diff 的二维前缀和（原地修改）
        for i in range(1, n + 1):
            for j in range(1, n + 1):
                diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1]

        # 保留中间 n*n 的部分，即为答案
        diff = diff[1:-1]
        for i, row in enumerate(diff):
            diff[i] = row[1:-1]
        return diff
```

```java [sol-Java]
class Solution {
    public int[][] rangeAddQueries(int n, int[][] queries) {
        // 二维差分
        int[][] diff = new int[n + 2][n + 2];
        for (int[] q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            diff[r1 + 1][c1 + 1]++;
            diff[r1 + 1][c2 + 2]--;
            diff[r2 + 2][c1 + 1]--;
            diff[r2 + 2][c2 + 2]++;
        }

        // 计算 diff 的二维前缀和（原地修改），然后填入 ans
        int[][] ans = new int[n][n];
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1];
                ans[i - 1][j - 1] = diff[i][j];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> rangeAddQueries(int n, vector<vector<int>>& queries) {
        // 二维差分
        vector<vector<int>> diff(n + 2, vector<int>(n + 2));
        for (auto& q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            diff[r1 + 1][c1 + 1]++;
            diff[r1 + 1][c2 + 2]--;
            diff[r2 + 2][c1 + 1]--;
            diff[r2 + 2][c2 + 2]++;
        }

        // 计算 diff 的二维前缀和（原地修改）
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1];
            }
        }

        // 保留中间 n*n 的部分，即为答案
        diff.pop_back();
        diff.erase(diff.begin());
        for (auto& row : diff) {
            row.pop_back();
            row.erase(row.begin());
        }
        return diff;
    }
};
```

```go [sol-Go]
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

	// 计算 diff 的二维前缀和（原地修改）
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

- 时间复杂度：$\mathcal{O}(n^2+q)$，其中 $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
