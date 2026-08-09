两个不重叠的正方形，要么一左一右，要么一上一下。

比如一上一下的情况，枚举水平分割线切开了 $i$ 行和 $i+1$ 行，问题变成：

- 计算 $\textit{mat}$ 的 $[0,i]$ 行中的 [221. 最大正方形](https://leetcode.cn/problems/maximal-square/)。
- 计算 $\textit{mat}$ 的 $[i+1,m-1]$ 行中的 [221. 最大正方形](https://leetcode.cn/problems/maximal-square/)。

取两个正方形的面积的较小值，即为一上一下情况中的单个正方形的最大面积。

> **注**：如果存在边长为 $k$ 的全 $\texttt{1}$ 正方形，那么也存在边长小于 $k$ 的全 $\texttt{1}$ 正方形。

一左一右的做法类似，枚举垂直分割线的位置。代码实现时，可以直接把 $\textit{mat}$ 转置（或者旋转 $90^\circ$），从而复用同一段代码。

对于 221 题，可以用动态规划，或者单调栈，或者二分答案+二维前缀和等。下面用的动态规划写法（空间优化），来自 [我的题解](https://leetcode.cn/problems/maximal-square/solutions/3704858/he-85-ti-yi-yang-de-zuo-fa-pythonjavaccg-az54/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def calc(self, mat: list[list[int]]) -> int:
        m, n = len(mat), len(mat[0])

        # 221. 最大正方形（空间优化写法）
        # 计算 mat 下半部分的最大正方形的边长
        suf_max = [0] * (m + 1)
        f = [0] * (n + 1)
        for i in range(m - 1, 0, -1):
            last = 0
            for j, x in enumerate(mat[i]):
                if x:
                    tmp = f[j + 1]
                    f[j + 1] = min(last, f[j + 1], f[j]) + 1
                    last = tmp
                else:
                    f[j + 1] = 0
                    last = 0
            suf_max[i] = max(suf_max[i + 1], max(f))

        # 计算 mat 上半部分的最大正方形的边长
        ans = pre_max = 0
        f = [0] * (n + 1)
        for i, row in enumerate(mat):
            last = 0
            for j, x in enumerate(row):
                if x:
                    tmp = f[j + 1]
                    f[j + 1] = min(last, f[j + 1], f[j]) + 1
                    last = tmp
                else:
                    f[j + 1] = 0
                    last = 0
            if suf_max[i + 1] <= ans:
                break  # 最优性优化：继续循环不会让 ans 变大
            pre_max = max(pre_max, max(f))
            ans = max(ans, min(pre_max, suf_max[i + 1]))  # 题目要求两个正方形的边长相等

        return ans * ans

    def maxArea(self, mat: list[list[int]]) -> int:
        return max(self.calc(mat), self.calc(list(zip(*mat))))
```

```java [sol-Java]
class Solution {
    public int maxArea(int[][] mat) {
        return Math.max(calc(mat), calc(transpose(mat)));
    }

    private int calc(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;

        // 221. 最大正方形（空间优化写法）
        // 计算 mat 下半部分的最大正方形的边长
        int[] sufMax = new int[m];
        int[] f = new int[n + 1];
        int mx = 0;
        for (int i = m - 1; i > 0; i--) {
            int last = 0;
            for (int j = 0; j < n; j++) {
                int x = mat[i][j];
                if (x == 1) {
                    int tmp = f[j + 1];
                    f[j + 1] = Math.min(Math.min(last, f[j + 1]), f[j]) + 1;
                    last = tmp;
                    mx = Math.max(mx, f[j + 1]);
                } else {
                    f[j + 1] = 0;
                    last = 0;
                }
            }
            sufMax[i] = mx;
        }

        int ans = 0;
        // 计算 mat 上半部分的最大正方形的边长
        int preMax = 0;
        Arrays.fill(f, 0);
        for (int i = 0; i < m - 1; i++) {
            int last = 0;
            for (int j = 0; j < n; j++) {
                int x = mat[i][j];
                if (x == 1) {
                    int tmp = f[j + 1];
                    f[j + 1] = Math.min(Math.min(last, f[j + 1]), f[j]) + 1;
                    last = tmp;
                    preMax = Math.max(preMax, f[j + 1]);
                } else {
                    f[j + 1] = 0;
                    last = 0;
                }
            }
            if (sufMax[i + 1] <= ans) {
                break; // 最优性优化：继续循环不会让 ans 变大
            }
            ans = Math.max(ans, Math.min(preMax, sufMax[i + 1])); // 题目要求两个正方形的边长相等
        }

        return ans * ans;
    }

    // 转置矩阵 mat
    private int[][] transpose(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;
        int[][] a = new int[n][m];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                a[i][j] = mat[j][i];
            }
        }
        return a;
    }
}
```

```cpp [sol-C++]
class Solution {
    int calc(const vector<vector<int>>& mat) {
        int m = mat.size(), n = mat[0].size();

        // 221. 最大正方形（空间优化写法）
        // 计算 mat 下半部分的最大正方形的边长
        vector<int> suf_max(m);
        vector<int> f(n + 1);
        int mx = 0;
        for (int i = m - 1; i > 0; i--) {
            int last = 0;
            for (int j = 0; j < n; j++) {
                if (mat[i][j]) {
                    int tmp = f[j + 1];
                    f[j + 1] = min(min(last, f[j + 1]), f[j]) + 1;
                    last = tmp;
                    mx = max(mx, f[j + 1]);
                } else {
                    f[j + 1] = 0;
                    last = 0;
                }
            }
            suf_max[i] = mx;
        }

        int ans = 0;
        // 计算 mat 上半部分的最大正方形的边长
        int pre_max = 0;
        ranges::fill(f, 0);
        for (int i = 0; i < m - 1; i++) {
            int last = 0;
            for (int j = 0; j < n; j++) {
                if (mat[i][j]) {
                    int tmp = f[j + 1];
                    f[j + 1] = min(min(last, f[j + 1]), f[j]) + 1;
                    last = tmp;
                    pre_max = max(pre_max, f[j + 1]);
                } else {
                    f[j + 1] = 0;
                    last = 0;
                }
            }
            if (suf_max[i + 1] <= ans) {
                break; // 最优性优化：继续循环不会让 ans 变大
            }
            ans = max(ans, min(pre_max, suf_max[i + 1])); // 题目要求两个正方形的边长相等
        }

        return ans * ans;
    }

    // 转置矩阵 mat
    vector<vector<int>> transpose(vector<vector<int>>& mat) {
        int m = mat.size(), n = mat[0].size();
        vector a(n, vector<int>(m));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < mat.size(); j++) {
                a[i][j] = mat[j][i];
            }
        }
        return a;
    }

