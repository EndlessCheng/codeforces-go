## 前置知识

1. [【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/)
2. [【图解】一张图秒懂二维前缀和](https://leetcode.cn/problems/range-sum-query-2d-immutable/solution/tu-jie-yi-zhang-tu-miao-dong-er-wei-qian-84qp/)

## 思路

用**二维差分** $\mathcal{O}(1)$ 处理每个 $\textit{queries}[i]$。

然后计算二维差分矩阵的**二维前缀和**，即为答案。

代码实现时，为方便计算二维前缀和，可以在二维差分矩阵最上面添加一排 $0$，最左边添加一列 $0$，这样计算二维前缀和无需考虑下标越界。

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

        # 原地计算 diff 的二维前缀和，然后填入答案
        ans = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(n):
                diff[i + 1][j + 1] += diff[i + 1][j] + diff[i][j + 1] - diff[i][j]
                ans[i][j] = diff[i + 1][j + 1]
        return ans
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

        // 原地计算 diff 的二维前缀和，然后填入答案
        int[][] ans = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                diff[i + 1][j + 1] += diff[i + 1][j] + diff[i][j + 1] - diff[i][j];
                ans[i][j] = diff[i + 1][j + 1];
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
        vector diff(n + 2, vector<int>(n + 2));
        for (auto& q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            diff[r1 + 1][c1 + 1]++;
            diff[r1 + 1][c2 + 2]--;
            diff[r2 + 2][c1 + 1]--;
            diff[r2 + 2][c2 + 2]++;
        }

        // 原地计算 diff 的二维前缀和，然后填入答案
        vector ans(n, vector<int>(n));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                diff[i + 1][j + 1] += diff[i + 1][j] + diff[i][j + 1] - diff[i][j];
                ans[i][j] = diff[i + 1][j + 1];
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int** rangeAddQueries(int n, int** queries, int queriesSize, int* queriesColSize, int* returnSize, int** returnColumnSizes) {
    // 二维差分
    int** diff = calloc(n + 2, sizeof(int*));
    for (int i = 0; i < n + 2; i++) {
        diff[i] = calloc(n + 2, sizeof(int));
    }
    for (int i = 0; i < queriesSize; i++) {
        int r1 = queries[i][0], c1 = queries[i][1], r2 = queries[i][2], c2 = queries[i][3];
        diff[r1 + 1][c1 + 1]++;
        diff[r1 + 1][c2 + 2]--;
        diff[r2 + 2][c1 + 1]--;
        diff[r2 + 2][c2 + 2]++;
    }

    // 原地计算 diff 的二维前缀和，然后填入答案
    int** ans = malloc(n * sizeof(int*));
    *returnSize = n;
    *returnColumnSizes = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        ans[i] = malloc(n * sizeof(int));
        (*returnColumnSizes)[i] = n;
    }
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            diff[i + 1][j + 1] += diff[i + 1][j] + diff[i][j + 1] - diff[i][j];
            ans[i][j] = diff[i + 1][j + 1];
        }
    }

    for (int i = 0; i < n + 2; i++) {
        free(diff[i]);
    }
    free(diff);
    return ans;
}
```

```go [sol-Go]
func rangeAddQueries(n int, queries [][]int) [][]int {
	// 二维差分
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff[r1+1][c1+1]++
		diff[r1+1][c2+2]--
		diff[r2+2][c1+1]--
		diff[r2+2][c2+2]++
	}

	// 原地计算 diff 的二维前缀和，然后填入答案
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
			ans[i][j] = diff[i+1][j+1]
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var rangeAddQueries = function(n, queries) {
    // 二维差分
    const diff = Array.from({ length: n + 2 }, () => Array(n + 2).fill(0));
    for (const [r1, c1, r2, c2] of queries) {
        diff[r1 + 1][c1 + 1]++;
        diff[r1 + 1][c2 + 2]--;
        diff[r2 + 2][c1 + 1]--;
        diff[r2 + 2][c2 + 2]++;
    }

    // 原地计算 diff 的二维前缀和，然后填入答案
    const ans = Array.from({ length: n }, () => Array(n).fill(0));
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < n; j++) {
            diff[i + 1][j + 1] += diff[i + 1][j] + diff[i][j + 1] - diff[i][j];
            ans[i][j] = diff[i + 1][j + 1];
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn range_add_queries(n: i32, queries: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let n = n as usize;
        // 二维差分
        let mut diff = vec![vec![0; n + 2]; n + 2];
        for q in queries {
            let (r1, c1, r2, c2) = (q[0] as usize, q[1] as usize, q[2] as usize, q[3] as usize);
            diff[r1 + 1][c1 + 1] += 1;
            diff[r1 + 1][c2 + 2] -= 1;
            diff[r2 + 2][c1 + 1] -= 1;
            diff[r2 + 2][c2 + 2] += 1;
        }

        // 原地计算 diff 的二维前缀和，然后填入答案
        let mut ans = vec![vec![0; n]; n];
        for i in 0..n {
            for j in 0..n {
                diff[i + 1][j + 1] += diff[i + 1][j] + diff[i][j + 1] - diff[i][j];
                ans[i][j] = diff[i + 1][j + 1];
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+q)$，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

**注**：也可以创建 $n\times n$ 大小的 $\textit{diff}$，原地计算二维前缀和，最后直接返回 $\textit{diff}$。

## 专题训练

见数据结构题单的「**§2.2 二维差分**」和「**§1.6 二维前缀和**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
