请看 [视频讲解](https://www.bilibili.com/video/BV1aC4y1G7dB/) 第四题。

**核心思想**：把矩阵拉成一维的，我们需要算出每个数左边所有数的乘积，以及右边所有数的乘积，这都可以用递推得到。

先算出从 $\textit{grid}[i][j]$ 的下一个元素开始，到最后一个元素 $\textit{grid}[n-1][m-1]$ 的乘积，记作 $\textit{suf}[i][j]$。这可以从最后一行最后一列开始，倒着遍历得到。

然后算出从第一个元素 $\textit{grid}[0][0]$ 开始，到 $\textit{grid}[i][j]$ 的上一个元素的乘积，记作 $\textit{pre}[i][j]$。这可以从第一行第一列开始，正着遍历得到。

那么

$$
p[i][j] = \textit{pre}[i][j]\cdot \textit{suf}[i][j]
$$

代码实现时，可以先初始化 $p[i][j]=\textit{suf}[i][j]$，然后把 $\textit{pre}[i][j]$ 乘到 $\textit{p}[i][j]$ 中，就得到了答案。这样 $\textit{pre}$ 和 $\textit{suf}$ 就可以压缩成一个变量。

关于取模的原理，见文末的「算法小课堂」。

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
                p[i][j] = p[i][j] * pre % MOD  # 然后再乘上前缀乘积
                pre = pre * x % MOD

        return p
```

```java [sol-Java]
class Solution {
    public int[][] constructProductMatrix(int[][] grid) {
        final int MOD = 12345;
        int n = grid.length, m = grid[0].length;
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
                p[i][j] = (int) (p[i][j] * pre % MOD); // 然后再乘上前缀乘积
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
    vector<vector<int>> constructProductMatrix(vector<vector<int>> &grid) {
        const int MOD = 12345;
        int n = grid.size(), m = grid[0].size();
        vector<vector<int>> p(n, vector<int>(m));

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
                p[i][j] = p[i][j] * pre % MOD; // 然后再乘上前缀乘积
                pre = pre * grid[i][j] % MOD;
            }
        }

        return p;
    }
};
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
			p[i][j] = p[i][j] * pre % mod // 然后再乘上前缀乘积
			pre = pre * x % mod
		}
	}
	return p
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 和 $m$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 练习：前后缀分解（右边的数字为难度分）

- [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) 和本题几乎一样
- [2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/) 1395
- [2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/) 1495
- [2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/) 1695
- [2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219
- [2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/) 2223
- [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/) 2432
- [2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/) 2433
- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)

## 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。