public:
    int maxArea(vector<vector<int>>& mat) {
        return max(calc(mat), calc(transpose(mat)));
    }
};
```

```go [sol-Go]
func calc(mat [][]int) int {
	m, n := len(mat), len(mat[0])

	// 221. 最大正方形（空间优化写法）
	// 计算 mat 下半部分的最大正方形的边长
	sufMax := make([]int, m)
	f := make([]int, n+1)
	mx := 0
	for i := m - 1; i > 0; i-- {
		last := 0
		for j, x := range mat[i] {
			if x == 1 {
				tmp := f[j+1]
				f[j+1] = min(last, f[j+1], f[j]) + 1
				last = tmp
				mx = max(mx, f[j+1])
			} else {
				f[j+1] = 0
				last = 0
			}
		}
		sufMax[i] = mx
	}

	ans := 0
	// 计算 mat 上半部分的最大正方形的边长
	preMax := 0
	clear(f)
	for i, row := range mat[:m-1] {
		last := 0
		for j, x := range row {
			if x == 1 {
				tmp := f[j+1]
				f[j+1] = min(last, f[j+1], f[j]) + 1
				last = tmp
				preMax = max(preMax, f[j+1])
			} else {
				f[j+1] = 0
				last = 0
			}
		}
		if sufMax[i+1] <= ans {
			break // 最优性优化：继续循环不会让 ans 变大
		}
		ans = max(ans, min(preMax, sufMax[i+1])) // 题目要求两个正方形的边长相等
	}

	return ans * ans
}

// 转置矩阵 mat
func transpose(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j, row := range mat {
			a[i][j] = row[i]
		}
	}
	return a
}

func maxArea(mat [][]int) int {
	return max(calc(mat), calc(transpose(mat)))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别是 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$ 或 $\mathcal{O}(m+n)$。如果不转置矩阵，则可以做到 $\mathcal{O}(m+n)$ 的空间复杂度。

## 相似题目

[3197. 包含所有 1 的最小矩形面积 II](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-ii/)

## 专题训练

1. 动态规划题单的「**专题：前后缀分解**」和「**§7.5 子矩形 DP**」。
2. 单调栈题单的「**二、矩形**」

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
