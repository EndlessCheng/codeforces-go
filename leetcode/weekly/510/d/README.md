本题属于**相邻相关型 DP**，代表题目为 [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)。本题相当于选一个「最长一致子序列」。

类似 300 题，定义 $f[i]$ 表示在 $[0,i]$ 列中能保留的最大列数。

用「**枚举选哪个**」思考：

- 枚举上一个保留的列号 $j\ (0\le j\le i-1)$，问题变成在 $[0,j]$ 列中能保留的最大列数，即 $f[j]$。
- 如果列 $i$ 和列 $j$ 是一致的，那么用 $f[j]+1$ 更新 $f[i]$ 的最大值。

初始值：$f[i]=1$。可以只保留一列。

答案：$\max(f)$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxConsistentColumns(self, grid: list[list[int]], limit: int) -> int:
        n = len(grid[0])
        f = [1] * n
        for i in range(n):
            for j in range(i):  # 枚举上一个保留的列
                if all(abs(row[i] - row[j]) <= limit for row in grid):
                    f[i] = max(f[i], f[j] + 1)
        return max(f)
```

```java [sol-Java]
class Solution {
    public int maxConsistentColumns(int[][] grid, int limit) {
        int n = grid[0].length;
        int[] f = new int[n];
        int ans = 0;

        for (int i = 0; i < n; i++) {
            f[i] = 1;
            next:
            for (int j = 0; j < i; j++) { // 枚举上一个保留的列
                for (int[] row : grid) {
                    if (Math.abs(row[i] - row[j]) > limit) {
                        continue next; // 列 i 和列 j 不是一致的，枚举下一个 j
                    }
                }
                f[i] = Math.max(f[i], f[j] + 1);
            }
            ans = Math.max(ans, f[i]);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxConsistentColumns(vector<vector<int>>& grid, int limit) {
        int n = grid[0].size();
        vector<int> f(n, 1);
        int ans = 0;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) { // 枚举上一个保留的列
                bool ok = true;
                for (const auto& row : grid) {
                    if (abs(row[i] - row[j]) > limit) {
                        ok = false; // 列 i 和列 j 不是一致的
                        break;
                    }
                }
                if (ok) {
                    f[i] = max(f[i], f[j] + 1);
                }
            }
            ans = max(ans, f[i]);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxConsistentColumns(grid [][]int, limit int) (ans int) {
	n := len(grid[0])
	f := make([]int, n)
	for i := range n {
		f[i] = 1 // 只保留列 i
	next:
		for j := range i { // 枚举上一个保留的列
			for _, row := range grid {
				if abs(row[i]-row[j]) > limit {
					continue next // 列 i 和列 j 不是一致的，枚举下一个 j
				}
			}
			f[i] = max(f[i], f[j]+1)
		}
		ans = max(ans, f[i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 优化

最优性优化：如果 $f[j]+1 \le f[i]$，那么 $f[i]$ 不会变大，无需判断一致性。

为了配合这个优化，可以从大到小枚举 $j$，这样更可能先把 $f[i]$ 变得很大，更可能触发最优性优化。

```py [sol-Python3]
class Solution:
    def maxConsistentColumns(self, grid: list[list[int]], limit: int) -> int:
        n = len(grid[0])
        f = [0] * n
        for i in range(n):
            mx = 0
            for j in range(i - 1, -1, -1):  # 枚举上一个保留的列
                if f[j] > mx and all(abs(row[i] - row[j]) <= limit for row in grid):
                    mx = f[j]
            f[i] = mx + 1
        return max(f)
```

```java [sol-Java]
class Solution {
    public int maxConsistentColumns(int[][] grid, int limit) {
        int n = grid[0].length;
        int[] f = new int[n];
        int ans = 0;

        for (int i = 0; i < n; i++) {
            next:
            for (int j = i - 1; j >= 0; j--) { // 枚举上一个保留的列
                if (f[j] <= f[i]) {
                    continue;
                }
                for (int[] row : grid) {
                    if (Math.abs(row[i] - row[j]) > limit) {
                        continue next; // 列 i 和列 j 不是一致的，枚举下一个 j
                    }
                }
                f[i] = f[j];
            }
            f[i]++;
            ans = Math.max(ans, f[i]);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxConsistentColumns(vector<vector<int>>& grid, int limit) {
        int n = grid[0].size();
        vector<int> f(n);
        int ans = 0;

        for (int i = 0; i < n; i++) {
            for (int j = i - 1; j >= 0; j--) { // 枚举上一个保留的列
                if (f[j] <= f[i]) {
                    continue;
                }
                bool ok = true;
                for (const auto& row : grid) {
                    if (abs(row[i] - row[j]) > limit) {
                        ok = false; // 列 i 和列 j 不是一致的
                        break;
                    }
                }
                if (ok) {
                    f[i] = f[j];
                }
            }
            f[i]++;
            ans = max(ans, f[i]);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxConsistentColumns(grid [][]int, limit int) (ans int) {
	n := len(grid[0])
	f := make([]int, n)
	for i := range n {
	next:
		for j := i - 1; j >= 0; j-- { // 枚举上一个保留的列
			if f[j] <= f[i] {
				continue
			}
			for _, row := range grid {
				if abs(row[i]-row[j]) > limit {
					continue next // 列 i 和列 j 不是一致的
				}
			}
			f[i] = f[j]
		}
		f[i]++
		ans = max(ans, f[i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn^2)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§4.2 最长递增子序列**」和「**§7.1 一维 DP**」。

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
