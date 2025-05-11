## 分析

设整个 $\textit{grid}$ 的元素和为 $\textit{total}$。

设第一部分的元素和为 $s$，那么第二部分的元素和为 $\textit{total}-s$。

- 不删元素：$s = \textit{total}-s$，即 $2s-\textit{total} = 0$。
- 删第一部分中的元素 $x$：$s - x = \textit{total}-s$，即 $2s-\textit{total} = x$。

据此，我们可以一边遍历 $\textit{grid}$，一边计算第一部分的元素和 $s$，一边用哈希集合记录遍历过的元素。

每一行/列遍历结束后，判断 $x=2s-\textit{total}$ 是否在哈希集合中，如果在，就说明存在 $x$，使得 $s - x = \textit{total}-s$ 成立。

小技巧：预先把 $0$ 加到哈希集合中，这样可以把不删和删合并成一种情况。

对于删第二部分中的元素的情况，可以把 $\textit{grid}$ 上下翻转，复用删第一部分中的元素的代码。

## 分类讨论

先计算水平分割的情况。

分类讨论：

- 对于只有一列（$n=1$）的情况，只能删除第一个数或者分割线上那个数。
- 对于只有一行（分割线在第一行与第二行之间）的情况，只能删除第一个数或者最后一个数。删除中间的数会导致第一部分不连通。
- 其余情况，可以随便删。

对于垂直分割，可以把 $\textit{grid}$ 旋转 $90$ 度，复用上述代码。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1h3EuzrEqV/?t=10m36s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def canPartitionGrid(self, grid: List[List[int]]) -> bool:
        total = sum(sum(row) for row in grid)

        # 能否水平分割
        def check(a: List[List[int]]) -> bool:
            m, n = len(a), len(a[0])

            # 删除上半部分中的一个数，能否满足要求
            def f(a: List[List[int]]) -> bool:
                st = {0}  # 0 对应不删除数字
                s = 0
                for i, row in enumerate(a[:-1]):
                    for j, x in enumerate(row):
                        s += x
                        # 第一行，不能删除中间元素
                        if i > 0 or j == 0 or j == n - 1:
                            st.add(x)
                    # 特殊处理只有一列的情况，此时只能删除第一个数或者分割线上那个数
                    if n == 1:
                        if s * 2 == total or s * 2 - total == a[0][0] or s * 2 - total == row[0]:
                            return True
                        continue
                    if s * 2 - total in st:
                        return True
                    # 如果分割到更下面，那么可以删第一行的元素
                    if i == 0:
                        st.update(row)
                return False

            # 删除上半部分中的数 or 删除下半部分中的数
            return f(a) or f(a[::-1])

        # 水平分割 or 垂直分割
        return check(grid) or check(list(zip(*grid)))
```

```java [sol-Java]
class Solution {
    public boolean canPartitionGrid(int[][] grid) {
        long total = 0;
        for (int[] row : grid) {
            for (int x : row) {
                total += x;
            }
        }
        return check(grid, total) || check(rotate(grid), total);
    }

    private boolean check(int[][] a, long total) {
        if (f(a, total)) {
            return true;
        }
        reverse(a);
        return f(a, total);
    }

    private boolean f(int[][] a, long total) {
        int m = a.length, n = a[0].length;
        Set<Long> st = new HashSet<>();
        st.add(0L); // 0 对应不删除数字
        long s = 0;
        for (int i = 0; i < m - 1; i++) {
            int[] row = a[i];
            for (int j = 0; j < n; j++) {
                int x = row[j];
                s += x;
                // 第一行，不能删除中间元素
                if (i > 0 || j == 0 || j == n - 1) {
                    st.add((long) x);
                }
            }
            // 特殊处理只有一列的情况，此时只能删除第一个数或者分割线上那个数
            if (n == 1) {
                if (s * 2 == total || s * 2 - total == a[0][0] || s * 2 - total == row[0]) {
                    return true;
                }
                continue;
            }
            if (st.contains(s * 2 - total)) {
                return true;
            }
            // 如果分割到更下面，那么可以删第一行的元素
            if (i == 0) {
                for (int x : row) {
                    st.add((long) x);
                }
            }
        }
        return false;
    }

    // 顺时针旋转矩阵 90°
    private int[][] rotate(int[][] a) {
        int m = a.length, n = a[0].length;
        int[][] b = new int[n][m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                b[j][m - 1 - i] = a[i][j];
            }
        }
        return b;
    }

    private void reverse(int[][] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int[] tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    // 顺时针旋转矩阵 90°
    vector<vector<int>> rotate(vector<vector<int>>& a) {
        int m = a.size(), n = a[0].size();
        vector b(n, vector<int>(m));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                b[j][m - 1 - i] = a[i][j];
            }
        }
        return b;
    }

public:
    bool canPartitionGrid(vector<vector<int>>& grid) {
        long long total = 0;
        for (auto& row : grid) {
            for (int x : row) {
                total += x;
            }
        }

        auto check = [&](vector<vector<int>> a) -> bool {
            int m = a.size(), n = a[0].size();

            auto f = [&]() -> bool {
                unordered_set<long long> st = {0}; // 0 对应不删除数字
                long long s = 0;
                for (int i = 0; i < m - 1; i++) {
                    auto& row = a[i];
                    for (int j = 0; j < n; j++) {
                        int x = row[j];
                        s += x;
                        // 第一行，不能删除中间元素
                        if (i > 0 || j == 0 || j == n - 1) {
                            st.insert(x);
                        }
                    }
                    // 特殊处理只有一列的情况，此时只能删除第一个数或者分割线上那个数
                    if (n == 1) {
                        if (s * 2 == total || s * 2 - total == a[0][0] || s * 2 - total == row[0]) {
                            return true;
                        }
                        continue;
                    }
                    if (st.contains(s * 2 - total)) {
                        return true;
                    }
                    // 如果分割到更下面，那么可以删第一行的元素
                    if (i == 0) {
                        for (int x : row) {
                            st.insert(x);
                        }
                    }
                }
                return false;
            };

            // 删除上半部分中的一个数
            if (f()) {
                return true;
            }
            // 删除下半部分中的一个数
            ranges::reverse(a);
            return f();
        };

        // 水平分割 or 垂直分割
        return check(grid) || check(rotate(grid));
    }
};
```

```go [sol-Go]
func canPartitionGrid(grid [][]int) bool {
	total := 0
	for _, row := range grid {
		for _, x := range row {
			total += x
		}
	}

	// 能否水平分割
	check := func(a [][]int) bool {
		m, n := len(a), len(a[0])
		f := func() bool {
			has := map[int]bool{0: true} // 0 对应不删除数字
			s := 0
			for i, row := range a[:m-1] {
				for j, x := range row {
					s += x
					// 第一行，不能删除中间元素
					if i > 0 || j == 0 || j == n-1 {
						has[x] = true
					}
				}
				// 特殊处理只有一列的情况，此时只能删除第一个数或者分割线上那个数
				if n == 1 {
					if s*2 == total || s*2-total == a[0][0] || s*2-total == row[0] {
						return true
					}
					continue
				}
				if has[s*2-total] {
					return true
				}
				// 如果分割到更下面，那么可以删第一行的元素
				if i == 0 {
					for _, x := range row {
						has[x] = true
					}
				}
			}
			return false
		}
		// 删除上半部分中的一个数
		if f() {
			return true
		}
		// 删除下半部分中的一个数
		slices.Reverse(a)
		return f()
	}

	// 水平分割 or 垂直分割
	return check(grid) || check(rotate(grid))
}

// 顺时针旋转矩阵 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
