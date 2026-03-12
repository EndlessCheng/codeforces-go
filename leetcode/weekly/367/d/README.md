请先完成本题的**一维版本**：[238. 除了自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)。

把矩阵平铺成一维数组，就是 238 题了。我们需要算出每个数左边所有数的乘积，以及右边所有数的乘积。

先算出从 $\textit{grid}[i][j]$ 的下一个元素开始，到最后一个元素 $\textit{grid}[n-1][m-1]$ 的乘积，记作 $\textit{suf}[i][j]$。这可以从最后一个数 $\textit{grid}[n-1][m-1]$ 开始，倒着遍历 $\textit{grid}$ 得到。

然后算出从第一个数 $\textit{grid}[0][0]$ 开始，到 $\textit{grid}[i][j]$ 的上一个元素的乘积，记作 $\textit{pre}[i][j]$。这可以从第一行第一列开始，正着遍历得到。

那么

$$
p[i][j] = \textit{pre}[i][j]\cdot \textit{suf}[i][j]
$$

代码实现时，可以先初始化 $p[i][j]=\textit{suf}[i][j]$，然后在计算 $\textit{pre}[i][j]$ 的过程中，把 $\textit{pre}[i][j]$ 乘到 $\textit{p}[i][j]$ 中，就得到了最终答案。这样写的话，$\textit{pre}$ 和 $\textit{suf}$ 可以直接用单个变量表示，无需创建数组。

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def constructProductMatrix(self, grid: List[List[int]]) -> List[List[int]]:
        MOD = 12345
        n, m = len(grid), len(grid[0])
        p = [[0] * m for _ in range(n)]

        suf = 1  # 后缀乘积
        for i in range(n - 1, -1, -1):
            for j in range(m - 1, -1, -1):
                p[i][j] = suf  # p[i][j] 先初始化成后缀乘积
                suf = suf * grid[i][j] % MOD

        pre = 1  # 前缀乘积
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                p[i][j] = p[i][j] * pre % MOD  # 乘上前缀乘积
                pre = pre * x % MOD

        return p
```

```java [sol-Java]
class Solution {
    public int[][] constructProductMatrix(int[][] grid) {
        final int MOD = 12345;
        int n = grid.length;
        int m = grid[0].length;
        int[][] p = new int[n][m];

        long suf = 1; // 后缀乘积
        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                p[i][j] = (int) suf; // p[i][j] 先初始化成后缀乘积
                suf = suf * grid[i][j] % MOD;
            }
        }

        long pre = 1; // 前缀乘积
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                p[i][j] = (int) (p[i][j] * pre % MOD); // 乘上前缀乘积
                pre = pre * grid[i][j] % MOD;
            }
        }

        return p;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> constructProductMatrix(vector<vector<int>>& grid) {
        constexpr int MOD = 12345;
        int n = grid.size(), m = grid[0].size();
        vector p(n, vector<int>(m));

        long long suf = 1; // 后缀乘积
        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                p[i][j] = suf; // p[i][j] 先初始化成后缀乘积
                suf = suf * grid[i][j] % MOD;
            }
        }

        long long pre = 1; // 前缀乘积
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                p[i][j] = p[i][j] * pre % MOD; // 乘上前缀乘积
                pre = pre * grid[i][j] % MOD;
            }
        }

        return p;
    }
};
```

```c [sol-C]
int** constructProductMatrix(int** grid, int gridSize, int* gridColSize, int* returnSize, int** returnColumnSizes) {
    const int MOD = 12345;
    int n = gridSize, m = gridColSize[0];
    int** p = malloc(n * sizeof(int*));
    *returnSize = n;
    *returnColumnSizes = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        p[i] = malloc(m * sizeof(int));
        (*returnColumnSizes)[i] = m;
    }

    long long suf = 1; // 后缀乘积
    for (int i = n - 1; i >= 0; i--) {
        for (int j = m - 1; j >= 0; j--) {
            p[i][j] = suf; // p[i][j] 先初始化成后缀乘积
            suf = suf * grid[i][j] % MOD;
        }
    }

    long long pre = 1; // 前缀乘积
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            p[i][j] = p[i][j] * pre % MOD; // 乘上前缀乘积
            pre = pre * grid[i][j] % MOD;
        }
    }

    return p;
}
```

```go [sol-Go]
func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	n, m := len(grid), len(grid[0])
	p := make([][]int, n)
	suf := 1 // 后缀乘积
	for i := n - 1; i >= 0; i-- {
		p[i] = make([]int, m)
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suf // p[i][j] 先初始化成后缀乘积
			suf = suf * grid[i][j] % mod
		}
	}

	pre := 1 // 前缀乘积
	for i, row := range grid {
		for j, x := range row {
			p[i][j] = p[i][j] * pre % mod // 乘上前缀乘积
			pre = pre * x % mod
		}
	}
	return p
}
```

```js [sol-JavaScript]
var constructProductMatrix = function(grid) {
    const MOD = 12345;
    const n = grid.length, m = grid[0].length;
    const p = Array.from({ length: n }, () => Array(m).fill(0));

    let suf = 1; // 后缀乘积
    for (let i = n - 1; i >= 0; i--) {
        for (let j = m - 1; j >= 0; j--) {
            p[i][j] = suf; // p[i][j] 先初始化成后缀乘积
            suf = suf * grid[i][j] % MOD;
        }
    }

    let pre = 1; // 前缀乘积
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            p[i][j] = p[i][j] * pre % MOD; // 乘上前缀乘积
            pre = pre * grid[i][j] % MOD;
        }
    }

    return p;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn construct_product_matrix(grid: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        const MOD: i64 = 12345;
        let n = grid.len();
        let m = grid[0].len();
        let mut p = vec![vec![0; m]; n];

        let mut suf = 1; // 后缀乘积
        for i in (0..n).rev() {
            for j in (0..m).rev() {
                p[i][j] = suf as i32; // p[i][j] 先初始化成后缀乘积
                suf = suf * grid[i][j] as i64 % MOD;
            }
        }

        let mut pre = 1; // 前缀乘积
        for i in 0..n {
            for j in 0..m {
                p[i][j] = (p[i][j] as i64 * pre % MOD) as i32; // 乘上前缀乘积
                pre = pre * grid[i][j] as i64 % MOD;
            }
        }

        p
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 和 $m$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
